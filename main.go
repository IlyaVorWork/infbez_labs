package main

import (
	"fmt"
	l1 "infbez_labs/Lab1"
)

func main() {
	lab1 := l1.NewAlphabet()

	var key string
	var text string
	_, err := fmt.Scanln(&key)
	if err != nil {
		return
	}

	_, err = fmt.Scanln(&text)
	if err != nil {
		return
	}

	table := lab1.BuildThrithemusAlphabet(key)
	fmt.Println(string(table))
	encodedText := lab1.EncodePolyThrithemus(text, key)
	decodedText := lab1.DecodePolyThrithemus(encodedText, key)
	fmt.Println(encodedText)
	fmt.Println(decodedText)

	encodedSText := lab1.EncodeSThrithemus(text, key)
	decodedSText := lab1.DecodeSThrithemus(encodedSText, key)
	fmt.Println(encodedSText)
	fmt.Println(decodedSText)

	encodeMergeBlock := lab1.EncodeMergeBlock(text, key)
	decodedMergeBlock := lab1.DecodeMergeBlock(encodeMergeBlock, key)
	fmt.Println(encodeMergeBlock)
	fmt.Println(decodedMergeBlock)

	encodedTrithemusM := lab1.EncodeSThrithemusM(text, key)
	decodedTrithemusM := lab1.DecodeSThrithemusM(encodedTrithemusM, key)
	fmt.Println(encodedTrithemusM)
	fmt.Println(decodedTrithemusM)

	/*
		fmt.Println(string(lab1.ShiftThrithemusAlphabet(table, "Х", 16)))
		fmt.Println(string(lab1.ShiftThrithemusAlphabet(table, "Х", 18)))
		fmt.Println(string(lab1.ShiftThrithemusAlphabet(table, "Я", 16)))
		fmt.Println(string(lab1.ShiftThrithemusAlphabet(table, "А", 8)))
		fmt.Println(string(lab1.ShiftThrithemusAlphabet(table, "А", 23)))
		fmt.Println(string(lab1.ShiftThrithemusAlphabet(table, "А", 30)))
		fmt.Println(string(lab1.ShiftThrithemusAlphabet(table, "А", 1)))
	*/

}
