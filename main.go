package main

import (
	"fmt"
	l1 "infbez_labs/Lab1"
	l2 "infbez_labs/Lab2"
)

func main() {
	var (
		TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")

		InnerState = [4]string{"________________",
			"ПРОЖЕКТОР_ЧЕПУХИ",
			"КОЛЫХАТЬ_ПАРОДИЮ",
			"КАРМАННЫЙ_АТАМАН"}
	)

	alphabet := *l1.NewAlphabet(TelegraphAlphabet)
	sponge := l2.NewSponge(InnerState, alphabet)

	var (
		IN  = "ХОРОШО_БЫТЬ_ВАМИ"
		IN1 = "КЬЕРГЕГОР_ПРОПАЛ"
		IN2 = "ХОРОШО_ПРОБРОСИЛ"
	)
	fmt.Println(sponge.Confuse(IN, IN1))
	fmt.Println(sponge.Confuse(IN, IN2))
	fmt.Println(sponge.Confuse(IN, IN))

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
}
