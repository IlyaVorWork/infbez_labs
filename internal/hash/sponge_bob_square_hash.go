package hash

import (
	"infbez_labs/internal/alphabet"
	"infbez_labs/internal/core"
	"strings"
)

var SpongeStarterState = [5][5]string{
	{"____", "____", "____", "____", "____"},
	{"____", "____", "____", "____", "____"},
	{"____", "____", "____", "____", "____"},
	{"____", "____", "____", "____", "____"},
	{"____", "____", "____", "____", "____"},
}

type Sponge struct {
	InnerState [5][5]string
	Alphabet   alphabet.Alphabet
	CBlock     core.CBlock
}

func NewSponge(spongeStarterState [5][5]string, Alphabet alphabet.Alphabet, CBlock core.CBlock) *Sponge {
	return &Sponge{
		spongeStarterState,
		Alphabet,
		CBlock}
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
	var str1 = concatStrings([]string{inputBlock, s.InnerState[0][0], inputBlock, s.InnerState[0][0]}, 16)

	for i := 0; i < 5; i++ {
		columnSums[i] = "____"
		for j := 0; j < 5; j++ {
			columnSums[i] = s.Alphabet.AddTxt(columnSums[i], s.InnerState[i][j])
		}
	}
	str2 := concatStrings([]string{columnSums[0], columnSums[1], columnSums[2], columnSums[3]}, 16)

	s.InnerState[0][0] = s.CBlock.Run([]string{str1, str2}, 4)
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
	return s.CBlock.Run([]string{str}, 4)
}

func concatStrings(stringsArr []string, capacity int) string {
	var builder strings.Builder
	builder.Grow(capacity)
	for _, str := range stringsArr {
		builder.WriteString(str)
	}
	return builder.String()
}
