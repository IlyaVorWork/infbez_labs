package alphabet

import (
	"errors"
	"math"
	"strings"
)

var TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")

type Alphabet struct {
	AlphabetArr    []rune
	AlphabetLength int
	AlphabetMap    map[rune]int
}

func NewAlphabet(inputAlphabet []rune) *Alphabet {
	AlphabetArr := inputAlphabet
	AlphabetLength := len(inputAlphabet)

	alphabetMap := make(map[rune]int, AlphabetLength)
	for i, char := range AlphabetArr {
		alphabetMap[char] = (i + 1) % AlphabetLength
	}
	return &Alphabet{AlphabetArr, len(AlphabetArr), alphabetMap}
}

func (a *Alphabet) GetCharByKey(key int) rune {
	return a.AlphabetArr[(a.AlphabetLength+key-1)%a.AlphabetLength]
}

func (a *Alphabet) GetKeyByChar(char rune) (int, bool) {
	item, ok := a.AlphabetMap[char]
	return item, ok
}

func (a *Alphabet) TextToArray(text []rune) []int {
	var res []int
	for _, char := range text {
		key, ok := a.GetKeyByChar(char)
		if !ok {
			panic("TextToArray: неправильный символ")
		}

		res = append(res, key)
	}
	return res
}

func (a *Alphabet) ArrayToText(array []int) []rune {
	var output = make([]rune, 0, len(array))

	for _, key := range array {
		char := a.GetCharByKey(key)
		output = append(output, char)
	}
	return output
}

func (a *Alphabet) AddChars(charX, charY rune) rune {
	charXIndex, okX := a.GetKeyByChar(charX)
	charYIndex, okY := a.GetKeyByChar(charY)
	if !okX || !okY {
		panic("AddChars: неправильный символ")
	}
	resultChar := a.GetCharByKey(charXIndex + charYIndex)
	return resultChar
}

func (a *Alphabet) SubtractChars(charX, charY rune) rune {
	charXIndex, okX := a.GetKeyByChar(charX)
	charYIndex, okY := a.GetKeyByChar(charY)
	if !okX || !okY {
		panic("AddChars: неправильный символ")
	}
	resultChar := a.GetCharByKey(a.AlphabetLength + charXIndex - charYIndex)
	return resultChar
}

func (a *Alphabet) AddTxt(txt1, txt2 string) string {
	r1 := []rune(txt1)
	r2 := []rune(txt2)

	if len(r1) < len(r2) {
		r1, r2 = r2, r1
	}
	var builder strings.Builder
	builder.Grow(len(r1))

	for i := 0; i < len(r2); i++ {
		builder.WriteRune(a.AddChars(r1[i], r2[i]))
	}
	for i := len(r2); i < len(r1); i++ {
		builder.WriteString(string(r1[i]))
	}

	return builder.String()
}

func (a *Alphabet) SubTxt(txt1, txt2 string) string {
	r1 := []rune(txt1)
	r2 := []rune(txt2)

	flag := 0
	TIN := r1
	if len(r1) < len(r2) {
		TIN = r2
		flag = 1
	}

	m := min(len(r1), len(r2))
	M := len(TIN)

	var builder strings.Builder
	builder.Grow(M)

	for i := 0; i < m; i++ {
		builder.WriteRune(a.SubtractChars(r1[i], r2[i]))
	}
	placeholder := '_'

	for i := m; i < M; i++ {
		t := TIN[i]
		if flag == 1 {
			builder.WriteRune(a.SubtractChars(placeholder, t))
		} else {
			builder.WriteRune(a.SubtractChars(t, placeholder))
		}
	}

	return builder.String()
}

func (a *Alphabet) BlockToNum(block []rune) int {
	if len(block) != 4 {
		panic("BlockToNum: вход имеет неверную длину")
	}
	pos := 1
	temp := a.TextToArray(block)
	var out int
	for i := 3; i >= 0; i-- {
		out = pos*temp[i] + out
		pos = 32 * pos
	}
	return out
}

func (a *Alphabet) div(dividend, divisor int) int {
	return int(math.Trunc(float64(dividend) / float64(divisor)))
}

func (a *Alphabet) NumToBlock(num int) []rune {
	temp := [4]int{}
	for i := 0; i < 4; i++ {
		temp[3-i] = num % a.AlphabetLength
		num = a.div(num, a.AlphabetLength)
	}
	return a.ArrayToText(temp[:])
}

func (a *Alphabet) DecToBin(num int) []int {
	rem := num
	out := make([]int, 20)
	for i := 0; i < 20; i++ {
		out[19-i] = rem % 2
		rem = a.div(rem, 2)
	}
	return out[:]
}

func (a *Alphabet) BinToDec(nums []int) int {
	out := 0
	for i := 0; i < len(nums); i++ {
		out = out*2 + nums[i]
	}
	return out
}

func (a *Alphabet) BlockToBin(block []rune) []int {
	temp := a.BlockToNum(block)
	return a.DecToBin(temp)
}

func (a *Alphabet) BinToBlock(bin []int) []rune {
	temp := a.BinToDec(bin)
	return a.NumToBlock(temp)
}

func (a *Alphabet) PushReg(bin []int, boolIn int) []int {
	return append(bin[1:], boolIn)
}

// -- Для S-P сети

func (a *Alphabet) SubBlocksXOR(miniBlock1, miniBlock2 []rune) []rune {
	var (
		miniNumBlock1 = a.BlockToNum(miniBlock1)
		miniNumBlock2 = a.BlockToNum(miniBlock2)
	)
	return a.NumToBlock(miniNumBlock1 ^ miniNumBlock2)
}

func (a *Alphabet) BlockXOR(block1, block2 []rune) []rune {
	if len(block1) != len(block2) || len(block1)%4 != 0 {
		panic("BlockXOR: вход имеет неверную длину")
	}

	var resultRunes = make([]rune, 0, len(block1))

	for i := 0; i < len(block1); i += 4 {
		sub1 := block1[i : i+4]
		sub2 := block2[i : i+4]
		resultRunes = append(resultRunes, a.SubBlocksXOR(sub1, sub2)...)
	}
	return resultRunes
}

func (a *Alphabet) MessageToBin(msgArr []rune) ([]byte, error) {
	var (
		msgLen        = len(msgArr)
		output        = make([]byte, 0, msgLen*5)
		charBin       [5]byte
		lastCharIndex = msgLen - 1
	)

	for index, char := range msgArr {
		if char == '0' || char == '1' {
			lastCharIndex = index
			break
		}

		key, ok := a.GetKeyByChar(char)
		if !ok {
			return []byte{}, errors.New("MessageToBin: плохой символ")
		}

		charBin = a.NumToBin(key)
		output = append(output, charBin[:]...)
	}
	output = append(output, a.binStrToNum(msgArr[lastCharIndex:])...)
	return output, nil
}

func (a *Alphabet) BinToMessage(bins []byte) []rune {
	var (
		wholeBlocks = len(bins) / 5
		remains     = len(bins) % 5
		msg         = make([]rune, 0, wholeBlocks+remains)
	)

	for i := 0; i < wholeBlocks; i++ {
		numOfChar := a.BinToNum([5]byte(bins[i*5 : i*5+5]))
		char := a.GetCharByKey(numOfChar)
		msg = append(msg, char)
	}

	for i := 0; i < remains; i++ {
		numOfChar := bins[wholeBlocks*5+i]
		msg = append(msg, a.numToStrBin(numOfChar))
	}
	return msg
}

func (a *Alphabet) binStrToNum(strBins []rune) []byte {
	var output []byte
	for i := range strBins {
		switch strBins[i] {
		case '0':
			output = append(output, 0)
		case '1':
			output = append(output, 1)
		}
	}
	return output
}

func (a *Alphabet) numToStrBin(num byte) rune {
	var runeOutput rune

	switch num {
	case 0:
		runeOutput = '0'
	case 1:
		runeOutput = '1'
	}

	return runeOutput
}

// NumToBin Возвращает двоичное число (5 символов) из числа
func (a *Alphabet) NumToBin(num int) [5]byte {
	rem := num
	out := make([]byte, 5)
	for i := 0; i < 5; i++ {
		out[4-i] = byte(rem % 2)
		rem = a.div(rem, 2)
	}
	return [5]byte(out)
}

// BinToNum Возвращает число из двоичного числа (5 символов)
func (a *Alphabet) BinToNum(nums [5]byte) int {
	out := 0
	for i := 0; i < len(nums); i++ {
		out = out*2 + int(nums[i])
	}
	return out
}
