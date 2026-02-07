package Thrithemus

import (
	"fmt"
	"slices"
)

var TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")

type Alphabet struct {
	alphabet []rune
	key      string
}

func NewAlphabet() *Alphabet {
	return &Alphabet{TelegraphAlphabet, ""}
}

func (a *Alphabet) GetCharByKey(key int) string {
	return string(a.alphabet[(32+key-1)%32])
}

func (a *Alphabet) GetKeyByChar(char string) int {
	for i, r := range a.alphabet {
		if r == []rune(char)[0] {
			return (i + 1) % 32
		}
	}

	return -1
}

func (a *Alphabet) TextToArray(text string) []int {
	var res []int
	for _, char := range text {
		res = append(res, a.GetKeyByChar(string(char)))
	}
	return res
}

func (a *Alphabet) ArrayToText(array []int) string {
	res := ""
	for _, key := range array {
		res += a.GetCharByKey(key)
	}
	return res
}

func (a *Alphabet) AddChars(charX, charY string) string {
	charXIndex := a.GetKeyByChar(charX)
	charYIndex := a.GetKeyByChar(charY)

	return a.GetCharByKey(charXIndex + charYIndex)
}

func (a *Alphabet) SubtractChars(charX, charY string) string {
	charXIndex := a.GetKeyByChar(charX)
	charYIndex := a.GetKeyByChar(charY)

	return a.GetCharByKey(32 + charXIndex - charYIndex)
}

func (a *Alphabet) BuildThrithemusAlphabet(key string) []rune {
	var ThrithemusAlphabet []rune

	for _, char := range []rune(key) {
		temp := char
		if len(ThrithemusAlphabet) == 32 {
			break
		}
		for slices.Contains(ThrithemusAlphabet, temp) {
			temp = []rune(a.GetCharByKey((a.GetKeyByChar(string(temp)) + 1) % 32))[0]
		}
		ThrithemusAlphabet = append(ThrithemusAlphabet, temp)
	}

	for _, r := range TelegraphAlphabet {
		if len(ThrithemusAlphabet) == 32 {
			break
		}
		if !slices.Contains(ThrithemusAlphabet, r) {
			ThrithemusAlphabet = append(ThrithemusAlphabet, r)
		}
	}
	fmt.Println("вышел")
	return ThrithemusAlphabet
}

func (a *Alphabet) EncodeThrithemus(text string, table []rune) string {
	var out []rune
	for _, char := range []rune(text) {
		temp := char
		pos := slices.Index(table, temp)
		newChar := table[(pos+8)%32]
		out = append(out, newChar)
	}
	//fmt.Println(text, string(out), string(table))
	return string(out)
}

func (a *Alphabet) DecodeThrithemus(text string, table []rune) string {
	var out []rune
	for _, char := range []rune(text) {
		temp := char
		pos := slices.Index(table, temp)
		newChar := table[(32+pos-8)%32]
		out = append(out, newChar)
	}
	return string(out)
}

func (a *Alphabet) ShiftThrithemusAlphabet(table []rune, char string, bias int) []rune {
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

func (a *Alphabet) EncodePolyThrithemus(text string, key string) string {
	var res []rune

	runeText := []rune(text)
	table := a.BuildThrithemusAlphabet(key)
	keyArr := a.TextToArray(key)
	for i, char := range runeText {

		k := i % len(keyArr)
		b := (i + len(keyArr)) % 32
		encodedChar := a.EncodeThrithemus(string(char), table)
		table = a.ShiftThrithemusAlphabet(table, a.GetCharByKey(keyArr[k]), b)

		res = append(res, []rune(encodedChar)[0])
	}

	return string(res)
}

func (a *Alphabet) DecodePolyThrithemus(text string, key string) string {
	var res []rune

	runeText := []rune(text)
	table := a.BuildThrithemusAlphabet(key)
	keyArr := a.TextToArray(key)
	for i, char := range runeText {

		k := i % len(keyArr)
		b := (i + len(keyArr)) % 32
		encodedChar := a.DecodeThrithemus(string(char), table)
		table = a.ShiftThrithemusAlphabet(table, a.GetCharByKey(keyArr[k]), b)

		res = append(res, []rune(encodedChar)[0])
	}

	return string(res)
}

func (a *Alphabet) EncodeSThrithemus(block, key string) string {
	res := "input error"
	if len([]rune(block)) == 4 && len([]rune(key)) == 16 {
		res = a.EncodePolyThrithemus(block, key)
	}
	return res
}

func (a *Alphabet) DecodeSThrithemus(block, key string) string {
	res := "input error"
	if len([]rune(block)) == 4 && len([]rune(key)) == 16 {
		res = a.DecodePolyThrithemus(block, key)
	}
	return res
}

func (a *Alphabet) EncodeMergeBlock(block, key string) string {
	res := "input error"
	M := []int{0, 1, 2, 3}
	if len([]rune(block)) == 4 && len([]rune(key)) == 16 {
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
		res = a.ArrayToText(blockArr)
	}
	return res
}

func (a *Alphabet) DecodeMergeBlock(block, key string) string {
	res := "input error"
	M := []int{0, 1, 2, 3}
	if len([]rune(block)) == 4 && len([]rune(key)) == 16 {
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
		res = a.ArrayToText(blockArr)
	}
	return res
}

func (a *Alphabet) EncodeSThrithemusM(block, key string) string {
	temp := a.DecodeMergeBlock(block, key)
	temp = a.EncodeSThrithemus(temp, key)
	return a.EncodeMergeBlock(temp, key)
}

func (a *Alphabet) DecodeSThrithemusM(block, key string) string {
	temp := a.DecodeMergeBlock(block, key)
	temp = a.DecodeSThrithemus(temp, key)
	return a.EncodeMergeBlock(temp, key)
}
