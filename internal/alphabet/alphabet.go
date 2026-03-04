package alphabet

import (
	"math"
	"strings"
)

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

func (a *Alphabet) GetCharByKey(key int) string {
	return string(a.AlphabetArr[(a.AlphabetLength+key-1)%a.AlphabetLength])
}

func (a *Alphabet) GetKeyByChar(char string) int {
	item, ok := a.AlphabetMap[[]rune(char)[0]]
	if !ok {
		panic("cимвол не найден в алфавите")
	}
	return item
}

func (a *Alphabet) TextToArray(text string) []int {
	var res []int
	for _, char := range text {
		key := a.GetKeyByChar(string(char))
		res = append(res, key)
	}
	return res
}

func (a *Alphabet) ArrayToText(array []int) string {
	var builder strings.Builder
	builder.Grow(len(array))

	for _, key := range array {
		char := a.GetCharByKey(key)
		builder.WriteString(char)
	}
	return builder.String()
}

func (a *Alphabet) AddChars(charX, charY string) string {
	charXIndex := a.GetKeyByChar(charX)
	charYIndex := a.GetKeyByChar(charY)
	resultChar := a.GetCharByKey(charXIndex + charYIndex)
	return resultChar
}

func (a *Alphabet) SubtractChars(charX, charY string) string {
	charXIndex := a.GetKeyByChar(charX)
	charYIndex := a.GetKeyByChar(charY)
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
		char1 := string(r1[i])
		char2 := string(r2[i])
		builder.WriteString(a.AddChars(char1, char2))
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
		char1 := string(r1[i])
		char2 := string(r2[i])
		builder.WriteString(a.SubtractChars(char1, char2))
	}
	placeholder := "_"

	for i := m; i < M; i++ {
		t := string(TIN[i])
		if flag == 1 {
			builder.WriteString(a.SubtractChars(placeholder, t))
		} else {
			builder.WriteString(a.SubtractChars(t, placeholder))
		}
	}

	return builder.String()
}

func (a *Alphabet) BlockToNum(block string) int {
	runeBlock := []rune(block)
	if len(runeBlock) != 4 {
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

func (a *Alphabet) NumToBlock(num int) string {
	rem := num
	temp := [4]int{}
	for i := 0; i < 4; i++ {
		temp[3-i] = rem % 32
		rem = a.div(rem, 32)
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

func (a *Alphabet) BlockToBin(block string) []int {
	temp := a.BlockToNum(block)
	return a.DecToBin(temp)
}

func (a *Alphabet) BinToBlock(bin []int) string {
	temp := a.BinToDec(bin)
	return a.NumToBlock(temp)
}

func (a *Alphabet) PushReg(bin []int, bool_in int) []int {
	return append(bin[1:], bool_in)
}
