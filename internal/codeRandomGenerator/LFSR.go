package codeRandomGenerator

import (
	"infbez_labs/internal/alphabet"
	sponge "infbez_labs/internal/hash"
	"slices"
	"strings"
)

var (
	TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")

	SpongeInnerState = [5][5]string{
		{"____", "____", "____", "____", "____"},
		{"____", "____", "____", "____", "____"},
		{"____", "____", "____", "____", "____"},
		{"____", "____", "____", "____", "____"},
		{"____", "____", "____", "____", "____"},
	}
)

func InitializePRNG(seed string) []string {
	STATE := []string{
		"ПЕРВОЕ_АКТЕРСТВО",
		"ВТОРОЙ_ДАЛЬТОНИК",
		"ТРЕТЬЯ_САДОВНИЦА",
		"ЧЕТВЕРТЫЙ_ГОБЛИН",
	}

	Alphabet := alphabet.NewAlphabet(TelegraphAlphabet)
	SpongeBob := sponge.NewSponge(SpongeInnerState, *Alphabet)

	value := [4]string{}
	for i := 0; i < 4; i++ {
		value[i] = SpongeBob.CBlock([]string{STATE[i], seed}, 16)
	}
	secret := SpongeBob.CBlock(value[:], 16)

	out := [4]string{}
	for i := 0; i < 4; i++ {
		temp := value[i]
		TEMP := ""
		for j := 0; j < 4; j++ {
			temp = Alphabet.AddTxt(temp, STATE[i])
			TEMP += SpongeBob.CBlock([]string{temp, secret}, 4)
			temp = Alphabet.AddTxt(temp, TEMP)
		}
		out[i] = string([]rune(TEMP)[4:16])
	}
	return out[:]
}

type LFSR struct {
	Alphabet alphabet.Alphabet
}

func NewLFSR(Alphabet alphabet.Alphabet) *LFSR {
	return &LFSR{Alphabet}
}

func (l *LFSR) TapsToBin(taps_in []int) []int {
	taps := taps_in
	slices.Sort(taps)
	slices.Reverse(taps)

	temp := [20]int{}
	for _, tap := range taps {
		temp[tap-1] = 1
	}
	out := temp[:]
	slices.Reverse(out)
	return out
}

func (l *LFSR) LFSR_Push(state, taps []int) []int {
	N := min(len(state), len(taps))
	temp := 0
	for i := 0; i < N; i++ {
		temp += state[i] * taps[i]
	}
	return l.Alphabet.PushReg(state, temp%2)
}

func (l *LFSR) LFSR_Next(state, taps []int) [][]int {
	tempState := state
	stream := [20]int{}
	for i := 0; i < 20; i++ {
		tempState = l.LFSR_Push(tempState, taps)
		stream[i] = tempState[19]
	}
	out := [][]int{stream[:], tempState}
	return out
}

func (l *LFSR) SeedToBins(seeds []string) [][]int {
	var out [][]int
	for _, seed := range seeds {
		out = append(out, l.Alphabet.BlockToBin(seed))
	}
	return out
}

func (l *LFSR) ASLFSR_Push(state, taps [][]int) (int, [][]int) {
	lfsr0 := l.LFSR_Push(state[0], taps[0])
	lfsr1 := l.LFSR_Push(state[1], taps[1])
	lfsr2 := l.LFSR_Push(state[2], taps[2])

	var stream int

	if lfsr0[19] == 0 {
		stream = lfsr1[19]
	} else {
		stream = lfsr2[19]
	}

	return stream, [][]int{lfsr0, lfsr1, lfsr2}
}

func (l *LFSR) ASLFSR_Next(state, taps [][]int) ([]int, [][]int) {
	stateSet := state
	stream := [20]int{}
	for i := 0; i < 20; i++ {
		tempStream, tempState := l.ASLFSR_Push(stateSet, taps)
		stateSet = tempState
		stream[i] = tempStream
	}
	return stream[:], stateSet
}

func (l *LFSR) Wrap_CASLFSR_Next(initFlag string, stateIn [][][]int, seed string, tapsSet [][][]int) (string, [][][]int) {
	var state [][][]int
	var stream strings.Builder
	if initFlag == "up" {
		init := InitializePRNG(seed)
		for i := 0; i < 4; i++ {
			state = append(state, l.SeedToBins([]string{
				string([]rune(init[i])[:4]),
				string([]rune(init[i])[4:8]),
				string([]rune(init[i])[8:12]),
			}))
		}
	} else if initFlag == "down" {
		state = stateIn
	} else {
		panic("Wrap CASLFSR_Next: wrong flag")
	}

	for j := 0; j < 4; j++ {
		var temp []int
		for k := 0; k < 4; k++ {
			Tstream, Tstate := l.ASLFSR_Next(state[k], tapsSet[j])
			state[k] = Tstate
			if k == 0 {
				temp = Tstream
			} else {
				for i := 0; i < 20; i++ {
					temp[i] = (Tstream[i] + temp[i]) % 2
				}
			}
		}
		block := l.Alphabet.BinToBlock(temp)
		stream.Grow(len(block))
		stream.WriteString(block)
	}
	return stream.String(), state
}
