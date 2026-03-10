package compositeCipher

import (
	"infbez_labs/internal/alphabet"
	generator "infbez_labs/internal/codeRandomGenerator"
	"infbez_labs/internal/core"
	"strings"
)

var (
	LfsrSet = [][][]int{
		{
			generator.TapsToBin([]int{19, 18}),
			generator.TapsToBin([]int{18, 7}),
			generator.TapsToBin([]int{17, 3}),
		},
		{
			generator.TapsToBin([]int{19, 18}),
			generator.TapsToBin([]int{18, 7}),
			generator.TapsToBin([]int{16, 14, 13, 11}),
		},
		{
			generator.TapsToBin([]int{19, 18}),
			generator.TapsToBin([]int{18, 7}),
			generator.TapsToBin([]int{15, 13, 12, 10}),
		},
		{
			generator.TapsToBin([]int{19, 18}),
			generator.TapsToBin([]int{18, 7}),
			generator.TapsToBin([]int{14, 5, 3, 1}),
		},
	}
)

type SPNet struct {
	alphabet *alphabet.Alphabet
	sBlock   *core.SBlock
	pBlock   *core.PBlock
	lfsr     *generator.LFSR
}

func NewSPNet(alphabet *alphabet.Alphabet, sBlock *core.SBlock, pBlock *core.PBlock, lfsr *generator.LFSR) *SPNet {
	return &SPNet{
		alphabet: alphabet,
		sBlock:   sBlock,
		pBlock:   pBlock,
		lfsr:     lfsr,
	}
}

func (s *SPNet) ProduceRoundKeys(key string, roundNum int) []string {
	if roundNum < 1 {
		return []string{}
	}

	var (
		out   = make([]string, roundNum)
		state = make([][][]int, 0)
	)
	out[0], state = s.lfsr.WrapCAsLfsrNext("up", make([][][]int, 0), key, LfsrSet)

	if roundNum == 1 {
		return out
	}

	for i := 1; i < roundNum; i++ {
		out[i], state = s.lfsr.WrapCAsLfsrNext("down", state, "", LfsrSet)
	}
	return out
}

func (s *SPNet) FrwRoundSP(blockIn string, Key string, roundNum int) string {
	var (
		builder = strings.Builder{}
	)
	builder.Grow(16)

	for i := 0; i < 16; i += 4 {
		blockPart := blockIn[i : i+4]
		builder.WriteString(s.sBlock.FrwRun(blockPart, Key))
	}
	afterPBlock := s.pBlock.FrwRound(builder.String(), roundNum)
	result := s.alphabet.BlockXOR(afterPBlock, Key)
	return result
}

func (s *SPNet) InvRoundSP(blockIn string, Key string, roundNum int) string {
	var (
		builder = strings.Builder{}
	)
	builder.Grow(16)

	afterXOR := s.alphabet.BlockXOR(blockIn, Key)
	afterPBlock := s.pBlock.InvRound(afterXOR, roundNum)

	for i := 0; i < 16; i += 4 {
		blockPart := afterPBlock[i : i+4]
		builder.WriteString(s.sBlock.FrwRun(blockPart, Key))
	}
	return builder.String()
}

func (s *SPNet) FrwSPNet(blockIn, key string, NumOfRound int) string {
	var (
		keys  = s.ProduceRoundKeys(key, NumOfRound)
		state = blockIn
	)

	for i := 0; i < NumOfRound; i++ {
		state = s.FrwRoundSP(state, keys[i], i)
	}

	return state
}

func (s *SPNet) InvSPNet(blockIn, key string, NumOfRound int) string {
	var (
		keys  = s.ProduceRoundKeys(key, NumOfRound)
		state = blockIn
	)

	for i := NumOfRound - 1; i >= 0; i-- {
		state = s.InvRoundSP(state, keys[i], i)
	}

	return state
}
