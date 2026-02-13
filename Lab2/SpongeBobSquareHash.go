package Lab2

import (
	l1 "infbez_labs/Lab1"
)

type Sponge struct {
	InnerState [4]string
	Alphabet   l1.Alphabet
}

func NewSponge(InnerState [4]string, Alphabet l1.Alphabet) *Sponge {
	return &Sponge{InnerState, Alphabet}
}

func (s *Sponge) CoreTrithemius(inPrime, inAux string) string {
	if (len([]rune(inPrime)) != 16) || (len([]rune(inAux)) != 16) {
		panic("неподходящая длина входных данных")
	}
	newTrithemius := l1.NewTrithemiusWithReadyAlphabet(s.Alphabet)

	return newTrithemius.EncodePolyTrithemius(inAux, inPrime)
}

func (s *Sponge) CBlock(inArr []string, outSize int) string {
	var (
		CountRow   = len(inArr)
		innerState = s.InnerState
	)
	for i := 0; i < CountRow; i++ {
		if len([]rune(inArr[i])) == 16 {
			innerState[i] = s.Alphabet.AddTxt(innerState[i], inArr[i])
		} else {
			panic("ошибка входных данных: строка должна быть длиной 16 символов")
		}
	}
	innerState[1] = s.Alphabet.AddTxt(innerState[1], inArr[0])
	TMP1 := s.CoreTrithemius(innerState[0], innerState[2])
	TMP2 := s.CoreTrithemius(innerState[3], innerState[1])
	TMP3 := s.Confuse(TMP1, TMP2)

	result := s.CoreTrithemius(TMP3, TMP1)
	result = s.Compress(result, outSize)
	return result

}

func (s *Sponge) Confuse(In1, In2 string) string {
	arr1 := s.Alphabet.TextToArray(In1)
	arr2 := s.Alphabet.TextToArray(In2)
	for i := 0; i < 16; i++ {
		arr1[i] = (max(arr1[i], arr2[i]) + i) % s.Alphabet.AlphabetLength
	}
	return s.Alphabet.ArrayToText(arr1)
}

func (s *Sponge) Compress(In16 string, outN int) string {
	if outN == 16 {
		return In16
	}
	var (
		In16Arr = []rune(In16)
		a1      = In16Arr[0:4]
		a2      = In16Arr[4:8]
		a3      = In16Arr[8:12]
		a4      = In16Arr[12:16]
	)
	if outN == 8 {
		a13 := append([]rune(nil), a1...)
		a13 = append(a13, a3...)

		a24 := append([]rune(nil), a2...)
		a24 = append(a24, a4...)

		return s.Alphabet.AddTxt(string(a13), string(a24))
	} else if outN == 4 {
		var (
			a13 = s.Alphabet.SubTxt(string(a1), string(a3))
			a24 = s.Alphabet.SubTxt(string(a2), string(a4))
		)
		return s.Alphabet.AddTxt(a13, a24)
	} else {
		panic("недопустимый размер выходных данных")
	}
}
