package trithemius

import (
	"slices"
)

// Trithemius TODO Убрать магические числа в коде (осталась длина алфавита)
type Trithemius struct {
	Alphabet *Alphabet
}

const DefaultShift int = 8

func NewTrithemius(TelegraphAlphabet []rune) *Trithemius {
	return &Trithemius{NewAlphabet(TelegraphAlphabet)}
}

func NewTrithemiusWithReadyAlphabet(Alphabet Alphabet) *Trithemius {
	return &Trithemius{Alphabet: &Alphabet}
}

func (a *Trithemius) BuildTrithemiusAlphabet(key string) []rune {
	var TrithemiusAlphabet []rune

	for _, char := range []rune(key) {
		temp := char
		if len(TrithemiusAlphabet) == a.Alphabet.AlphabetLength {
			break
		}
		for slices.Contains(TrithemiusAlphabet, temp) {
			receivedKey := a.Alphabet.GetKeyByChar(string(temp))
			temp = []rune(a.Alphabet.GetCharByKey((receivedKey + 1) % 32))[0]
		}
		TrithemiusAlphabet = append(TrithemiusAlphabet, temp)
	}

	for _, r := range a.Alphabet.AlphabetArr {
		if len(TrithemiusAlphabet) == a.Alphabet.AlphabetLength {
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
		newChar := table[(pos+DefaultShift)%32]
		out = append(out, newChar)
	}
	return string(out)
}

func (a *Trithemius) DecodeTrithemius(text string, table []rune) string {
	var out []rune
	for _, char := range []rune(text) {
		pos := slices.Index(table, char)
		newChar := table[(32+pos-DefaultShift)%32]
		out = append(out, newChar)
	}
	return string(out)
}

func (a *Trithemius) ShiftTrithemiusAlphabet(table []rune, char string, bias int) []rune {
	s := []rune(char)[0]
	str := table[bias:]
	rem := table[:bias]
	for slices.Contains(rem, s) {
		receivedKey := a.Alphabet.GetKeyByChar(string(s))
		s = []rune(a.Alphabet.GetCharByKey((receivedKey + 1) % 32))[0]
	}
	x := slices.Index(str, s)
	str = slices.Concat(str[:x], str[x+1:])
	return slices.Concat([]rune{s}, rem, str)
}

func (a *Trithemius) EncodePolyTrithemius(text string, key string) string {
	var res []rune

	runeText := []rune(text)
	table := a.BuildTrithemiusAlphabet(key)
	keyArr := a.Alphabet.TextToArray(key)
	for i, char := range runeText {

		k := i % len(keyArr)
		b := (i + len(keyArr)) % 32
		encodedChar := a.EncodeTrithemius(string(char), table)
		table = a.ShiftTrithemiusAlphabet(table, a.Alphabet.GetCharByKey(keyArr[k]), b)

		res = append(res, []rune(encodedChar)[0])
	}

	return string(res)
}

func (a *Trithemius) DecodePolyTrithemius(text string, key string) string {
	var res []rune

	runeText := []rune(text)
	table := a.BuildTrithemiusAlphabet(key)
	keyArr := a.Alphabet.TextToArray(key)
	for i, char := range runeText {

		k := i % len(keyArr)
		b := (i + len(keyArr)) % 32
		encodedChar := a.DecodeTrithemius(string(char), table)
		table = a.ShiftTrithemiusAlphabet(table, a.Alphabet.GetCharByKey(keyArr[k]), b)

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

	keyArr := a.Alphabet.TextToArray(key)
	var sum int
	for i := 0; i < 16; i++ {
		sign := 1
		if i%2 == 1 {
			sign = -1
		}
		sum = ((24+sum+sign*keyArr[i])%24 + 24) % 24
	}
	for k := 0; k < 3; k++ {
		t := sum % (4 - k)
		sum = (sum - t) / (4 - k)
		temp := M[k]
		M[k] = M[k+t]
		M[k+t] = temp
	}
	blockArr := a.Alphabet.TextToArray(block)
	for j := 0; j < 4; j++ {
		b := M[(1+j)%4]
		c := M[j%4]
		blockArr[b] = (blockArr[b] + blockArr[c]) % 32
	}
	return a.Alphabet.ArrayToText(blockArr)
}

func (a *Trithemius) DecodeMergeBlock(block, key string) string {
	if !a.validBlockAndKey(block, key) {
		return "input error"
	}

	M := []int{0, 1, 2, 3}
	keyArr := a.Alphabet.TextToArray(key)
	var sum int
	for i := 0; i < 16; i++ {
		sign := 1
		if i%2 == 1 {
			sign = -1
		}
		sum = ((24+sum+sign*keyArr[i])%24 + 24) % 24
	}
	for k := 0; k < 3; k++ {
		t := sum % (4 - k)
		sum = (sum - t) / (4 - k)
		temp := M[k]
		M[k] = M[k+t]
		M[k+t] = temp
	}
	blockArr := a.Alphabet.TextToArray(block)
	for j := 3; j >= 0; j-- {
		b := M[(1+j)%4]
		c := M[j%4]
		blockArr[b] = (32 + blockArr[b] - blockArr[c]) % 32
	}
	return a.Alphabet.ArrayToText(blockArr)
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

func (a *Trithemius) validBlockAndKey(block, key string) bool {
	const (
		shortBlockLen = 4
		fullKeyLen    = 16
	)

	return len([]rune(block)) == shortBlockLen && len([]rune(key)) == fullKeyLen
}
