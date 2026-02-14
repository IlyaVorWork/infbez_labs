package Lab2

import (
	l1 "infbez_labs/Lab1"
	"strings"
)

var (
	cBlockInnerState = [4]string{
		"________________",
		"ПРОЖЕКТОР_ЧЕПУХИ",
		"КОЛЫХАТЬ_ПАРОДИЮ",
		"КАРМАННЫЙ_АТАМАН",
	}

	SpongeInnerState = [5][5]string{
		{"____", "____", "____", "____", "____"},
		{"____", "____", "____", "____", "____"},
		{"____", "____", "____", "____", "____"},
		{"____", "____", "____", "____", "____"},
		{"____", "____", "____", "____", "____"},
	}
)

type Sponge struct {
	InnerState [5][5]string
	Alphabet   l1.Alphabet
}

func NewSponge(InnerState [5][5]string, Alphabet l1.Alphabet) *Sponge {
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
		innerState = cBlockInnerState
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

func (s *Sponge) MixCols() [5][5]string {
	var (
		X [5]string
	)
	for i := 0; i < 5; i++ {
		X[i] = "____"
		for j := 0; j < 5; j++ {
			X[i] = s.Alphabet.AddTxt(X[i], s.InnerState[j][i])
		}
		q := (i + 1) % 5
		for j := 0; j < 5; j++ {
			tmp := s.Alphabet.AddTxt(X[i], s.InnerState[j][q])
			s.InnerState[j][q] = s.Alphabet.SubTxt(tmp, s.InnerState[j][i])
		}
	}
	return s.InnerState
}

//func (s *Sponge) ShiftBlock(block string) string {
//	blockArr := s.Alphabet.TextToArray(block)
//	lenBlockArr := len(blockArr)
//	result := make([]int, lenBlockArr)
//	for i := 0; i < lenBlockArr; i++ {
//		result[(i+1)%lenBlockArr] = blockArr[i]
//	}
//	return s.Alphabet.ArrayToText(result)
//}

func (s *Sponge) ShiftBlock(block string) string {
	blockArr := s.Alphabet.TextToArray(block)
	shifted := append(blockArr[len(blockArr)-1:], blockArr[:len(blockArr)-1]...)
	return s.Alphabet.ArrayToText(shifted)
}

func (s *Sponge) ShatterBlocks() [5][5]string {
	for i := 0; i < 5; i++ {
		s.InnerState[i][i] = s.ShiftBlock(s.InnerState[i][i])
	}
	return s.InnerState
}

func (s *Sponge) ShiftRows() [5][5]string {
	var result [5][5]string

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			q := (j + i) % 5
			result[j][i] = s.InnerState[q][i]
		}
	}
	s.InnerState = result
	return result
}

func (s *Sponge) SpongeAbsorb(inputBlock string) [5][5]string {
	var columnSums [5]string
	var str2 = concatStrings([]string{inputBlock, s.InnerState[0][0], inputBlock, s.InnerState[0][0]}, 16)

	for i := 0; i < 5; i++ {
		columnSums[i] = "____"
		for j := 0; j < 5; j++ {
			columnSums[i] = s.Alphabet.AddTxt(columnSums[i], s.InnerState[i][j])
		}
	}
	str1 := concatStrings([]string{columnSums[0], columnSums[1], columnSums[2], columnSums[3]}, 16)

	s.InnerState[0][0] = s.CBlock([]string{str1, str2}, 4)
	s.MixCols()
	s.ShatterBlocks()
	s.ShiftRows()
	return s.InnerState
}

func (s *Sponge) SpongeSqueeze() string {
	s.MixCols()
	s.ShatterBlocks()
	s.ShiftRows()

	var columnSums [5]string

	for i := 0; i < 5; i++ {
		columnSums[i] = "____"
		for j := 0; j < 5; j++ {
			columnSums[i] = s.Alphabet.AddTxt(columnSums[i], s.InnerState[i][j])
		}
	}
	str := concatStrings([]string{columnSums[0], columnSums[1], columnSums[2], columnSums[3]}, 16)
	return s.CBlock([]string{str}, 4)
}

func SpongeHash(message string, alphabet l1.Alphabet) string {
	messageLen := len([]rune(message))
	sponge := NewSponge(SpongeInnerState, alphabet)
	var builder strings.Builder

	K := 4 - (messageLen % 4)
	if K < 4 {
		message = concatStrings([]string{message, strings.Repeat("_", K)}, messageLen+K)
	}

	messageArr := []rune(message)
	messageLen = len(messageArr)

	for i := 0; i < messageLen; i += 4 {
		tpm := messageArr[i : i+4]
		sponge.SpongeAbsorb(string(tpm))
	}

	builder.Grow(64)
	for i := 0; i < 16; i++ {
		tpm := sponge.SpongeSqueeze()
		builder.WriteString(tpm)
	}

	return builder.String()
}

func concatStrings(strs []string, capacity int) string {
	var builder strings.Builder
	builder.Grow(capacity)
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}
