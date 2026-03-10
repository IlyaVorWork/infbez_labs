package main

import (
	"fmt"
	"infbez_labs/internal/alphabet"
	"infbez_labs/internal/codeRandomGenerator"
	"infbez_labs/internal/compositeCipher"
	"infbez_labs/internal/core"
)

func main() {
	var (
		telegraphAlphabet = alphabet.NewAlphabet(alphabet.TelegraphAlphabet)
		pBlock            = core.NewPBlock(telegraphAlphabet)
		sBlockSTM         = core.NewSBlockSTM(telegraphAlphabet)
		lfsr              = codeRandomGenerator.NewLFSR(telegraphAlphabet)
		sPNet             = compositeCipher.NewSPNet(telegraphAlphabet, sBlockSTM, pBlock, lfsr)
	)
	fmt.Println(sPNet.FrwSPNet("КОРЫСТЬ_СЛОНА_ЭХ", "МТВ_ВСЕ_ЕЩЕ_ТЛЕН", 8))
	fmt.Println(sPNet.InvSPNet("ЯПЛЦБФСЖХРЮИШФПФ", "МТВ_ВСЕ_ЕЩЕ_ТЛЕН", 8))
	fmt.Println(sPNet.FrwSPNet("ЛЕРА_КЛОНКА_КОНЯ", "МТВ_ВСЕ_ЕЩЕ_ТЛЕН", 8))
	//fmt.Println(pBlock.FrwRound("ЗОЛОТАЯ_СЕРЕДИНА", 1))
	//fmt.Println(pBlock.InvRound("ПЯУЦШВГЖ_СПВЕЖЧШ", 1))
	//fmt.Println(pBlock.FrwRound("ЗОЛОТАЯ_СЕРЕДИНА", 2))

	//var a = pBlock.Text16CharTo80Bit("ЗОЛОТАЯ_СЕРЕДИНА")
	//fmt.Println(a)
	//fmt.Println(pBlock.Text80BitTo16Char(a))

	//in1 := []int{0, 0, 1, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1}
	//fmt.Println(pBlock.BinaryShift(in1, -41))

	//in1 := "АБВГДЕЖЗИЙКЛМНОП"
	//
	//out1 := pBlock.FrwMagicSquare(in1, core.MS1)
	//out2 := pBlock.FrwMagicSquare(in1, core.MS2)
	//out3 := pBlock.FrwMagicSquare(in1, core.MS3)
	//
	//fmt.Println(out1)
	//fmt.Println(out2)
	//fmt.Println(out3)
	//
	//fmt.Println(pBlock.InvMagicSquare(out1, core.MS1))
	//fmt.Println(pBlock.InvMagicSquare(out2, core.MS2))
	//fmt.Println(pBlock.InvMagicSquare(out3, core.MS3))

	//var (
	//	inA = "АГАТ"
	//	inB = "ТАГА"
	//
	//	inA1 = "КОЛЕНЬКА"
	//	inB1 = "МТВ_ТЛЕН"
	//
	//	inA2 = "ТОРТ_ХОЧЕТ_ГОРКУ"
	//	inB2 = "МТВ_ВСЕ_ЕЩЕ_ТЛЕН"
	//)
	//
	//fmt.Println(telegraphAlphabet.BlockXOR(inA, inB))
	//fmt.Println(telegraphAlphabet.BlockXOR(inA1, inB1))
	//fmt.Println(telegraphAlphabet.BlockXOR(inA2, inB2))

	// -------------

	//var sBlock = core.NewSBlock(*telegraphAlphabet)
	//var lfsr = generator.NewLFSR(*telegraphAlphabet)
	//var sPNet = compositeCipher.NewSPNet(*sBlock, *lfsr)
	//
	//fmt.Println(sPNet.ProduceRoundKeys("ПОЛИМАТ_ТЕХНОБОГ", 6))

	// -------------

}
