package trithemius

import (
	"slices"
)

// Trithemius TODO Убрать магические числа в коде
type Trithemius struct {
	alphabet       []rune
	alphabetLength int //TODO Правильно использовать в функциях, а не просто 32
}

func NewTrithemius(TelegraphAlphabet []rune) *Trithemius {
	return &Trithemius{TelegraphAlphabet, len(TelegraphAlphabet)} //TODO Узнать как правильно задать значение по умолчанию
}

func (a *Trithemius) GetCharByKey(key int) string {
	return string(a.alphabet[(32+key-1)%32])
}

func (a *Trithemius) GetKeyByChar(char string) int {
	for i, r := range a.alphabet {
		if r == []rune(char)[0] {
			return (i + 1) % 32
		}
	}
	return -1
}

func (a *Trithemius) TextToArray(text string) []int {
	var res []int
	for _, char := range text {
		res = append(res, a.GetKeyByChar(string(char)))
	}
	return res
}

func (a *Trithemius) ArrayToText(array []int) string {
	res := ""
	for _, key := range array {
		res += a.GetCharByKey(key)
	}
	return res
}

func (a *Trithemius) AddChars(charX, charY string) string {
	charXIndex := a.GetKeyByChar(charX)
	charYIndex := a.GetKeyByChar(charY)

	return a.GetCharByKey(charXIndex + charYIndex)
}

func (a *Trithemius) SubtractChars(charX, charY string) string {
	charXIndex := a.GetKeyByChar(charX)
	charYIndex := a.GetKeyByChar(charY)

	return a.GetCharByKey(32 + charXIndex - charYIndex)
}

func (a *Trithemius) BuildTrithemiusAlphabet(key string) []rune {
	var TrithemiusAlphabet []rune

	for _, char := range []rune(key) {
		temp := char
		if len(TrithemiusAlphabet) == 32 {
			break
		}
		for slices.Contains(TrithemiusAlphabet, temp) {
			temp = []rune(a.GetCharByKey((a.GetKeyByChar(string(temp)) + 1) % 32))[0]
		}
		TrithemiusAlphabet = append(TrithemiusAlphabet, temp)
	}

	for _, r := range a.alphabet {
		if len(TrithemiusAlphabet) == 32 {
			break
		}
		if !slices.Contains(TrithemiusAlphabet, r) {
			TrithemiusAlphabet = append(TrithemiusAlphabet, r)
		}
	}
	return TrithemiusAlphabet
}

func (a *Trithemius) EncodeTrithemius(text string, table []rune) string {
	var out []rune
	for _, char := range []rune(text) {
		pos := slices.Index(table, char)
		newChar := table[(pos+8)%32]
		out = append(out, newChar)
	}
	return string(out)
}

func (a *Trithemius) DecodeTrithemius(text string, table []rune) string {
	var out []rune
	for _, char := range []rune(text) {
		pos := slices.Index(table, char)
		newChar := table[(32+pos-8)%32]
		out = append(out, newChar)
	}
	return string(out)
}

func (a *Trithemius) ShiftTrithemiusAlphabet(table []rune, char string, bias int) []rune {
	s := []rune(char)[0]
	str := table[bias:]
	rem := table[:bias]
	for slices.Contains(rem, s) {
		s = []rune(a.GetCharByKey((a.GetKeyByChar(string(s)) + 1) % 32))[0]
	}
	x := slices.Index(str, s)
	str = slices.Concat(str[:x], str[x+1:])
	return slices.Concat([]rune{s}, rem, str)
}

func (a *Trithemius) EncodePolyTrithemius(text string, key string) string {
	var res []rune

	runeText := []rune(text)
	table := a.BuildTrithemiusAlphabet(key)
	keyArr := a.TextToArray(key)
	for i, char := range runeText {

		k := i % len(keyArr)
		b := (i + len(keyArr)) % 32
		encodedChar := a.EncodeTrithemius(string(char), table)
		table = a.ShiftTrithemiusAlphabet(table, a.GetCharByKey(keyArr[k]), b)

		res = append(res, []rune(encodedChar)[0])
	}

	return string(res)
}

func (a *Trithemius) DecodePolyTrithemius(text string, key string) string {
	var res []rune

	runeText := []rune(text)
	table := a.BuildTrithemiusAlphabet(key)
	keyArr := a.TextToArray(key)
	for i, char := range runeText {

		k := i % len(keyArr)
		b := (i + len(keyArr)) % 32
		encodedChar := a.DecodeTrithemius(string(char), table)
		table = a.ShiftTrithemiusAlphabet(table, a.GetCharByKey(keyArr[k]), b)

		res = append(res, []rune(encodedChar)[0])
	}

	return string(res)
}

func (a *Trithemius) EncodeSTrithemius(block, key string) string {
	if !a.validBlockAndKey(block, key) {
		return "input error"
	}
	return a.EncodePolyTrithemius(block, key)

}

func (a *Trithemius) DecodeSTrithemius(block, key string) string {
	if !a.validBlockAndKey(block, key) {
		return "input error"
	}
	return a.DecodePolyTrithemius(block, key)
}

func (a *Trithemius) EncodeMergeBlock(block, key string) string {
	if !a.validBlockAndKey(block, key) {
		return "input error"
	}

	M := []int{0, 1, 2, 3}

	keyArr := a.TextToArray(key)
	var sum int
	for i := 0; i < 16; i++ {
		sign := 1
		if i%2 == 1 {
			sign = -1
		}
		sum = (24 + sum + sign*keyArr[i]) % 24
	}
	for k := 0; k < 3; k++ {
		t := sum % (4 - k)
		sum = (sum - t) / (4 - k)
		temp := M[k]
		M[k] = M[k+t]
		M[k+t] = temp
	}
	blockArr := a.TextToArray(block)
	for j := 0; j < 4; j++ {
		b := M[(1+j)%4]
		c := M[j%4]
		blockArr[b] = (blockArr[b] + blockArr[c]) % 32
	}
	return a.ArrayToText(blockArr)
}

func (a *Trithemius) DecodeMergeBlock(block, key string) string {
	if !a.validBlockAndKey(block, key) {
		return "input error"
	}

	M := []int{0, 1, 2, 3}
	keyArr := a.TextToArray(key)
	var sum int
	for i := 0; i < 16; i++ {
		sign := 1
		if i%2 == 1 {
			sign = -1
		}
		sum = (24 + sum + sign*keyArr[i]) % 24
	}
	for k := 0; k < 3; k++ {
		t := sum % (4 - k)
		sum = (sum - t) / (4 - k)
		temp := M[k]
		M[k] = M[k+t]
		M[k+t] = temp
	}
	blockArr := a.TextToArray(block)
	for j := 3; j >= 0; j-- {
		b := M[(1+j)%4]
		c := M[j%4]
		blockArr[b] = (32 + blockArr[b] - blockArr[c]) % 32
	}
	return a.ArrayToText(blockArr)
}

func (a *Trithemius) EncodeSTrithemiusM(block, key string) string {
	temp := a.DecodeMergeBlock(block, key)
	temp = a.EncodeSTrithemius(temp, key)
	return a.EncodeMergeBlock(temp, key)
}

func (a *Trithemius) DecodeSTrithemiusM(block, key string) string {
	temp := a.DecodeMergeBlock(block, key)
	temp = a.DecodeSTrithemius(temp, key)
	return a.EncodeMergeBlock(temp, key)
}

// TODO Довольно узкая функция по магическим числам, нужно придумать более универсальную функцию для проверки входных данных
func (a *Trithemius) validBlockAndKey(block, key string) bool {
	const (
		shortBlockLen = 4
		fullKeyLen    = 16
	)

	return len([]rune(block)) == shortBlockLen && len([]rune(key)) == fullKeyLen
}
