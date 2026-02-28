package main

import (
	"fmt"
	l3 "infbez_labs/Lab3"
	shared "infbez_labs/Shared"
)

func main() {
	var TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")
	var alphabet = *shared.NewAlphabet(TelegraphAlphabet)
	var LFSR = l3.NewLFSR(alphabet)

	SET := [][][]int{
		{
			LFSR.TapsToBin([]int{19, 18}),
			LFSR.TapsToBin([]int{18, 7}),
			LFSR.TapsToBin([]int{17, 3}),
		},
		{
			LFSR.TapsToBin([]int{19, 18}),
			LFSR.TapsToBin([]int{18, 7}),
			LFSR.TapsToBin([]int{16, 14, 13, 11}),
		},
		{
			LFSR.TapsToBin([]int{19, 18}),
			LFSR.TapsToBin([]int{18, 7}),
			LFSR.TapsToBin([]int{15, 13, 12, 10}),
		},
		{
			LFSR.TapsToBin([]int{19, 18}),
			LFSR.TapsToBin([]int{18, 7}),
			LFSR.TapsToBin([]int{14, 5, 3, 1}),
		},
	}

	seed := "АБВГДЕЖЗИЙКЛМНОП"
	out, intern := LFSR.Wrap_CASLFSR_Next("up", [][][]int{}, seed, SET)

	fmt.Println(out)

	for _, instate := range intern {
		for _, t := range instate {
			fmt.Println(t)
		}
	}

	IN1 := "ХОРОШО_БЫТЬ_ВАМИ"
	fmt.Println(l3.InitializePRNG(IN1))

	seed1 := "ЛЕРА"
	seed2 := "КЛОН"
	seed3 := "КОНЯ"

	S_set := [][]int{
		alphabet.BlockToBin(seed1),
		alphabet.BlockToBin(seed2),
		alphabet.BlockToBin(seed3),
	}

	T1 := []int{1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	T2 := []int{0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}
	T3 := []int{0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	T_set := [][]int{
		T1, T2, T3,
	}

	_, state := LFSR.ASLFSR_Next(S_set, T_set)
	fmt.Println(state)

	B1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	B2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}
	B3 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1}
	B4 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0}

	fmt.Println(alphabet.BinToBlock(B1))
	fmt.Println(alphabet.BinToBlock(B2))
	fmt.Println(alphabet.BinToBlock(B3))
	fmt.Println(alphabet.BinToBlock(B4))
}

func Display(inp [5][5]string) {
	for i := 0; i < 5; i++ {
		fmt.Println(inp[i])
	}
	fmt.Println("")
}
