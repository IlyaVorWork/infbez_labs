package sponge_test

import (
	l1 "infbez_labs/Lab1"
	l2 "infbez_labs/Lab2"
	"strings"
	"testing"
)

func TestSponge_CBlock(t *testing.T) {
	var (
		TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")
		alphabet          = *l1.NewAlphabet(TelegraphAlphabet)
		sponge            = l2.NewSponge(l2.SpongeInnerState, alphabet)

		IN1 = []string{"ХОРОШО_БЫТЬ_ВАМИ"}
		IN2 = []string{"ХОРОШО_БЫТЬ_ВАМИ", "________________", "________________", "________________"}
		IN3 = []string{"ХОРОШО_БЫТЬ_ВАМИ", "________А_______"}

		IN1_16OUT = "ВУГТБТЩЩЬБЩШЕЗЛМ"
		IN2_16OUT = "ВУГТБТЩЩЬБЩШЕЗЛМ"
		IN3_16OUT = "ЫХВТАСЩШЬБШШНЗЛМ"

		IN1_8OUT = "ДЖЮМБЙЕЕ"
		IN2_8OUT = "ДЖЮМБЙЕЕ"
		IN3_8OUT = "ЬЗЭЛЙЙДЕ"

		IN4 = []string{"ХОРОШО_БЫТЬ_ВАМИ", "КЬЕРКЕГОР_ПРОПАЛ"}
		IN5 = []string{"ЧЕРНЫЙ_АББАТ_ПОЛ", "ХОРОШО_БЫТЬ_ВАМИ", "КЬЕРКЕГОР_ПРОПАЛ"}

		IN4_16OUT = "_ЭВТЮМШРЙШГЛДХЖН"
		IN4_8OUT  = "ЮЙЬГООКЩ"
		IN4_4OUT  = "ОЫРЙ"

		IN5_16OUT = "ШЯФЩЦНБЩЭФЭЬСХКК"
		IN5_8OUT  = "ПМЦУОКЗЖ"
		IN5_4OUT  = "АБОМ"

		IN6 = []string{
			"______А_________",
			"________________",
			"________________",
			"________________",
		}
		IN7 = []string{
			"________________",
			"________________",
			"________________",
			"________________",
		}
		IN8 = []string{
			"_____А__________",
			"___________А____",
			"_А______________",
			"____________А___",
		}

		IN6_16OUT = "ЭХЕБФ_ИЖАЕАУХЖОК"
		IN6_8OUT  = "СХОИЦМПЯ"
		IN6_4OUT  = "ЫИЯЙ"

		IN7_16OUT = "ЭХЕОМВЙЗЖУЕШХЦЮК"
		IN7_8OUT  = "ЙШПЦЭКГГ"
		IN7_4OUT  = "МНЛТ"

		IN8_16OUT = "ЭТЕЫМЮЛЙСУЮЗХТБЛ"
		IN8_8OUT  = "ЙРСДЗЖ_У"
		IN8_4OUT  = "БЙСР"

		IN6_16_SUB_IN8_16_OUT = "_В_ЖЗБЭЭОСВЛ_УМЯ"
		IN6_16_SUB_IN7_16_OUT = "___ТЗЭЯЯЩСЫЫ_ПР_"
		IN7_16_SUB_IN8_16_OUT = "_В_У_ДЮЮФ_ЗР_ГЬЯ"
	)

	tests := []struct {
		name   string
		input  []string
		size   int
		output string
	}{
		{"{ " + strings.Join(IN1, " ") + " } -> " + IN1_16OUT, IN1, 16, IN1_16OUT},
		{"{ " + strings.Join(IN1, " ") + " } -> " + IN1_8OUT, IN1, 8, IN1_8OUT},
		{"{ " + strings.Join(IN2, " ") + " } -> " + IN2_16OUT, IN2, 16, IN2_16OUT},
		{"{ " + strings.Join(IN2, " ") + " } -> " + IN2_8OUT, IN2, 8, IN2_8OUT},
		{"{ " + strings.Join(IN3, " ") + " } -> " + IN3_16OUT, IN3, 16, IN3_16OUT},
		{"{ " + strings.Join(IN3, " ") + " } -> " + IN3_8OUT, IN3, 8, IN3_8OUT},

		{"{ " + strings.Join(IN4, " ") + " } -> " + IN4_16OUT, IN4, 16, IN4_16OUT},
		{"{ " + strings.Join(IN4, " ") + " } -> " + IN4_8OUT, IN4, 8, IN4_8OUT},
		{"{ " + strings.Join(IN4, " ") + " } -> " + IN4_4OUT, IN4, 4, IN4_4OUT},

		{"{ " + strings.Join(IN5, " ") + " } -> " + IN5_16OUT, IN5, 16, IN5_16OUT},
		{"{ " + strings.Join(IN5, " ") + " } -> " + IN5_8OUT, IN5, 8, IN5_8OUT},
		{"{ " + strings.Join(IN5, " ") + " } -> " + IN5_4OUT, IN5, 4, IN5_4OUT},

		{"{ " + strings.Join(IN6, " ") + " } -> " + IN6_16OUT, IN6, 16, IN6_16OUT},
		{"{ " + strings.Join(IN6, " ") + " } -> " + IN6_8OUT, IN6, 8, IN6_8OUT},
		{"{ " + strings.Join(IN6, " ") + " } -> " + IN6_4OUT, IN6, 4, IN6_4OUT},

		{"{ " + strings.Join(IN7, " ") + " } -> " + IN7_16OUT, IN7, 16, IN7_16OUT},
		{"{ " + strings.Join(IN7, " ") + " } -> " + IN7_8OUT, IN7, 8, IN7_8OUT},
		{"{ " + strings.Join(IN7, " ") + " } -> " + IN7_4OUT, IN7, 4, IN7_4OUT},

		{"{ " + strings.Join(IN8, " ") + " } -> " + IN8_16OUT, IN8, 16, IN8_16OUT},
		{"{ " + strings.Join(IN8, " ") + " } -> " + IN8_8OUT, IN8, 8, IN8_8OUT},
		{"{ " + strings.Join(IN8, " ") + " } -> " + IN8_4OUT, IN8, 4, IN8_4OUT},
	}

	subTests := []struct {
		name   string
		input1 []string
		input2 []string
		size   int
		output string
	}{
		{"IN6_16 - IN8_16", IN6, IN8, 16, IN6_16_SUB_IN8_16_OUT},
		{"IN6_16 - IN7_16", IN6, IN7, 16, IN6_16_SUB_IN7_16_OUT},
		{"IN7_16 - IN8_16", IN7, IN8, 16, IN7_16_SUB_IN8_16_OUT},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sponge.CBlock(tt.input, tt.size)

			if tt.output != got {
				t.Errorf("Failed CBlock(input=%q, size=%q), want %v but return %v", tt.input, tt.size, tt.output, got)
				return
			}

		})
	}

	for _, tt := range subTests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := sponge.CBlock(tt.input1, tt.size)
			got2 := sponge.CBlock(tt.input2, tt.size)
			subGot := sponge.Alphabet.SubTxt(got1, got2)

			if tt.output != subGot {
				t.Errorf("Failed Subtract CBlock, want %v but return %v", tt.output, subGot)
				return
			}

		})
	}
}

func TestSponge_MixCols_ShatterBlocks_ShiftRows(t *testing.T) {
	var (
		TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")
		alphabet          = *l1.NewAlphabet(TelegraphAlphabet)

		state0 = [5][5]string{
			{"____", "____", "____", "____", "____"},
			{"__А_", "____", "____", "____", "____"},
			{"____", "____", "____", "____", "____"},
			{"____", "____", "____", "____", "____"},
			{"____", "____", "____", "____", "____"},
		}

		sponge = l2.NewSponge(state0, alphabet)

		state11 = [5][5]string{
			{"__М_", "__А_", "__В_", "__М_", "__Т_"},
			{"__М_", "____", "__Г_", "__Л_", "__У_"},
			{"__М_", "__А_", "__В_", "__М_", "__Т_"},
			{"__М_", "__А_", "__В_", "__М_", "__Т_"},
			{"__М_", "__А_", "__В_", "__М_", "__Т_"},
		}

		state12 = [5][5]string{
			{"___М", "__А_", "__В_", "__М_", "__Т_"},
			{"__М_", "____", "__Г_", "__Л_", "__У_"},
			{"__М_", "__А_", "___В", "__М_", "__Т_"},
			{"__М_", "__А_", "__В_", "___М", "__Т_"},
			{"__М_", "__А_", "__В_", "__М_", "___Т"},
		}

		state13 = [5][5]string{
			{"___М", "____", "___В", "___М", "___Т"},
			{"__М_", "__А_", "__В_", "__М_", "__Т_"},
			{"__М_", "__А_", "__В_", "__М_", "__У_"},
			{"__М_", "__А_", "__В_", "__Л_", "__Т_"},
			{"__М_", "__А_", "__Г_", "__М_", "__Т_"},
		}

		state21 = [5][5]string{
			{"__ЯА", "__У_", "___Ц", "__ЭИ", "__ЙВ"},
			{"__ЦЙ", "__ЗМ", "__ОЖ", "__ЫЛ", "__ЯМ"},
			{"__ХЙ", "__ЗМ", "__ОЖ", "__ЫЛ", "___М"},
			{"__ХЙ", "__ЗМ", "__ОЖ", "__ЩЛ", "___М"},
			{"__ХЙ", "__ЗМ", "__ПЖ", "__ЩЛ", "___М"},
		}

		state22 = [5][5]string{
			{"А__Я", "__У_", "___Ц", "__ЭИ", "__ЙВ"},
			{"__ЦЙ", "М__З", "__ОЖ", "__ЫЛ", "__ЯМ"},
			{"__ХЙ", "__ЗМ", "Ж__О", "__ЫЛ", "___М"},
			{"__ХЙ", "__ЗМ", "__ОЖ", "Л__Щ", "___М"},
			{"__ХЙ", "__ЗМ", "__ПЖ", "__ЩЛ", "М___"},
		}

		state23 = [5][5]string{
			{"А__Я", "М__З", "Ж__О", "Л__Щ", "М___"},
			{"__ЦЙ", "__ЗМ", "__ОЖ", "__ЩЛ", "__ЙВ"},
			{"__ХЙ", "__ЗМ", "__ПЖ", "__ЭИ", "__ЯМ"},
			{"__ХЙ", "__ЗМ", "___Ц", "__ЫЛ", "___М"},
			{"__ХЙ", "__У_", "__ОЖ", "__ЫЛ", "___М"},
		}

		state31 = [5][5]string{
			{"Ф_ОИ", "М_ШП", "К_ЦЙ", "Л_ЦЦ", "Ш_НЛ"},
			{"П_ЧД", "А_ЙЙ", "П_ФЗ", "Ы_ТК", "Э_ЬЫ"},
			{"П_ДЧ", "А_КЙ", "П_ФЗ", "Ы_ХЗ", "Э_НЗ"},
			{"П_СК", "А_КЙ", "П_ДЧ", "Ы_ГЫ", "Э_АФ"},
			{"П_ОН", "А_ЦЭ", "П_ЗФ", "Ы_АЮ", "Э_ГС"},
		}

		state32 = [5][5]string{
			{"ИФ_О", "М_ШП", "К_ЦЙ", "Л_ЦЦ", "Ш_НЛ"},
			{"П_ЧД", "ЙА_Й", "П_ФЗ", "Ы_ТК", "Э_ЬЫ"},
			{"П_ДЧ", "А_КЙ", "ЗП_Ф", "Ы_ХЗ", "Э_НЗ"},
			{"П_СК", "А_КЙ", "П_ДЧ", "ЫЫ_Г", "Э_АФ"},
			{"П_ОН", "А_ЦЭ", "П_ЗФ", "Ы_АЮ", "СЭ_Г"},
		}

		state33 = [5][5]string{
			{"ИФ_О", "ЙА_Й", "ЗП_Ф", "ЫЫ_Г", "СЭ_Г"},
			{"П_ЧД", "А_КЙ", "П_ДЧ", "Ы_АЮ", "Ш_НЛ"},
			{"П_ДЧ", "А_КЙ", "П_ЗФ", "Л_ЦЦ", "Э_ЬЫ"},
			{"П_СК", "А_ЦЭ", "К_ЦЙ", "Ы_ТК", "Э_НЗ"},
			{"П_ОН", "М_ШП", "П_ФЗ", "Ы_ХЗ", "Э_АФ"},
		}

		state41 = [5][5]string{
			{"ЯРШЦ", "ЙАЮ_", "ЬГ_У", "ЩЫР_", "ЬМУП"},
			{"ЖБСЕ", "ЩФРЙ", "У_СМ", "БГ_А", "ЫЖТЦ"},
			{"УБЦТ", "ЩФГЦ", "У_БЭ", "ТГЕЙ", "НЖЫЭ"},
			{"ЗБЮЧ", "ЩФВЦ", "О_СС", "ЖГСИ", "ЩЖАК"},
			{"ОБСЭ", "ЕФЗЖ", "З_К_", "НГЬЧ", "ТЖЙИ"},
		}

		state42 = [5][5]string{
			{"ЦЯРШ", "ЙАЮ_", "ЬГ_У", "ЩЫР_", "ЬМУП"},
			{"ЖБСЕ", "ЙЩФР", "У_СМ", "БГ_А", "ЫЖТЦ"},
			{"УБЦТ", "ЩФГЦ", "ЭУ_Б", "ТГЕЙ", "НЖЫЭ"},
			{"ЗБЮЧ", "ЩФВЦ", "О_СС", "ИЖГС", "ЩЖАК"},
			{"ОБСЭ", "ЕФЗЖ", "З_К_", "НГЬЧ", "ИТЖЙ"},
		}

		state43 = [5][5]string{
			{"ЦЯРШ", "ЙЩФР", "ЭУ_Б", "ИЖГС", "ИТЖЙ"},
			{"ЖБСЕ", "ЩФГЦ", "О_СС", "НГЬЧ", "ЬМУП"},
			{"УБЦТ", "ЩФВЦ", "З_К_", "ЩЫР_", "ЫЖТЦ"},
			{"ЗБЮЧ", "ЕФЗЖ", "ЬГ_У", "БГ_А", "НЖЫЭ"},
			{"ОБСЭ", "ЙАЮ_", "У_СМ", "ТГЕЙ", "ЩЖАК"},
		}

		tests = []struct {
			name      string
			input     [5][5]string
			operation string
			output    [5][5]string
		}{
			{"MIX 1", state0, "MIX", state11},
			{"SHATTER 1", state11, "SHATTER", state12},
			{"SHIFT 1", state12, "SHIFT", state13},
			{"MIX 2", state13, "MIX", state21},
			{"SHATTER 2", state21, "SHATTER", state22},
			{"SHIFT 2", state22, "SHIFT", state23},
			{"MIX 3", state23, "MIX", state31},
			{"SHATTER 3", state31, "SHATTER", state32},
			{"SHIFT 3", state32, "SHIFT", state33},
			{"MIX 4", state33, "MIX", state41},
			{"SHATTER 4", state41, "SHATTER", state42},
			{"SHIFT 4", state42, "SHIFT", state43},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got [5][5]string
			switch tt.operation {
			case "MIX":
				got = sponge.MixCols()
			case "SHATTER":
				got = sponge.ShatterBlocks()
			case "SHIFT":
				got = sponge.ShiftRows()
			}

			if tt.output != got {
				t.Errorf("Failed %v, want %v but return %v", tt.operation, tt.output, got)
				return
			}

		})
	}
}

func TestSponge_SpongeAbsorb(t *testing.T) {
	var (
		TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")
		alphabet          = *l1.NewAlphabet(TelegraphAlphabet)

		state0 = [5][5]string{
			{"____", "____", "____", "____", "____"},
			{"____", "____", "____", "____", "____"},
			{"____", "____", "____", "____", "____"},
			{"____", "____", "____", "____", "____"},
			{"____", "____", "____", "____", "____"},
		}

		sponge = l2.NewSponge(state0, alphabet)

		IN1 = "_А__"
		IN2 = "ВИЛЯ"
		IN3 = "ОЗЛ_"

		state1 = [5][5]string{
			{"ЖИХЬ", "ВМНЛ", "ИЖЙГ", "ЖИХЬ", "ШЦЙГ"},
			{"ИХЬЖ", "МНЛВ", "ЖЙГИ", "ИХЬЖ", "ГЧПЬ"},
			{"ИХЬЖ", "МНЛВ", "ЖЙГИ", "ЬЗПГ", "ЦЙГШ"},
			{"ИХЬЖ", "МНЛВ", "УЧПЛ", "ИХЬЖ", "ЦЙГШ"},
			{"ИХЬЖ", "____", "ЖЙГИ", "ИХЬЖ", "ЦЙГШ"},
		}

		state2 = [5][5]string{
			{"ГНЙЬ", "ЗННУ", "ЕККС", "ЭКХЛ", "МСКЙ"},
			{"РУ_И", "ННУЗ", "ЧШЮИ", "КХЛЭ", "СЭЕЫ"},
			{"РУ_И", "ННУЗ", "ЧШЮИ", "МРХЩ", "СКЙМ"},
			{"РУ_И", "А_ЗД", "УРНА", "ЧГЧ_", "СКЙМ"},
			{"РУ_И", "ЖДЮЗ", "ККСЕ", "КХЛЭ", "СКЙМ"},
		}

		state3 = [5][5]string{
			{"ПНЯС", "ДОЫН", "МЭРУ", "ЖЦСЯ", "ЖДЬЭ"},
			{"КР_М", "ОЫНД", "ЕЦПЗ", "ШОДБ", "ДБЦН"},
			{"МЮЕЧ", "БМББ", "ЦЛЮЙ", "УПЛЧ", "КТОЧ"},
			{"ОКЛВ", "ЗСЧД", "Н_АВ", "ТЙОЯ", "ИЕИМ"},
			{"РЗСЮ", "ЛЮНЧ", "ЭРУМ", "ФДШЬ", "ЖШВБ"},
		}

		state4 = [5][5]string{
			{"ЬВИГ", "НИРУ", "УЫШД", "ЗССО", "ЫПЖЧ"},
			{"ЗХПР", "ЩХБ_", "НФЧФ", "ЖКХТ", "РШЧХ"},
			{"ЯВДЧ", "ЮНСЧ", "ЫЦЫЖ", "ЕДОТ", "ЙЧЙЮ"},
			{"ОЫАУ", "_ЭБП", "НОЫЙ", "КВЧЬ", "ФЧЫБ"},
			{"ЗЮУД", "ЭРЯУ", "ЫИДГ", "ЕУЯБ", "ЖМДР"},
		}

		tests = []struct {
			name   string
			input  string
			output [5][5]string
		}{
			{"ABSORB 1", IN1, state1},
			{"ABSORB 2", IN2, state2},
			{"ABSORB 3", IN3, state3},
			{"ABSORB 4", IN3, state4},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sponge.SpongeAbsorb(tt.input)

			if tt.output != got {
				t.Errorf("Failed Sponge Absorb, want %v but return %v", tt.output, got)
				return
			}

		})
	}
}

func TestSponge_SpongeSqueeze(t *testing.T) {
	var (
		TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")
		alphabet          = *l1.NewAlphabet(TelegraphAlphabet)

		stateX = [5][5]string{
			{"БЫ_Щ", "ЙЖ_Б", "ЮФ_Е", "БЫ_Щ", "ЮД_Е"},
			{"Ы_ЩБ", "Ж_БЙ", "Ф_ЕЮ", "Ы_ЩБ", "Л_ЗЗ"},
			{"Ы_ЩБ", "Ж_БЙ", "Ф_ЕЮ", "У_ЧЧ", "Д_ЕЮ"},
			{"Ы_ЩБ", "Ж_БЙ", "Ь_ЗЗ", "Ы_ЩБ", "Д_ЕЮ"},
			{"Ы_ЩБ", "____", "Ф_ЕЮ", "Ы_ЩБ", "Д_ЕЮ"},
		}

		sponge = l2.NewSponge(stateX, alphabet)

		OUT1 = "ШНОТ"
		OUT2 = "НБНЮ"
		OUT3 = "ИФСУ"

		tests = []struct {
			name   string
			output string
		}{
			{"SQUEEZE 1", OUT1},
			{"SQUEEZE 2", OUT2},
			{"SQUEEZE 3", OUT3},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sponge.SpongeSqueeze()

			if tt.output != got {
				t.Errorf("Failed Sponge Squeeze, want %v but return %v", tt.output, got)
				return
			}

		})
	}
}

func TestSponge_Hash(t *testing.T) {
	var (
		TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")
		alphabet          = *l1.NewAlphabet(TelegraphAlphabet)

		IN1 = "КАТЕГОРИЧЕСКИЙ_ИМПЕРАТИВ"
		IN2 = "________________________________________________________________"
		IN3 = "______________________А_________________________________________"
		IN4 = "________А_______________________________________________________"
		IN5 = "ПЕТЯ_ПИЛ_ПИВО_В_КАЛЬЯННОЙ_И_КУРИЛ_БАМБУК_ЧЕРЕЗ_АНАНАС_ТЧК_НАСТЯ_ПИЛА_ВОДУ_И_НЕ_ПОШЛА_В_КАЛЬЯННУЮ_ЗПТ_ЧТОБЫ_ВЫСПАТЬСЯ"
		//IN6 = "ЗОЛОТЫЕ_ВРЕМЕНА_ПРОШЛИ_ТЧК_НАСТАЛА_ПОРА_ГРУЗИТЬ_АПЕЛЬСИНЫ_БОЧКАМИ_И_НЕ_ОГЛЯДЫВАТЬСЯ_НАЗАД_ТЧК_КОГДАТО_СНОВА_МЫ_БУДЕМ_ТАМ_ГДЕ_НАС_ЖДУТ_ТЧК"

		OUT1 = "Й_НСПЮБЮЛЯОЯЛЩГЩВЧЫЫШЩГДЖФОНЕЙЮЫДПЬФШШАНЦДДЮ_ШЯДЕОНЬДЦА_ГЬЭДЩЙПИ"
		OUT2 = "_КЫЫШГЛВЯЖМНАЫ_ТУТЬИВЬЧПЖФЕПЕЬРХУЬСАЖЗЦТОБЯСЫУЫЬФЮШБШНЮДУЮЗЩКЙАХ"
		OUT3 = "ГЗМВЭЮГШКЖЯЛБИЧМРЩОЭЯМСМХУВЦШЮПЯЗДЩСЖЬТВЕЩНРЭУШЛВЗРЬМИИ_ДХИВЧТДВ"
		OUT4 = "_КЫЫШГЛВЯЖМНАЫ_ТУТЬИВЬЧПЖФЕПЕЬРХУЬСАЖЗЦТОБЯСЫУЫЬФЮШБШНЮДУЮЗЩКЙАХ"
		OUT5 = "БЕЭНТЛБСОПМЫЬЛЧБХЭЗЩИЗЛЭЦТЕЗЖЩЖРМТЯЫЮКЫРЫХОЕЮДКРФИИЕМДТЮ_ОНМЙШЦЭ"
		//OUT6 = "ХБХШЯБЕГККТСЦМПТРЖИЩМ_АКЗВЖАЧЕ_ДАПКФИЖЕЧДСЖБРЦЗПШЖЩСОЙЖОЯГЮГДОЗП"

		SUB_OUT1 = "ЬВНЧЬЕЗЙУ_НБЯСЗЕВШМЛГОЕВРАВШМЮАЦЛЦЧО_ЛГПИЗРАЮ_БПСХЗЕЛДФДОЗЯЦТЦЬТ"
		SUB_OUT2 = "________________________________________________________________"

		tests = []struct {
			name   string
			input  string
			output string
		}{
			{IN1 + "->" + OUT1, IN1, OUT1},
			{IN2 + "->" + OUT2, IN2, OUT2},
			{IN3 + "->" + OUT3, IN3, OUT3},
			{IN4 + "->" + OUT4, IN4, OUT4},
			{IN5 + "->" + OUT5, IN5, OUT5},
			//{IN6 + "->" + OUT6, IN6, OUT6},
		}

		subTests = []struct {
			name   string
			input1 string
			input2 string
			output string
		}{
			{IN2 + "-" + IN3 + "->" + OUT1, IN2, IN3, SUB_OUT1},
			{IN2 + "-" + IN4 + "->" + OUT1, IN2, IN4, SUB_OUT2},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := l2.SpongeHash(tt.input, alphabet)

			if tt.output != got {
				t.Errorf("Failed Sponge Hash, want %v but return %v", tt.output, got)
				return
			}

		})
	}

	for _, tt := range subTests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := l2.SpongeHash(tt.input1, alphabet)
			got2 := l2.SpongeHash(tt.input2, alphabet)
			subGot := alphabet.SubTxt(got1, got2)

			if tt.output != subGot {
				t.Errorf("Failed Sponge Sub Hash, want %v but return %v", tt.output, subGot)
				return
			}

		})
	}
}
