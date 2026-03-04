package codeRandomGenerator_test

import (
	"fmt"
	alpha "infbez_labs/internal/alphabet"
	"reflect"

	generator "infbez_labs/internal/codeRandomGenerator"
	"testing"
)

var (
	telegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")
	alphabet          = alpha.NewAlphabet(telegraphAlphabet)
)

func TestAlphabet_Block2Num(t *testing.T) {

	var (
		Input1 = "АБВГ"
		Input2 = "_ЯЗЬ"
		Input3 = "ЯЯЯЯ"

		Output1 = 34916
		Output2 = 32028
		Output3 = 1048575
	)

	tests := []struct {
		name        string
		inputBlock  string
		outputBlock int
	}{
		{Input1 + "->" + fmt.Sprint(Output1), Input1, Output1},
		{Input2 + "->" + fmt.Sprint(Output2), Input2, Output2},
		{Input3 + "->" + fmt.Sprint(Output3), Input3, Output3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alphabet.BlockToNum(tt.inputBlock)

			if tt.outputBlock != got {
				t.Errorf("Failed Block2Num(input=%q), want %v but return %v", tt.inputBlock, tt.outputBlock, got)
				return
			}
		})
	}
}

func TestAlphabet_Num2Block(t *testing.T) {

	var (
		Output1 = "АБВГ"
		Output2 = "_ЯЗЬ"
		Output3 = "ЯЯЯЯ"

		Input1 = 34916
		Input2 = 32028
		Input3 = 1048575
	)

	tests := []struct {
		name        string
		inputBlock  int
		outputBlock string
	}{
		{fmt.Sprint(Input1) + "->" + Output1, Input1, Output1},
		{fmt.Sprint(Input2) + "->" + Output2, Input2, Output2},
		{fmt.Sprint(Input3) + "->" + Output3, Input3, Output3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alphabet.NumToBlock(tt.inputBlock)

			if tt.outputBlock != got {
				t.Errorf("Failed Num2Block(input=%q), want %v but return %v", tt.inputBlock, tt.outputBlock, got)
				return
			}
		})
	}
}

func TestAlphabet_Dec2Bin(t *testing.T) {

	var (
		Input1 = 34916
		Input2 = 32028
		Input3 = 1048575

		Output1 = []int{0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0}
		Output2 = []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0}
		Output3 = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	)

	tests := []struct {
		name        string
		inputBlock  int
		outputBlock []int
	}{
		{fmt.Sprint(Input1) + "->" + fmt.Sprint(Output1), Input1, Output1},
		{fmt.Sprint(Input2) + "->" + fmt.Sprint(Output2), Input2, Output2},
		{fmt.Sprint(Input3) + "->" + fmt.Sprint(Output3), Input3, Output3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alphabet.DecToBin(tt.inputBlock)

			if !reflect.DeepEqual(tt.outputBlock, got) {
				t.Errorf("Failed Dec2Bin(input=%q), want %v but return %v", tt.inputBlock, tt.outputBlock, got)
				return
			}
		})
	}
}

func TestAlphabet_Bin2Dec(t *testing.T) {

	var (
		Input1 = []int{0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0}
		Input2 = []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0}
		Input3 = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

		Output1 = 34916
		Output2 = 32028
		Output3 = 1048575
	)

	tests := []struct {
		name        string
		inputBlock  []int
		outputBlock int
	}{
		{fmt.Sprint(Input1) + "->" + fmt.Sprint(Output1), Input1, Output1},
		{fmt.Sprint(Input2) + "->" + fmt.Sprint(Output2), Input2, Output2},
		{fmt.Sprint(Input3) + "->" + fmt.Sprint(Output3), Input3, Output3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alphabet.BinToDec(tt.inputBlock)

			if !reflect.DeepEqual(tt.outputBlock, got) {
				t.Errorf("Failed Bin2Dec(input=%q), want %v but return %v", tt.inputBlock, tt.outputBlock, got)
				return
			}
		})
	}
}

func TestAlphabet_Block2Bin(t *testing.T) {

	var (
		Input1 = "____"
		Input2 = "___А"
		Input3 = "__Б_"
		Input4 = "__БГ"

		Output1 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		Output2 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
		Output3 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}
		Output4 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0}
	)

	tests := []struct {
		name        string
		inputBlock  string
		outputBlock []int
	}{
		{fmt.Sprint(Input1) + "->" + fmt.Sprint(Output1), Input1, Output1},
		{fmt.Sprint(Input2) + "->" + fmt.Sprint(Output2), Input2, Output2},
		{fmt.Sprint(Input3) + "->" + fmt.Sprint(Output3), Input3, Output3},
		{fmt.Sprint(Input4) + "->" + fmt.Sprint(Output4), Input4, Output4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alphabet.BlockToBin(tt.inputBlock)

			if !reflect.DeepEqual(tt.outputBlock, got) {
				t.Errorf("Failed Block2Bin(input=%q), want %v but return %v", tt.inputBlock, tt.outputBlock, got)
				return
			}
		})
	}
}

func TestAlphabet_Bin2Block(t *testing.T) {

	var (
		Input1 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		Input2 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
		Input3 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}
		Input4 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0}

		Output1 = "____"
		Output2 = "___А"
		Output3 = "__Б_"
		Output4 = "__БГ"
	)

	tests := []struct {
		name        string
		inputBlock  []int
		outputBlock string
	}{
		{fmt.Sprint(Input1) + "->" + fmt.Sprint(Output1), Input1, Output1},
		{fmt.Sprint(Input2) + "->" + fmt.Sprint(Output2), Input2, Output2},
		{fmt.Sprint(Input3) + "->" + fmt.Sprint(Output3), Input3, Output3},
		{fmt.Sprint(Input4) + "->" + fmt.Sprint(Output4), Input4, Output4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alphabet.BinToBlock(tt.inputBlock)

			if !reflect.DeepEqual(tt.outputBlock, got) {
				t.Errorf("Failed Bin2Block((input=%q), want %v but return %v", tt.inputBlock, tt.outputBlock, got)
				return
			}
		})
	}
}

func TestAlphabet_PushReg(t *testing.T) {

	var (
		Input1 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		Input2 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
		Input3 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}
		Input4 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0}

		Output1 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
		Output2 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}
		Output3 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1}
		Output4 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0}
	)

	tests := []struct {
		name        string
		inputBlock  []int
		bool        int
		outputBlock []int
	}{
		{fmt.Sprint(Input1) + "->" + fmt.Sprint(Output1), Input1, 1, Output1},
		{fmt.Sprint(Input2) + "->" + fmt.Sprint(Output2), Input2, 0, Output2},
		{fmt.Sprint(Input3) + "->" + fmt.Sprint(Output3), Input3, 1, Output3},
		{fmt.Sprint(Input4) + "->" + fmt.Sprint(Output4), Input4, 0, Output4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alphabet.PushReg(tt.inputBlock, tt.bool)

			if !reflect.DeepEqual(tt.outputBlock, got) {
				t.Errorf("Failed Bin2Block((input=%q, bool_in=%q), want %v but return %v", tt.inputBlock, tt.bool, tt.outputBlock, got)
				return
			}
		})
	}
}

func TestPRNG_InitializePRNG(t *testing.T) {

	//var

	var (
		Input1 = "ХОРОШО_БЫТЬ_ВАМИ"
		Input2 = "________________"
		Input3 = "___А____________"
		Input4 = "ХОРОШО_ВЫТЬ_ВАМИ"

		Output1 = []string{
			"ФЕШЗЕЧЧ_ВЕГГ",
			"ОГЛНЖРСЙФДПХ",
			"КЖЮОХЛТКЬЕХЩ",
			"ЛШЧЫМССЭМЗФЕ",
		}
		Output2 = []string{
			"ЫНЖГИТЛМРШЗА",
			"ЧБСЗЮИФШЙДЙШ",
			"ЕУШЮЙАЭЩЖЗБА",
			"ДБЧЯЧРДВЩ_ОК",
		}
		Output3 = []string{
			"ПСЛГ_ЕКНБШЧУ",
			"ШИС_АЙПРЭФОФ",
			"ЫБТЯЧМЙЮШЕУР",
			"ДФДЭИДИИЙМЬН",
		}
		Output4 = []string{
			"НЫЖПИЛЗУТТФЦ",
			"АЗРВВШВРВЙЫШ",
			"ПЦОЖБЙЮЯРЙРФ",
			"ДОЖЭ_Я_ЙВЬУК",
		}
	)

	tests := []struct {
		name        string
		inputBlock  string
		outputBlock []string
	}{
		{fmt.Sprint(Input1) + "->" + fmt.Sprint(Output1), Input1, Output1},
		{fmt.Sprint(Input2) + "->" + fmt.Sprint(Output2), Input2, Output2},
		{fmt.Sprint(Input3) + "->" + fmt.Sprint(Output3), Input3, Output3},
		{fmt.Sprint(Input3) + "->" + fmt.Sprint(Output4), Input4, Output4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generator.InitializePRNG(tt.inputBlock)

			if !reflect.DeepEqual(tt.outputBlock, got) {
				t.Errorf("Failed InitializePRNG(input=%q), want %v but return %v", tt.inputBlock, tt.outputBlock, got)
				return
			}
		})
	}
}

func TestPRNG_TapsToBin(t *testing.T) {

	var (
		gen = generator.LFSR{Alphabet: *alphabet}

		Input1 = []int{20, 17}
		Input2 = []int{19, 18, 17, 4}
		Input3 = []int{18, 11}
		Input4 = []int{20, 19, 4, 3}
		Input5 = []int{19, 18, 17, 13}
		Input6 = []int{18, 17, 16, 13}

		Output1 = []int{1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		Output2 = []int{0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}
		Output3 = []int{0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		Output4 = []int{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0}
		Output5 = []int{0, 1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		Output6 = []int{0, 0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	)

	tests := []struct {
		name        string
		inputBlock  []int
		outputBlock []int
	}{
		{fmt.Sprint(Input1) + "->" + fmt.Sprint(Output1), Input1, Output1},
		{fmt.Sprint(Input2) + "->" + fmt.Sprint(Output2), Input2, Output2},
		{fmt.Sprint(Input3) + "->" + fmt.Sprint(Output3), Input3, Output3},
		{fmt.Sprint(Input4) + "->" + fmt.Sprint(Output4), Input4, Output4},
		{fmt.Sprint(Input5) + "->" + fmt.Sprint(Output5), Input5, Output5},
		{fmt.Sprint(Input6) + "->" + fmt.Sprint(Output6), Input6, Output6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gen.TapsToBin(tt.inputBlock)

			if !reflect.DeepEqual(tt.outputBlock, got) {
				t.Errorf("Failed TapsToBin((input=%q), want %v but return %v", tt.inputBlock, tt.outputBlock, got)
				return
			}
		})
	}
}

func TestPRNG_LFSR_Push(t *testing.T) {

	var (
		lfsr = generator.NewLFSR(*alphabet)
		S0   = []int{1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1}

		S1 = []int{0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0}
		S2 = []int{1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1}
		S3 = []int{1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1}
		S4 = []int{1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0}
		S5 = []int{0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 1}
		S6 = []int{1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 1, 0}
		S7 = []int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 1, 0, 1}
		S8 = []int{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0}
		S9 = []int{0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 0}

		T1 = []int{1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	)

	tests := []struct {
		name        string
		inputBlock  []int
		outputBlock []int
	}{
		{fmt.Sprint(S0) + "->" + fmt.Sprint(S1), S0, S1},
		{fmt.Sprint(S1) + "->" + fmt.Sprint(S2), S1, S2},
		{fmt.Sprint(S2) + "->" + fmt.Sprint(S3), S2, S3},
		{fmt.Sprint(S3) + "->" + fmt.Sprint(S4), S3, S4},
		{fmt.Sprint(S4) + "->" + fmt.Sprint(S5), S4, S5},
		{fmt.Sprint(S5) + "->" + fmt.Sprint(S6), S5, S6},
		{fmt.Sprint(S6) + "->" + fmt.Sprint(S7), S6, S7},
		{fmt.Sprint(S7) + "->" + fmt.Sprint(S8), S7, S8},
		{fmt.Sprint(S8) + "->" + fmt.Sprint(S9), S8, S9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lfsr.LFSR_Push(tt.inputBlock, T1)

			if !reflect.DeepEqual(tt.outputBlock, got) {
				t.Errorf("Failed LFSR_Push((input=%q, taps=%q), want %v but return %v", tt.inputBlock, T1, tt.outputBlock, got)
				return
			}
		})
	}
}

func TestPRNG_LFSR_Next(t *testing.T) {

	var (
		lfsr = generator.NewLFSR(*alphabet)

		seed = alphabet.BlockToBin("ОРИМ")
		T1   = []int{1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		T2   = []int{0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}

		tmp1 = lfsr.LFSR_Next(seed, T1)
		tmp2 = lfsr.LFSR_Next(seed, T2)

		seq1 = [10][]int{tmp1[0]}
		seq2 = [10][]int{tmp2[0]}

		Out1 = []string{"МВСЙ", "ДЯ_Ы", "ЙЖЕА", "ЫЮХО", "ГКДХ", "ЕСЗЖ", "С_ИЫ", "ББЖВ", "СТЯЯ"}
		Out2 = []string{"ЙБАА", "Н_НР", "ЦЛКЯ", "ЦДМЕ", "ЖЫПД", "ГМЩВ", "ЭШЯИ", "РЧИК", "ЭИОЧ"}
	)

	tests := []struct {
		name        string
		inputBlock  [10][]int
		T           []int
		outputBlock []string
	}{
		{"LFSR_Next_1", seq1, T1, Out1},
		{"LFSR_Next_2", seq2, T2, Out2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lfsr = generator.NewLFSR(*alphabet)
			for i := 1; i < 10; i++ {
				tt.inputBlock[i] = lfsr.LFSR_Next(tt.inputBlock[i-1], tt.T)[1]
			}
			var got []string
			for i := 1; i < 10; i++ {
				got = append(got, alphabet.BinToBlock(tt.inputBlock[i]))
			}

			if !reflect.DeepEqual(tt.outputBlock, got) {
				t.Errorf("Failed LFSR_Next((input=%q, taps=%q), want %v but return %v", tt.inputBlock, T1, tt.outputBlock, got)
				return
			}
		})
	}
}

func TestPRNG_AS_LFSR_Push(t *testing.T) {

	var (
		lfsr = generator.NewLFSR(*alphabet)

		seed1 = alphabet.BlockToBin("ЛЕРА")
		seed2 = alphabet.BlockToBin("КЛОН")
		seed3 = alphabet.BlockToBin("КОНЯ")

		T1 = []int{1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		T2 = []int{0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}
		T3 = []int{0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

		T_set = [][]int{T1, T2, T3}
		S_set = [][]int{seed1, seed2, seed3}

		Out1 = []int{0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1}
		Out2 = []int{1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1}
		Out3 = []int{1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0}

		Out_set = [][]int{Out1, Out2, Out3}
	)

	t.Run("AS_LFSR_Push", func(t *testing.T) {
		lfsr = generator.NewLFSR(*alphabet)
		kod1, got1 := lfsr.ASLFSR_Push(S_set, T_set)
		kod2, got2 := lfsr.ASLFSR_Push(got1, T_set)
		kod3, got3 := lfsr.ASLFSR_Push(got2, T_set)

		if kod1 != 1 || kod2 != 1 || kod3 != 0 {
			t.Errorf("Failed AS_LFSR_Push((input=%q, input2=%q), want %v but return %v", S_set, T_set, Out_set, got3)
			return
		}

		if !reflect.DeepEqual(Out_set, got3) {
			t.Errorf("Failed AS_LFSR_Push((input=%q, input2=%q), want %v but return %v", S_set, T_set, Out_set, got3)
			return
		}
	})
}

func TestPRNG_AS_LFSR_Next(t *testing.T) {

	var (
		lfsr = generator.NewLFSR(*alphabet)

		seed1 = alphabet.BlockToBin("ЛЕРА")
		seed2 = alphabet.BlockToBin("КЛОН")
		seed3 = alphabet.BlockToBin("КОНЯ")

		T1 = []int{1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		T2 = []int{0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}
		T3 = []int{0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

		TSet = [][]int{T1, T2, T3}
		SSet = [][]int{seed1, seed2, seed3}

		OutSet = []string{
			"ОРЩИ",
			"ЙАУС",
			"_ШИИ",
			"ЦЕЖГ",
			"БЬМЗ",
			"ОЯРЙ",
			"ДЯТЩ",
			"ИСЭП",
			"РЧУШ",
		}
	)

	t.Run("AS_LFSR_Push", func(t *testing.T) {
		lfsr = generator.NewLFSR(*alphabet)
		streams := [10][]int{}
		stateSets := [10][][]int{}
		streams[0], stateSets[0] = lfsr.ASLFSR_Next(SSet, TSet)

		for i := 1; i < 10; i++ {
			streams[i], stateSets[i] = lfsr.ASLFSR_Next(stateSets[i-1], TSet)
		}

		got := make([]string, 9)
		for i := 1; i < 10; i++ {
			got[i-1] = alphabet.BinToBlock(streams[i])
		}

		if !reflect.DeepEqual(OutSet, got) {
			t.Errorf("Failed AS_LFSR_Next((input=%q, input2=%q), want %v but return %v", SSet, TSet, OutSet, got)
			return
		}
	})
}

func TestPRNG_C_AS_LFSR_Next(t *testing.T) {

	var (
		lfsr = generator.NewLFSR(*alphabet)

		SET = [][][]int{
			{
				lfsr.TapsToBin([]int{19, 18}),
				lfsr.TapsToBin([]int{18, 7}),
				lfsr.TapsToBin([]int{17, 3}),
			},
			{
				lfsr.TapsToBin([]int{19, 18}),
				lfsr.TapsToBin([]int{18, 7}),
				lfsr.TapsToBin([]int{16, 14, 13, 11}),
			},
			{
				lfsr.TapsToBin([]int{19, 18}),
				lfsr.TapsToBin([]int{18, 7}),
				lfsr.TapsToBin([]int{15, 13, 12, 10}),
			},
			{
				lfsr.TapsToBin([]int{19, 18}),
				lfsr.TapsToBin([]int{18, 7}),
				lfsr.TapsToBin([]int{14, 5, 3, 1}),
			},
		}

		seed = "АБВГДЕЖЗИЙКЛМНОП"

		Out    = make([]string, 9)
		Intern = make([][][][]int, 9)

		Output = []string{
			"ЙЮЬХЖХРЬБЦБЬЧРЫ_",
			"ЧЖЙЕЭНЭДГХЛЛИЧЛЛ",
			"ПЙТН_ПХВЕФШИТ_СП",
			"ЙЯЧСНБКГЦЩЦФЮ_БЦ",
			"ВНРЬЛНАМЯХОЕКЕАЭ",
			"ЫММЮЗЯБМЩЙХИБЧИО",
			"ЯБ_ЭТЩНЙЦЫАМКЛВО",
			"ЬАЭАВЩЛОЖЭВЖЬЙНЬ",
			"ГЛЮШРЙЙЕТЖЗЦБКЫА",
		}
	)

	t.Run("C_AS_LFSR_Next", func(t *testing.T) {
		Out[0], Intern[0] = lfsr.WrapCAsLfsrNext("up", [][][]int{}, seed, SET)

		for i := 1; i < 9; i++ {
			Out[i], Intern[i] = lfsr.WrapCAsLfsrNext("down", Intern[i-1], seed, SET)
		}

		if !reflect.DeepEqual(Output, Out) {
			t.Errorf("Failed C_AS_LFSR_Next, want %v but return %v", Output, Out)
			return
		}
	})
}
