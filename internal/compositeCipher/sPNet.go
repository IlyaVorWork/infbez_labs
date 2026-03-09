package compositeCipher

import (
	generator "infbez_labs/internal/codeRandomGenerator"
	"infbez_labs/internal/core"
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
	sBlock core.SBlock
	lfsr   generator.LFSR
}

func NewSPNet(sBlock core.SBlock, lfsr generator.LFSR) *SPNet {
	return &SPNet{
		sBlock: sBlock,
		lfsr:   lfsr,
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
