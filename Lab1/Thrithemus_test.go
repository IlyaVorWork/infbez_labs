package trithemius_test

import (
	"fmt"
	l1 "infbez_labs/Lab1"
	"slices"
	"testing"
)

var TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")

func TestTrithemius_BuildTrithemiusAlphabet(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		K1 = "ДИНОЗАВР_ЗАУРОПОД"
		K2 = "ГАМЕЛЬНСКИЙ_АНТИКВАР"
		K3 = "ГАРРИ_ПОТТЕР_И_ФИЛОСОФСКИЙ_КАМЕНЬ"

		OutputTable1 = "ДИНОЗАВР_ЙБУСПТФЕГЖКЛМХЦЧШЩЫЬЭЮЯ"
		OutputTable2 = "ГАМЕЛЬНСКИЙ_БОТПРВДУЖЗФХЦЧШЩЫЭЮЯ"
		OutputTable3 = "ГАРСИ_ПОТУЕФБЙВХКЛЦЧШЩЫМНЬДЭЖЮЗЯ"
	)

	tests := []struct {
		name        string
		inputTable  string
		outputTable string
	}{
		{K1 + "->" + OutputTable1, K1, OutputTable1},
		{K2 + "->" + OutputTable2, K2, OutputTable2},
		{K3 + "->" + OutputTable3, K3, OutputTable3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(trithemius.BuildTrithemiusAlphabet(tt.inputTable))

			if tt.outputTable != got {
				t.Errorf("Failed BuildTrithemiusAlphabet(input=%q), want %v but return %v", tt.inputTable, tt.outputTable, got)
				return
			}

		})
	}
}

func TestTrithemius_GetCharByKey(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		Key1 = 0
		Key2 = 1
		Key3 = 7
		Key4 = 14
		Key5 = 32

		OutputChar1 = "_"
		OutputChar2 = "А"
		OutputChar3 = "Ж"
		OutputChar4 = "Н"
		OutputChar5 = "_"
	)

	tests := []struct {
		name       string
		inputKey   int
		outputChar string
	}{
		{fmt.Sprint(Key1) + "->" + OutputChar1, Key1, OutputChar1},
		{fmt.Sprint(Key2) + "->" + OutputChar2, Key2, OutputChar2},
		{fmt.Sprint(Key3) + "->" + OutputChar3, Key3, OutputChar3},
		{fmt.Sprint(Key4) + "->" + OutputChar4, Key4, OutputChar4},
		{fmt.Sprint(Key5) + "->" + OutputChar5, Key5, OutputChar5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trithemius.Alphabet.GetCharByKey(tt.inputKey); tt.outputChar != got {
				t.Errorf("Failed GetCharByKey(Key=%v), want %v but return %v", tt.inputKey, tt.outputChar, got)
				return
			}

		})
	}
}

func TestTrithemius_GetKeyByChar(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		InputChar1 = "_"
		InputChar2 = "А"
		InputChar3 = "Ж"
		InputChar4 = "Н"
		InputChar5 = "Е"

		OutputKey1 = 0
		OutputKey2 = 1
		OutputKey3 = 7
		OutputKey4 = 14
		OutputKey5 = 6
	)

	tests := []struct {
		name      string
		inputChar string
		OutputKey int
	}{
		{InputChar1 + "->" + fmt.Sprint(OutputKey1), InputChar1, OutputKey1},
		{InputChar2 + "->" + fmt.Sprint(OutputKey2), InputChar2, OutputKey2},
		{InputChar3 + "->" + fmt.Sprint(OutputKey3), InputChar3, OutputKey3},
		{InputChar4 + "->" + fmt.Sprint(OutputKey4), InputChar4, OutputKey4},
		{InputChar5 + "->" + fmt.Sprint(OutputKey5), InputChar5, OutputKey5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trithemius.Alphabet.GetKeyByChar(tt.inputChar); tt.OutputKey != got {
				t.Errorf("Failed GetKeyByChar(Char=%v), want %v but return %v", tt.inputChar, tt.OutputKey, got)
				return
			}
		})
	}
}

func TestTrithemius_AddTxt(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)
	var (
		TT1 = "ЕЖИК"
		TT2 = "В_ТУМАНЕ"
		TT3 = "БАРОН"
		TT4 = "ВАРАН"
		TT5 = "ИЖЬЯМАНЕ"
		TT6 = "Я__Н_"
	)

	tests := []struct {
		name   string
		input1 string
		input2 string
		Output string
	}{
		{TT1 + "||" + TT2 + "->" + TT5, TT1, TT2, TT5},
		{TT6 + "||" + TT4 + "->" + TT3, TT6, TT4, TT3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trithemius.Alphabet.AddTxt(tt.input1, tt.input2); tt.Output != got {
				t.Errorf("Failed AddTxt(Text1=%v, Text2=%v), want %v but return %v", tt.input1, tt.input2, tt.Output, got)
				return
			}
		})
	}
}

func TestTrithemius_SubTxt(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)
	var (
		TT1 = "ЕЖИК"
		TT2 = "В_ТУМАНЕ"
		TT3 = "БАРОН"
		TT4 = "ВАРАН"
		TT5 = "ИЖЬЯМАНЕ"
		TT6 = "Я__Н_"
	)

	tests := []struct {
		name   string
		input1 string
		input2 string
		Output string
	}{
		{TT3 + "||" + TT4 + "->" + TT6, TT3, TT4, TT6},
		{TT5 + "||" + TT2 + "->" + "ЕЖИК____", TT5, TT2, "ЕЖИК____"},
		{TT5 + "||" + TT1 + "->" + TT2, TT5, TT1, TT2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trithemius.Alphabet.SubTxt(tt.input1, tt.input2); tt.Output != got {
				t.Errorf("Failed SubTxt(Text1=%v, Text2=%v), want %v but return %v", tt.input1, tt.input2, tt.Output, got)
				return
			}
		})
	}
}

func TestTrithemius_ArrayToText(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		IN1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 0}
		IN2 = []int{0, 0, 0, 0, 0}
		IN3 = []int{1, 4, 4}
		IN4 = []int{10, 23, 20, 11, 6, 14}
		IN5 = []int{13, 9, 17}
	)

	var (
		OUT1 = "АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_"
		OUT2 = "_____"
		OUT3 = "АГГ"
		OUT4 = "ЙЦУКЕН"
		OUT5 = "МИР"
	)

	tests := []struct {
		name       string
		inputArray []int
		outputText string
	}{
		{fmt.Sprint(IN1) + "->" + OUT1, IN1, OUT1},
		{fmt.Sprint(IN2) + "->" + OUT2, IN2, OUT2},
		{fmt.Sprint(IN3) + "->" + OUT3, IN3, OUT3},
		{fmt.Sprint(IN4) + "->" + OUT4, IN4, OUT4},
		{fmt.Sprint(IN5) + "->" + OUT5, IN5, OUT5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trithemius.Alphabet.ArrayToText(tt.inputArray)

			if got != tt.outputText {
				t.Errorf("ArrayToText(array=%q), want %q but return %q", tt.inputArray, tt.outputText, got)
				return
			}
		})
	}
}

func TestTrithemius_TextToArray(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		IN1 = "АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_"
		IN2 = "_____"
		IN3 = "АГГ"
		IN4 = "ЙЦУКЕН"
		IN5 = "МИР"
	)

	var (
		OUT1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 0}
		OUT2 = []int{0, 0, 0, 0, 0}
		OUT3 = []int{1, 4, 4}
		OUT4 = []int{10, 23, 20, 11, 6, 14}
		OUT5 = []int{13, 9, 17}
	)

	tests := []struct {
		name        string
		inputText   string
		outputArray []int
	}{
		{IN1 + "->" + fmt.Sprint(OUT1), IN1, OUT1},
		{IN2 + "->" + fmt.Sprint(OUT2), IN2, OUT2},
		{IN3 + "->" + fmt.Sprint(OUT3), IN3, OUT3},
		{IN4 + "->" + fmt.Sprint(OUT4), IN4, OUT4},
		{IN5 + "->" + fmt.Sprint(OUT5), IN5, OUT5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trithemius.Alphabet.TextToArray(tt.inputText)

			if !slices.Equal(got, tt.outputArray) {
				t.Errorf("TextToArray(text=%q), want %q but return %q", tt.inputText, tt.outputArray, got)
				return
			}
		})
	}
}

func TestTrithemius_EncodeTrithemius_BasicTests(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		IN = "ГОЛОВНОЙ_ОФИС"
	)

	var (
		KEY1 = "ЧЕРНОСОТЕНЦЫ"
		KEY2 = "АБВГД"
		KEY3 = "А"
	)

	var (
		OUT1 = "ФАЮАМЫАЬТА_ЩБ"
		OUT2 = "ЛЦУЦКХЦСЗЦЭРЩ"
	)

	tests := []struct {
		name       string
		openText   string
		key        string
		cipherText string
	}{
		{IN + "||" + KEY1 + "->" + OUT1, IN, KEY1, OUT1},
		{IN + "||" + KEY2 + "->" + OUT2, IN, KEY2, OUT2},
		{IN + "||" + KEY3 + "->" + OUT2, IN, KEY3, OUT2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table := trithemius.BuildTrithemiusAlphabet(tt.key)
			got := trithemius.EncodeTrithemius(tt.openText, table)

			if got != tt.cipherText {
				t.Errorf("EncodeThrithemus(text=%q , key=%q), want %q but return %q", tt.openText, tt.key, tt.cipherText, got)
				return
			}

		})
	}
}

func TestTrithemius_DecodeTrithemius_BasicTests(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		IN1 = "ФАЮАМЫАЬТА_ЩБ"
		IN2 = "ЛЦУЦКХЦСЗЦЭРЩ"
	)

	var (
		KEY1 = "ЧЕРНОСОТЕНЦЫ"
		KEY2 = "АБВГД"
		KEY3 = "А"
	)

	var (
		OUT = "ГОЛОВНОЙ_ОФИС"
	)

	tests := []struct {
		name       string
		openText   string
		key        string
		cipherText string
	}{
		{IN1 + "||" + KEY1 + "->" + OUT, IN1, KEY1, OUT},
		{IN2 + "||" + KEY2 + "->" + OUT, IN2, KEY2, OUT},
		{IN2 + "||" + KEY3 + "->" + OUT, IN2, KEY3, OUT},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table := trithemius.BuildTrithemiusAlphabet(tt.key)
			got := trithemius.DecodeTrithemius(tt.openText, table)

			if got != tt.cipherText {
				t.Errorf("DecodeThrithemus(text=%q , key=%q), want %q but return %q", tt.openText, tt.key, tt.cipherText, got)
				return
			}

		})
	}
}

func TestTrithemius_ShiftTrithemiusAlphabet(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		table = trithemius.BuildTrithemiusAlphabet("ХОРОШО_БЫТЬ_ВАМИ")

		SymIn1 = "Х"
		SymIn2 = "Я"
		SymIn3 = "А"

		BIAS1 = 16
		BIAS2 = 18
		BIAS3 = 8
		BIAS4 = 23
		BIAS5 = 30
		BIAS6 = 1

		K1 = "ЦХОРПШС_БЫТЬАВГМИДЕЖЗЙКЛНУФЧЩЭЮЯ"
		K2 = "ЦХОРПШС_БЫТЬАВГМИДЕЖЗЙКЛНУФЧЩЭЮЯ"
		K3 = "ЯХОРПШС_БЫТЬАВГМИДЕЖЗЙКЛНУФЦЧЩЭЮ"
		K4 = "АХОРПШС_БЫТЬВГМИДЕЖЗЙКЛНУФЦЧЩЭЮЯ"
		K5 = "НХОРПШС_БЫТЬАВГМИДЕЖЗЙКЛУФЦЧЩЭЮЯ"
		K6 = "ЮХОРПШС_БЫТЬАВГМИДЕЖЗЙКЛНУФЦЧЩЭЯ"
		K7 = "АХОРПШС_БЫТЬВГМИДЕЖЗЙКЛНУФЦЧЩЭЮЯ"
	)

	tests := []struct {
		name        string
		inputSym    string
		inputBias   int
		outputTable string
	}{
		{SymIn1 + "||" + string(rune(BIAS1)) + "->" + K1, SymIn1, BIAS1, K1},
		{SymIn1 + "||" + string(rune(BIAS2)) + "->" + K2, SymIn1, BIAS2, K2},
		{SymIn2 + "||" + string(rune(BIAS1)) + "->" + K3, SymIn2, BIAS1, K3},
		{SymIn3 + "||" + string(rune(BIAS3)) + "->" + K4, SymIn3, BIAS3, K4},
		{SymIn3 + "||" + string(rune(BIAS4)) + "->" + K5, SymIn3, BIAS4, K5},
		{SymIn3 + "||" + string(rune(BIAS5)) + "->" + K6, SymIn3, BIAS5, K6},
		{SymIn3 + "||" + string(rune(BIAS6)) + "->" + K7, SymIn3, BIAS6, K7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(trithemius.ShiftTrithemiusAlphabet(table, tt.inputSym, tt.inputBias))
			if tt.outputTable != got {
				t.Errorf("Failed ShiftTrithemiusAlphabet(sym=%v, bias%v), want %v but return %v", tt.inputSym, tt.inputBias, tt.outputTable, got)
				return
			}
		})
	}
}

func TestTrithemius_EncodePolyTrithemius(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		IN1 = "ОТКРЫТЫЙ_ТЕКСТ"
		IN2 = "ДЛИННЫЙ_ОТКРЫТЫЙ_ТЕКСТ"
	)

	var (
		KEY1 = "АББАТ_ТРИТИМУС"
		KEY2 = "БАБАТ_ТРИТИМУС"
		KEY3 = "ТРИТИМУС_АББАТ"
	)

	var (
		OUT1 = "ЭХЩКДХЖШСХУВНХ"
		OUT2 = "ЭХЩКДХЖШСХУВНХ"
		OUT3 = "ПЫКЬЬЕЧСШХВПЛХЦЬСХУВЫН"
		OUT4 = "ПЬБЮЯЧГЖН_МАТ_ТГЮ_ЧФИВ"
	)

	tests := []struct {
		name       string
		openText   string
		key        string
		cipherText string
	}{
		{IN1 + "||" + KEY1 + "->" + OUT1, IN1, KEY1, OUT1},
		{IN1 + "||" + KEY2 + "->" + OUT2, IN1, KEY2, OUT2},
		{IN2 + "||" + KEY1 + "->" + OUT3, IN2, KEY1, OUT3},
		{IN2 + "||" + KEY3 + "->" + OUT4, IN2, KEY3, OUT4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trithemius.EncodePolyTrithemius(tt.openText, tt.key)

			if got != tt.cipherText {
				t.Errorf("EncodePolyTrithemius(text=%q , key=%q), want %q but return %q", tt.openText, tt.key, tt.cipherText, got)
				return
			}

		})
	}
}

func TestTrithemius_DecodePolyTrithemius_BasicTests(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		In1 = "ЭХЩКДХЖШСХУВНХ"
		In2 = "ЭХЩКДХЖШСХУВНХ"
		In3 = "ПЫКЬЬЕЧСШХВПЛХЦЬСХУВЫН"
		In4 = "ПЬБЮЯЧГЖН_МАТ_ТГЮ_ЧФИВ"
	)

	var (
		Key1 = "АББАТ_ТРИТИМУС"
		Key2 = "БАБАТ_ТРИТИМУС"
		Key3 = "ТРИТИМУС_АББАТ"
	)

	var (
		Out1 = "ОТКРЫТЫЙ_ТЕКСТ"
		Out2 = "ДЛИННЫЙ_ОТКРЫТЫЙ_ТЕКСТ"
		Out3 = "ДКБЗЖСЬЦЬМУ_ЯМДБЦМЛУЮЬ"
	)

	tests := []struct {
		name       string
		cipherText string
		key        string
		openText   string
	}{
		{In1 + "||" + Key1 + "->" + Out1, In1, Key1, Out1},
		{In2 + "||" + Key2 + "->" + Out1, In2, Key2, Out1},
		{In2 + "||" + Key1 + "->" + Out1, In2, Key1, Out1},
		{In3 + "||" + Key3 + "->" + Out3, In3, Key3, Out3},
		{In4 + "||" + Key3 + "->" + Out2, In4, Key3, Out2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trithemius.DecodePolyTrithemius(tt.cipherText, tt.key)

			if got != tt.openText {
				t.Errorf("DecodePolyTrithemius(text=%q , key=%q), want %q but return %q", tt.cipherText, tt.key, tt.openText, got)
				return
			}
		})
	}
}

func TestTrithemius_EncodeSTrithemius_BasicTests(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		IN1 = "БЛОК"
		IN2 = "БРОК"
	)

	var (
		KEY1 = "НЕТ_ЗВЕЗД_В_НОЧИ"
		KEY2 = "ХОРОШО_БЫТЬ_ВАМИ"
	)

	var (
		OUT1 = "РЩФЖ"
		OUT2 = "РИФЖ"
		OUT3 = "ИЯТФ"
		OUT4 = "ИЬТФ"
	)

	tests := []struct {
		name       string
		openText   string
		key        string
		cipherText string
	}{
		{IN1 + "||" + KEY1 + "->" + OUT1, IN1, KEY1, OUT1},
		{IN2 + "||" + KEY1 + "->" + OUT2, IN2, KEY1, OUT2},
		{IN1 + "||" + KEY2 + "->" + OUT3, IN1, KEY2, OUT3},
		{IN2 + "||" + KEY2 + "->" + OUT4, IN2, KEY2, OUT4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trithemius.EncodeSTrithemius(tt.openText, tt.key)

			if got != tt.cipherText {
				t.Errorf("SThrithemus(text=%q , key=%q), want %q but return %q", tt.openText, tt.key, tt.cipherText, got)
				return
			}

		})
	}
}

func TestTrithemius_DecodeSTrithemius_BasicTests(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		In1 = "БЛОК"
		In2 = "БРОК"
	)

	var (
		Key1 = "НЕТ_ЗВЕЗД_В_НОЧИ"
		Key2 = "ХОРОШО_БЫТЬ_ВАМИ"
	)

	var (
		Out1 = "РЩФЖ"
		Out2 = "РИФЖ"
		Out3 = "ИЯТФ"
		Out4 = "ИЬТФ"
	)

	tests := []struct {
		name       string
		cipherText string
		key        string
		openText   string
	}{
		{Out1 + "||" + Key1 + "->" + In1, Out1, Key1, In1},
		{Out1 + "||" + Key2 + "->" + "ЦЗДЬ", Out1, Key2, "ЦЗДЬ"},
		{Out2 + "||" + Key1 + "->" + In2, Out2, Key1, In2},
		{Out2 + "||" + Key2 + "->" + "ЦБДЬ", Out2, Key2, "ЦБДЬ"},
		{Out3 + "||" + Key1 + "->" + "ЯХЬБ", Out3, Key1, "ЯХЬБ"},
		{Out3 + "||" + Key2 + "->" + In1, Out3, Key2, In1},
		{Out4 + "||" + Key1 + "->" + "ЯСЬБ", Out4, Key1, "ЯСЬБ"},
		{Out4 + "||" + Key2 + "->" + In2, Out4, Key2, In2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trithemius.DecodeSTrithemius(tt.cipherText, tt.key)

			if got != tt.openText {
				t.Errorf("STrithemius(cipherText=%q , key=%q), want %q but return %q", tt.cipherText, tt.key, tt.openText, got)
				return
			}
		})
	}
}

func TestTrithemius_EncodeMergeBlock(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		In1 = "БЛОК"
		In2 = "БРОК"
	)

	var (
		Key1 = "ХОРОШО_ВЫТЬ_ВАМИ"
		Key2 = "ХОРОШО_БЫТЬ_ВАМИ"
		// Key3 = "ХОРОШО_ВЫТЬ_БАМИ"
	)

	var (
		Out1 = "ЬЗЦЩ"
		Out2 = "ЬМЬЩ"
		Out3 = "МЗЬТ"
		Out4 = "ММЬЧ"
	)

	tests := []struct {
		name    string
		intText string
		key     string
		outText string
	}{
		{In1 + "||" + Key1 + "->" + Out1, In1, Key1, Out1},
		{In2 + "||" + Key1 + "->" + Out2, In2, Key1, Out2},
		{In1 + "||" + Key2 + "->" + Out3, In1, Key2, Out3},
		{In2 + "||" + Key2 + "->" + Out4, In2, Key2, Out4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trithemius.EncodeMergeBlock(tt.intText, tt.key)

			if got != tt.outText {
				t.Errorf("EncodeMergeBlock(text=%q , key=%q), want %q but return %q", tt.intText, tt.key, tt.outText, got)
				return
			}

		})
	}
}

func TestTrithemius_DecodeMergeBlock(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		In1 = "БЛОК"
		In2 = "БРОК"
	)

	var (
		Key1 = "ХОРОШО_ВЫТЬ_ВАМИ"
		Key2 = "ХОРОШО_БЫТЬ_ВАМИ"
		Key3 = "ХОРОШО_ВЫТЬ_БАМИ"
	)

	var (
		Out1 = "ЬЗЦЩ"
		Out2 = "ЬМЬЩ"
		Out3 = "МЗЬТ"
		Out4 = "ММЬЧ"
	)

	tests := []struct {
		name    string
		intText string
		key     string
		outText string
	}{
		{In1 + "||" + Key1 + "->" + Out1, In1, Key1, Out1},
		{In2 + "||" + Key1 + "->" + Out2, In2, Key1, Out2},
		{"ЩЫУЯ" + "||" + Key1 + "->" + Out3, "ЩЫУЯ", Key1, Out3},
		{"Ф_ОИ" + "||" + Key1 + "->" + Out4, "Ф_ОИ", Key1, Out4},

		{"ЙРЫС" + "||" + Key2 + "->" + Out1, "ЙРЫС", Key2, Out1},
		{"ОР_М" + "||" + Key2 + "->" + Out2, "ОР_М", Key2, Out2},
		{In2 + "||" + Key2 + "->" + Out3, In1, Key2, Out3},
		{In2 + "||" + Key2 + "->" + Out4, In2, Key2, Out4},

		{"ДЛЭН" + "||" + Key3 + "->" + Out1, "ДЛЭН", Key3, Out1},
		{"_РБИ" + "||" + Key3 + "->" + Out2, "_РБИ", Key3, Out2},
		{"РЫИЧ" + "||" + Key3 + "->" + Out3, "РЫИЧ", Key3, Out3},
		{"Р_ГЧ" + "||" + Key3 + "->" + Out4, "Р_ГЧ", Key3, Out4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trithemius.DecodeMergeBlock(tt.outText, tt.key)

			if got != tt.intText {
				t.Errorf("DecodeMergeBlock(text=%q , key=%q), want %q but return %q", tt.outText, tt.key, tt.intText, got)
				return
			}

		})
	}
}

func TestTrithemius_EncodeSTrithemiusM_BasicTests(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		IN1 = "БЛОК"
	)

	var (
		KEY1 = "ХОРОШО_ВЫТЬ_ВАМИ"
	)

	var (
		OUT1 = "ЧФЮЖ"
	)

	tests := []struct {
		name    string
		inText  string
		key     string
		outText string
	}{
		{IN1 + "||" + KEY1 + "->" + OUT1, IN1, KEY1, OUT1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trithemius.EncodeSTrithemiusM(tt.inText, tt.key)

			if got != tt.outText {
				t.Errorf("STrithemiusM(text=%q , key=%q), want %q but return %q", tt.inText, tt.key, tt.outText, got)
				return
			}

		})
	}
}

func TestTrithemius_DecodeSTrithemiusM_BasicTests(t *testing.T) {
	trithemius := l1.NewTrithemius(TelegraphAlphabet)

	var (
		IN1 = "БЛОК"
	)

	var (
		KEY1 = "ХОРОШО_ВЫТЬ_ВАМИ"
	)

	var (
		OUT1 = "ЧФЮЖ"
	)

	tests := []struct {
		name    string
		inText  string
		key     string
		outText string
	}{
		{OUT1 + "||" + KEY1 + "->" + IN1, IN1, KEY1, OUT1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trithemius.DecodeSTrithemiusM(tt.outText, tt.key)

			if got != tt.inText {
				t.Errorf("STrithemiusM(text=%q , key=%q), want %q but return %q", tt.outText, tt.key, tt.inText, got)
				return
			}
		})
	}
}
