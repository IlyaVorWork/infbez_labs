package main

import (
	"fmt"
	l1 "infbez_labs/Lab1"
	l2 "infbez_labs/Lab2"
)

func main() {
	var TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")
	var alphabet = *l1.NewAlphabet(TelegraphAlphabet)
	//var (
	//	TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")

	//	state0            = [5][5]string{
	//		{"____", "____", "____", "____", "____"},
	//		{"____", "____", "____", "____", "____"},
	//		{"____", "____", "____", "____", "____"},
	//		{"____", "____", "____", "____", "____"},
	//		{"____", "____", "____", "____", "____"},
	//	}
	//
	//	stateX = [5][5]string{
	//		{"БЫ_Щ", "ЙЖ_Б", "ЮФ_Е", "БЫ_Щ", "ЮД_Е"},
	//		{"Ы_ЩБ", "Ж_БЙ", "Ф_ЕЮ", "Ы_ЩБ", "Л_ЗЗ"},
	//		{"Ы_ЩБ", "Ж_БЙ", "Ф_ЕЮ", "У_ЧЧ", "Д_ЕЮ"},
	//		{"Ы_ЩБ", "Ж_БЙ", "Ь_ЗЗ", "Ы_ЩБ", "Д_ЕЮ"},
	//		{"Ы_ЩБ", "____", "Ф_ЕЮ", "Ы_ЩБ", "Д_ЕЮ"},
	//	}
	//)

	//sponge := l2.NewSponge(InnerState, alphabet)

	//var (
	//	IN  = "ХОРОШО_БЫТЬ_ВАМИ"
	//	IN1 = "КЬЕРГЕГОР_ПРОПАЛ"
	//	IN2 = "ХОРОШО_ПРОБРОСИЛ"
	//)
	//fmt.Println(sponge.Confuse(IN, IN1))
	//fmt.Println(sponge.Confuse(IN, IN2))
	//fmt.Println(sponge.Confuse(IN, IN))

	//var (
	//	IN1 = []string{"ХОРОШО_БЫТЬ_ВАМИ"}
	//	IN2 = []string{"ХОРОШО_БЫТЬ_ВАМИ", "________________", "________________", "________________"}
	//	IN3 = []string{"ХОРОШО_БЫТЬ_ВАМИ", "________А_______"}
	//	IN4 = []string{"ХОРОШО_БЫТЬ_ВАМИ", "КЬЕРГЕГОР_ПРОПАЛ"}
	//)
	//
	//fmt.Println(l2.NewSponge(InnerState, alphabet).CBlock(IN1, 16))
	//fmt.Println(l2.NewSponge(InnerState, alphabet).CBlock(IN2, 16))
	//fmt.Println(l2.NewSponge(InnerState, alphabet).CBlock(IN3, 16))
	//
	//fmt.Println(l2.NewSponge(InnerState, alphabet).CBlock(IN1, 8))
	//fmt.Println(l2.NewSponge(InnerState, alphabet).CBlock(IN2, 8))
	//fmt.Println(l2.NewSponge(InnerState, alphabet).CBlock(IN3, 8))
	//
	//fmt.Println(l2.NewSponge(InnerState, alphabet).CBlock(IN4, 16))
	//fmt.Println(l2.NewSponge(InnerState, alphabet).CBlock(IN4, 8))
	//fmt.Println(l2.NewSponge(InnerState, alphabet).CBlock(IN4, 4))

	//sponge := l2.NewSponge(state0, alphabet)
	//Display(state0)
	//state11 := sponge.MixCols()
	//Display(state11)
	//state12 := sponge.ShatterBlocks()
	//Display(state12)
	//state13 := sponge.ShiftRows()
	//Display(state13)
	//
	//state21 := sponge.MixCols()
	//Display(state21)
	//state22 := sponge.ShatterBlocks()
	//Display(state22)
	//state23 := sponge.ShiftRows()
	//Display(state23)
	//
	//state31 := sponge.MixCols()
	//Display(state31)
	//state32 := sponge.ShatterBlocks()
	//Display(state32)
	//state33 := sponge.ShiftRows()
	//Display(state33)

	//sponge := l2.NewSponge(state0, alphabet)
	//sponge.SpongeAbsorb("_А__")
	//sponge.SpongeAbsorb("ВИЛЯ")
	//sponge.SpongeAbsorb("ОЗЛ_")
	//sponge.SpongeAbsorb("ОЗЛ_")
	//Display(sponge.InnerState)
	//
	//spongeX := l2.NewSponge(stateX, alphabet)
	//out1 := spongeX.SpongeSqueeze()
	//out2 := spongeX.SpongeSqueeze()
	//out3 := spongeX.SpongeSqueeze()
	//fmt.Println(out1, out2, out3)
	//Display(spongeX.InnerState)

	fmt.Println(l2.SpongeHash("КАТЕГОРИЧЕСКИЙ_ИМПЕРАТИВ", alphabet))
	fmt.Println(l2.SpongeHash("________________________________________________________________", alphabet))
}

func Display(inp [5][5]string) {
	for i := 0; i < 5; i++ {
		fmt.Println(inp[i])
	}
	fmt.Println("")
}
