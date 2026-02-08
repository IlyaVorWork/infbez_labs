package main

import (
	"fmt"
	l1 "infbez_labs/Lab1"
)

func main() {
	var TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")

	lab1 := l1.NewTrithemius(TelegraphAlphabet)

	var (
		key  string
		text string
	)

	_, err := fmt.Scanln(&text)
	if err != nil {
		return
	}

	_, err = fmt.Scanln(&key)
	if err != nil {
		return
	}

	table := lab1.BuildTrithemiusAlphabet(key)
	fmt.Println(string(table))
	encodedText := lab1.EncodePolyTrithemius(text, key)
	decodedText := lab1.DecodePolyTrithemius(encodedText, key)
	fmt.Println(encodedText)
	fmt.Println(decodedText)

	encodedSText := lab1.EncodeSTrithemius(text, key)
	decodedSText := lab1.DecodeSTrithemius(encodedSText, key)
	fmt.Println(encodedSText)
	fmt.Println(decodedSText)

	encodeMergeBlock := lab1.EncodeMergeBlock(text, key)
	decodedMergeBlock := lab1.DecodeMergeBlock(encodeMergeBlock, key)
	fmt.Println(encodeMergeBlock)
	fmt.Println(decodedMergeBlock)

	encodedTrithemiusM := lab1.EncodeSTrithemiusM(text, key)
	decodedTrithemiusM := lab1.DecodeSTrithemiusM(encodedTrithemiusM, key)
	fmt.Println(encodedTrithemiusM)
	fmt.Println(decodedTrithemiusM)

}
