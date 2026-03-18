package compositeCipher_test

import (
	"fmt"
	"infbez_labs/internal/alphabet"
	"infbez_labs/internal/codeRandomGenerator"
	"infbez_labs/internal/compositeCipher"
	"infbez_labs/internal/core"
	"reflect"
	"testing"
)

func TestCompositeCipher_XOR(t *testing.T) {
	var (
		inA  = "АГАТ"
		inA1 = "КОЛЕНЬКА"
		inA2 = "ТОРТ_ХОЧЕТ_ГОРКУ"
		inB  = "ТАГА"
		inB1 = "МТВ_ТЛЕН"
		inB2 = "МТВ_ВСЕ_ЕЩЕ_ТЛЕН"

		sub1out     = "СДДС"
		block1out   = "СДДС"
		addText1out = "УДДУ"

		block2out   = "ЕЬОЕЭПМО"
		block3out   = "ЮЬСТВГИЧ_ИЕГЬЭМЩ"
		addText2out = "ЧБОЕАЗРО"
		addText3out = "_БУТВЗФЧЛМЕГБЭРБ"
	)

	tests := []struct {
		name   string
		action string
		input1 string
		input2 string
		output string
	}{
		{fmt.Sprintf("%s+%s->%s", inA, inB, sub1out), "subblock", inA, inB, sub1out},
		{fmt.Sprintf("%s+%s->%s", inA, inB, block1out), "block", inA, inB, block1out},
		{fmt.Sprintf("%s+%s->%s", inA, inB, addText1out), "addtext", inA, inB, addText1out},
		{fmt.Sprintf("%s+%s->%s", inA1, inB1, block2out), "block", inA1, inB1, block2out},
		{fmt.Sprintf("%s+%s->%s", inA2, inB2, block3out), "block", inA2, inB2, block3out},
		{fmt.Sprintf("%s+%s->%s", inA1, inB1, addText2out), "addtext", inA1, inB1, addText2out},
		{fmt.Sprintf("%s+%s->%s", inA2, inB2, addText3out), "addtext", inA2, inB2, addText3out},
		{fmt.Sprintf("%s+%s->%s", block3out, inB2, inA2), "block", block3out, inB2, inA2},
		{fmt.Sprintf("%s+%s->%s", block3out, inA2, inB2), "block", block3out, inA2, inB2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Alphabet := alphabet.NewAlphabet(alphabet.TelegraphAlphabet)

			var got string
			switch tt.action {
			case "subblock":
				got = Alphabet.SubBlocksXOR(tt.input1, tt.input2)
			case "block":
				got = Alphabet.BlockXOR(tt.input1, tt.input2)
			case "addtext":
				got = Alphabet.AddTxt(tt.input1, tt.input2)
			}

			if tt.output != got {
				t.Errorf("Failed %s(input1=%s, input2=%s), want %s but return %s", tt.action, tt.input1, tt.input2, tt.output, got)
				return
			}
		})
	}
}

func TestCompositeCipher_ProduceRoundKeys(t *testing.T) {
	var (
		key    = "ПОЛИМАТ_ТЕХНОБОГ"
		numin  = 6
		output = []string{
			"ЫТШЙЕО_ВЬЛЯГО_ЩП",
			"УЗШЯБЬЗЛАСЧЩББШШ",
			"ЮБЙ_Р_УОСХЫЬДШПШ",
			"ЩЖЬСМЕХСЛБМСЮРНБ",
			"УЛИСЛОБ_ЦАХПЮЕХН",
			"ЙЛППЯИОЕВЮЩЬБЙКЖ",
		}
	)

	tests := []struct {
		name   string
		key    string
		numin  int
		output []string
	}{
		{fmt.Sprintf("%s,%d", key, numin), key, numin, output},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Alphabet := alphabet.NewAlphabet(alphabet.TelegraphAlphabet)
			Sblock := core.NewSBlockSTM(Alphabet)
			Pblock := core.NewPBlock(Alphabet)
			LFSR := codeRandomGenerator.NewLFSR(Alphabet)
			SPNet := compositeCipher.NewSPNet(Alphabet, Sblock, Pblock, LFSR)

			got := SPNet.ProduceRoundKeys(tt.key, tt.numin)

			if !reflect.DeepEqual(tt.output, got) {
				t.Errorf("Failed ProduceRoundKeys(key=%s, numin=%d), want %v but return %v", tt.key, tt.numin, tt.output, got)
				return
			}
		})
	}
}

func TestCompositeCipher_RoundSP(t *testing.T) {
	var (
		//in  = "АБВГДЕЖЗИЙКЛМНОП"
		in1 = "КОРЫСТЬ_СЛОНА_ЭХ"
		in2 = "КОРЫСТЬ_СЛОН__ЭХ"

		key = "МТВ_ВСЕ_ЕЩЕ_ТЛЕН"
		//key1 = "МТВ_ВСЕ_ЕЩЕ_ТЛЕН"
		//key2 = "НТВ_ВСЕ_ЕЩЕ_ТЛЕН"

		out1t = "ДДПЕЯЛТРЫММПЕВЦЩ"
		out2t = "ДДПЕЦЛТРЫММПЕБХЩ"
		out3t = "ЩШЖАЮЛИГВПОМНЛКФ"
		out4t = "ХШ_НЮЛИВКПОМНЙЖФ"
		//lout1t  = "ЛСТЧЧРАМНТЛ_ОЯР_"
		//lout2t  = "ЛСТЧЧРАМНТЛ_Т_Т_"
		sub1out = "____З________АА_"
		sub2out = "Г_ЖТ___АЧ____БГ_"
	)

	tests := []struct {
		name   string
		action string
		input  string
		key    string
		output string
	}{
		{fmt.Sprintf("%s+%s->%s", in1, key, out1t), "frw", in1, key, out1t},
		{fmt.Sprintf("%s+%s->%s", in2, key, out2t), "frw", in2, key, out2t},
		//{fmt.Sprintf("%s+%s->%s", out1t, key, lout1t), "inv", out1t, key, lout1t},
		//{fmt.Sprintf("%s+%s->%s", out2t, key, lout2t), "inv", out2t, key, lout2t},
		{fmt.Sprintf("%s+%s->%s", out1t, out2t, sub1out), "sub", out1t, out2t, sub1out},
		//{fmt.Sprintf("%s+%s->%s", in, key1, out3t), "frw", in, key1, out3t},
		//{fmt.Sprintf("%s+%s->%s", in, key2, out4t), "frw", in, key2, out4t},
		{fmt.Sprintf("%s+%s->%s", out3t, out4t, sub2out), "sub", out3t, out4t, sub2out},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Alphabet := alphabet.NewAlphabet(alphabet.TelegraphAlphabet)
			Sblock := core.NewSBlockSTM(Alphabet)
			Pblock := core.NewPBlock(Alphabet)
			LFSR := codeRandomGenerator.NewLFSR(Alphabet)
			SPNet := compositeCipher.NewSPNet(Alphabet, Sblock, Pblock, LFSR)

			var got string
			switch tt.action {
			case "frw":
				got = SPNet.FrwRoundSP(tt.input, tt.key, 0)
			case "inv":
				got = SPNet.InvRoundSP(tt.input, tt.key, 0)
			case "sub":
				got = Alphabet.SubTxt(tt.input, tt.key)
			}

			if tt.output != got {
				t.Errorf("Failed %s(input=%s, key=%s), want %s but return %s", tt.action, tt.input, tt.key, tt.output, got)
				return
			}
		})
	}
}

func TestCompositeCipher_SeqRoundSP(t *testing.T) {
	var (
		in1 = "КОРЫСТЬ_СЛОНА_ЭХ"
		in2 = "КОРЫСТЬ_СЛОН__ЭХ"
		in3 = "КОРЫСТЬ_СЛОНА_ЭХ"
		in4 = "КОРЫСТЬ_СЛОН_АЭХ"
		key = "МТВ_ВСЕ_ЕЩЕ_ТЛЕН"

		out1t = "ДДПЕЯЛТРЫММПЕВЦЩ"
		out2t = "ДДПЕЦЛТРЫММПЕБХЩ"
		out3t = "ДДПЕЯЛТРЫММПЕВЦЩ"
		out4t = "ДФПЕЖЛТРЫММПЕБДЙ"
		/*
			qq = []string{
				"____З________АА_",
				"БЬВ____ТПВЛ____Ч",
				"ЗЗЦЧ_ЫЧЛЯШФЦЯЮПМ",
				"ХШБЖЭЖЖОХ_ЖЧЯЦЖС",
				"ЖЭГЛЖХФСЙБШЧЬЯЮЛ",
				"ДВАНТБХЖФШДЬОЦЧТ",
				"ЬЯНЧНЖШЮЩИЫЯМДЦГ",
				"БССОЯРАЭЬКМАОЮЫЙ",
			}

			dd = []string{
				"_П__Ч________АСП",
				"ИУАЗ___УПВУ_ЧЗ_Ч",
				"ЯЧЩППЯ_ПЛРУЖЦЧ_Ж",
				"СЙОЫХГКЧГОХЙОНОЭ",
				"РАВЕАТЭГЖЫОМПКУД",
				"ЗЦОБСЧЙЩАЬКВИЖШЦ",
				"МЯВРЧЛГЧМЗЩПЕЮ_У",
				"ЦЦОШЫБЗЯШППЯ_ЛШЙ",
			}*/
	)

	tests1 := []struct {
		name   string
		input  string
		key    string
		output string
	}{
		{fmt.Sprintf("%s+%s->%s", in1, key, out1t), in1, key, out1t},
		{fmt.Sprintf("%s+%s->%s", in2, key, out2t), in2, key, out2t},
		{fmt.Sprintf("%s+%s->%s", in3, key, out3t), in3, key, out3t},
		{fmt.Sprintf("%s+%s->%s", in4, key, out4t), in4, key, out4t},
	}

	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			Alphabet := alphabet.NewAlphabet(alphabet.TelegraphAlphabet)
			Sblock := core.NewSBlockSTM(Alphabet)
			Pblock := core.NewPBlock(Alphabet)
			LFSR := codeRandomGenerator.NewLFSR(Alphabet)
			SPNet := compositeCipher.NewSPNet(Alphabet, Sblock, Pblock, LFSR)

			got := SPNet.FrwRoundSP(tt.input, tt.key, 0)

			if tt.output != got {
				t.Errorf("Failed FrwRoundSP(input=%s, key=%s), want %s but return %s", tt.input, tt.key, tt.output, got)
				return
			}
		})
	}
	/*
		tests2 := []struct {
			name   string
			input1 string
			input2 string
			key    string
			output []string
		}{
			{fmt.Sprintf("%s / %s+%s->%s", out1t, out2t, key, qq), out1t, out2t, key, qq},
			{fmt.Sprintf("%s / %s+%s->%s", out3t, out4t, key, dd), out3t, out4t, key, dd},
		}

		for _, tt := range tests2 {
			t.Run(tt.name, func(t *testing.T) {
				Alphabet := alphabet.NewAlphabet(alphabet.TelegraphAlphabet)
				Sblock := core.NewSBlockSTM(Alphabet)
				Pblock := core.NewPBlock(Alphabet)
				LFSR := codeRandomGenerator.NewLFSR(Alphabet)
				SPNet := compositeCipher.NewSPNet(Alphabet, Sblock, Pblock, LFSR)

				got1 := []string{
					tt.input1,
				}
				for i := 1; i < 8; i++ {
					out := SPNet.FrwRoundSP(got1[i-1], tt.key, i)
					got1 = append(got1, out)
				}

				got2 := []string{
					tt.input2,
				}
				for i := 1; i < 8; i++ {
					out := SPNet.FrwRoundSP(got2[i-1], tt.key, i)
					got2 = append(got2, out)
				}

				var got []string
				for i := 0; i < 8; i++ {
					got = append(got, Alphabet.SubTxt(got1[i], got2[i]))
				}

				if !reflect.DeepEqual(got, tt.output) {
					t.Errorf("Failed SeqFrwRoundSP(input1=%s, input2=%s, key=%s), want %v but return %v", tt.input1, tt.input2, tt.key, tt.output, got)
					return
				}
			})
		}*/
}

func TestCompositeCipher_SPNet(t *testing.T) {
	var (
		in1 = "КОРЫСТЬ_СЛОНА_ЭХ"
		in2 = "ЛЕРА_КЛОНКА_КОНЯ"
		key = "МТВ_ВСЕ_ЕЩЕ_ТЛЕН"

		keysLT = []string{
			"ЯРЛОЫВЕЕЗФГППШМФ",
			"ДЛШЯАЫЬФЯУТПОЦЧЛ",
			"ЭНБХУЙХЬ_ИДДБ_ГУ",
			"БЖР_ДЯКЦХЦККЯХЬЛ",
			"ЬШЫЫФЖЕОЯОБМАЖФУ",
			"ЮОПЬЗЩКПИМЯЖЩАЗЛ",
			"ЛСПШШПЦУЮЗЭЖЙБТО",
			"В_АНЫЩЦСЖКЮОПРЗФ",
		}

		out1tf = "ЯПЛЦБФСЖХРЮИШФПФ"
		out2tf = "ПЬЖЗОВЗФРИИНКАЙЭ"
	)

	t.Run("KeysLT", func(t *testing.T) {
		Alphabet := alphabet.NewAlphabet(alphabet.TelegraphAlphabet)
		Sblock := core.NewSBlockSTM(Alphabet)
		Pblock := core.NewPBlock(Alphabet)
		LFSR := codeRandomGenerator.NewLFSR(Alphabet)
		SPNet := compositeCipher.NewSPNet(Alphabet, Sblock, Pblock, LFSR)

		got := SPNet.ProduceRoundKeys(key, 8)

		if !reflect.DeepEqual(got, keysLT) {
			t.Errorf("Failed ProduceRoundKeys(key=%s), want %v but return %v", key, keysLT, got)
			return
		}
	})

	tests := []struct {
		name   string
		action string
		input  string
		output string
	}{
		{fmt.Sprintf("%s->%s", in1, out1tf), "frw", in1, out1tf},
		{fmt.Sprintf("%s->%s", in2, out2tf), "frw", in2, out2tf},
		{fmt.Sprintf("%s->%s", out1tf, in1), "inv", out1tf, in1},
		{fmt.Sprintf("%s->%s", out2tf, in2), "inv", out2tf, in2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Alphabet := alphabet.NewAlphabet(alphabet.TelegraphAlphabet)
			Sblock := core.NewSBlockSTM(Alphabet)
			Pblock := core.NewPBlock(Alphabet)
			LFSR := codeRandomGenerator.NewLFSR(Alphabet)
			SPNet := compositeCipher.NewSPNet(Alphabet, Sblock, Pblock, LFSR)

			var got string
			switch tt.action {
			case "frw":
				got = SPNet.FrwSPNet(tt.input, key, 8)
			case "inv":
				got = SPNet.InvSPNet(tt.input, key, 8)
			}

			if got != tt.output {
				t.Errorf("Failed %s SPNet(input=%s), want %s but return %s", tt.action, tt.input, tt.output, got)
				return
			}
		})
	}
}
