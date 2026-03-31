package core

import (
	"infbez_labs/internal/alphabet"
)

var (
	MS1 = [][]int{
		{16, 3, 2, 13},
		{5, 10, 11, 8},
		{9, 6, 7, 12},
		{4, 15, 14, 1},
	}

	MS2 = [][]int{
		{7, 14, 4, 9},
		{12, 1, 15, 6},
		{13, 8, 10, 3},
		{2, 11, 5, 16},
	}

	MS3 = [][]int{
		{4, 14, 15, 1},
		{9, 7, 6, 12},
		{5, 11, 10, 8},
		{16, 2, 3, 13},
	}

	mSet = [][][]int{MS1, MS2, MS3}
)

type PBlock struct {
	alphabet *alphabet.Alphabet
}

func NewPBlock(alphabet *alphabet.Alphabet) *PBlock {
	return &PBlock{
		alphabet: alphabet,
	}
}

func (pb *PBlock) FrwRound(block string, roundNum int) string {
	numMS := roundNum % 3
	shift := 4*(roundNum%4) + 2
	blockAfterMS := pb.FrwMagicSquare(block, mSet[numMS])
	bitsBlock := pb.Text16CharTo80Bit(blockAfterMS)
	resultBlock := pb.BinaryShift(bitsBlock, shift)
	return pb.Text80BitTo16Char(resultBlock)
}

func (pb *PBlock) InvRound(block string, roundNum int) string {
	numMS := roundNum % 3
	shift := -(4*(roundNum%4) + 2)
	bits := pb.Text16CharTo80Bit(block)
	bitsAfterShift := pb.BinaryShift(bits, shift)
	blockAfterShift := pb.Text80BitTo16Char(bitsAfterShift)
	blockAfterMS := pb.InvMagicSquare(blockAfterShift, mSet[numMS])
	return blockAfterMS
}

func (pb *PBlock) FrwMagicSquare(block string, magicSquare [][]int) string {
	if !pb.isInput16CharValidate(block) {
		panic("InvMagicSquare: вход имеет неверную длину")
	}

	runes := []rune(block)
	result := make([]rune, 16)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i*4+j] = runes[magicSquare[i][j]-1]
		}
	}
	return string(result)
}

func (pb *PBlock) InvMagicSquare(block string, magicSquare [][]int) string {
	if !pb.isInput16CharValidate(block) {
		panic("InvMagicSquare: вход имеет неверную длину")
	}

	runes := []rune(block)
	result := make([]rune, 16)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[magicSquare[i][j]-1] = runes[i*4+j]
		}
	}
	return string(result)
}

func (pb *PBlock) BinaryShift(array []int, shift int) []int {
	arrLength := len(array)
	arrResult := make([]int, arrLength)
	b := shift % arrLength

	if b != 0 {
		for i := 0; i < arrLength; i++ {
			arrResult[(i+arrLength+b)%arrLength] = array[i]
		}
	}

	return arrResult
}

func (pb *PBlock) Text16CharTo80Bit(block string) []int {
	if !pb.isInput16CharValidate(block) {
		panic("Text16CharTo80Bit: вход имеет неверную длину")
	}

	var (
		inputBlockArr = []rune(block)
		bits          []int
		binsBlock     []int
	)

	for i := 0; i < 16; i += 4 {
		binsBlock = pb.alphabet.BlockToBin(inputBlockArr[i : i+4])
		bits = append(bits, binsBlock...)
	}
	return bits
}

func (pb *PBlock) Text80BitTo16Char(bits []int) string {
	if len(bits) != 80 {
		panic("Text80BitTo16Char: вход имеет неверную длину")
	}

	var (
		out       = make([]rune, 0, 16)
		BinsBlock []int
	)

	for i := 0; i < 4; i++ {
		BinsBlock = bits[i*20 : i*20+20]
		out = append(out, pb.alphabet.BinToBlock(BinsBlock)...)
	}
	return string(out)
}

func (pb *PBlock) isInput16CharValidate(block string) bool {
	return len([]rune(block)) == 16
}
