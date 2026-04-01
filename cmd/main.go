package main

import (
	"bufio"
	"fmt"
	"infbez_labs/internal/alphabet"
	"infbez_labs/internal/authEncryptionProtocol"
	"os"
	"strings"
)

func main() {
	var (
		telegraphAlphabet = alphabet.NewAlphabet(alphabet.TelegraphAlphabet)
		//packet            = authEncryptionProtocol.CreatePacket(telegraphAlphabet)
		//	cBlock            = core.NewCBlock(telegraphAlphabet)
		//	sBlockPoly        = core.NewSBlockPloy(telegraphAlphabet)
		//	hasher            = hash.NewHasher(telegraphAlphabet, cBlock)
		//	key               = auth.NewKeyDerivation(hasher, telegraphAlphabet)
		//
		//	pass1, pass2 = "ЧЕЧЕТКА", "АПРОЛ"
		//	salt1, salt2 = "СЕАНС", "АТЛЕТ"
		//	size         = []int{32, 16}
		//	context      = []string{"СЕАНСОВЫЙ_КЛЮЧ", "КЛЮЧ_РАСПРЕДЕЛЕНИЯ_КЛЮЧЕЙ"}
	)
	//
	//fmt.Println(sBlockPoly.FrwRun("ХОРОШО_БЫТЬ_ВАМИ", "КЬЕРКЕГОР_ПРОПАЛ"))
	//
	//fmt.Println(hasher.Hash("КАТЕГОРИЧЕСКИЙ_ИМПЕРАТИВ"))
	//
	//fmt.Println(key.Run(pass1, salt1, context, size, 2))
	//fmt.Println(key.Run(pass2, salt1, context, size, 2))
	//fmt.Println(key.Run(pass1, salt2, context, size, 2))
	//fmt.Println(key.Run(pass1, salt2, []string{"МАСТЕР_КЛЮЧ"}, []int{120}, 2))
	//in1 := "ГНОЛЛЫ_ПИЛИЛИ_ПЫЛЕСОС_ЛОСОСЕМ"

	//a, _ := telegraphAlphabet.MessageToBin([]rune("ГНОЛЛЫ_ПИЛИЛИ_ПЫЛЕСОС_ЛОСОСЕМ1110011011011"))
	//fmt.Println(a)
	//fmt.Println(telegraphAlphabet.BinToMessage(a))

	//fileAd, _ := os.ReadFile("sources/ad.txt")
	//fileInp, _ := os.ReadFile("sources/inp.txt")
	//
	InputsArray := make([]string, 0)
	AssocdataArray := make([]string, 0)

	err := readFileByLines("sources/inp.txt", func(line string) {
		InputsArray = append(InputsArray, line)
	})
	if err != nil {
		return
	}

	err = readFileByLines("sources/ad.txt", func(line string) {
		AssocdataArray = append(AssocdataArray, line)
	})
	if err != nil {
		return
	}
	//
	//fmt.Println(AssocdataArray)
	//fmt.Println(InputsArray)
	//
	//IN1 := []rune(InputsArray[0])
	//fmt.Println(len(IN1))
	//mtb1, _ := telegraphAlphabet.MessageToBin(IN1)
	//fmt.Println(len(mtb1))
	//inter := packet.PadMessage(IN1)
	//fmt.Println(len(inter))
	//mtb2, _ := telegraphAlphabet.MessageToBin(inter)
	//fmt.Println(len(mtb2))
	//out := packet.PadMessage(inter)
	//fmt.Println(len(out))
	//mtb3, _ := telegraphAlphabet.MessageToBin(out)
	//fmt.Println(len(mtb3))
	//
	//fmt.Println(telegraphAlphabet.AddTxt("________________", "АБВГД"))

	//gt := packet.PreparePacket(
	//	[4]string(strings.Fields(strings.ReplaceAll(AssocdataArray[1], `"`, ""))),
	//	"КОЛЕСО",
	//	InputsArray[1])
	//
	//fmt.Println(gt.String())
	//
	//packett := packet.Recieve(packet.Transmit(gt))
	//fmt.Println(packett.String())

	TST := InputsArray[0]
	fmt.Println(TST)

	IV1 := []rune("АЛИСА_УМЕЕТ_ПЕТЬ")
	//IV2 := "БОБ_НЕМНОГО_ПЬЯН"

	cfb := authEncryptionProtocol.NewCFB(telegraphAlphabet)
	cfb1 := cfb.Forward([]rune(TST), IV1, "СЕАНСОВЫЙ_КЛЮЧИК", 1)
	fmt.Println(cfb1)
	cfd11 := cfb.Invert([]rune(cfb1), IV1, "СЕАНСОВЫЙ_КЛЮЧИК", 1)
	fmt.Println(cfd11)

	fmt.Println(TST)
	//fmt.Println(AD)
	fmt.Println("----------Тестики----------")
	eax := authEncryptionProtocol.NewEAX(telegraphAlphabet)

	SEC := "ТОЖЕ_ЕЩЕ_НЕВАЖНО"
	AssocData := ADPrepare(AssocdataArray[1], "АБВГД")

	packet := &authEncryptionProtocol.Packet{
		AssocData: AssocData,
		IV:        []rune("БОБ_НЕМНОГО_ПЬЯН"),
		Message:   []rune(InputsArray[0]),
		MAC:       make([]rune, 0),
	}

	ab := append([]rune{}, AssocData.MType[:]...)
	ab = append(ab, AssocData.Sender[:]...)
	ab = append(ab, AssocData.Reciever[:]...)
	ab = append(ab, AssocData.Transmission[:]...)
	ab = append(ab, []rune("___")...)

	fmt.Println(string(packet.Message))

	l := eax.ForwardEaxCfb(*packet, eax.CFB.Forward(ab, []rune(SEC), "СЕАНСОВЫЙ_КЛЮЧИК", -1), "СЕАНСОВЫЙ_КЛЮЧИК", SEC, 0)
	fmt.Println(l.String())

	l2 := eax.InvertEaxCfb(l, "СЕАНСОВЫЙ_КЛЮЧИК", SEC, 0)
	fmt.Println(l2.String())

	ad := ADPrepare(AssocdataArray[3], "_____")

	fmt.Println(ad.String())

	protocol := authEncryptionProtocol.NewProtocol(telegraphAlphabet)

	protocol.NewConnection(ad, "СЕАНСОВЫЙ_КЛЮЧИК", "СЕМИХАТОВ_КВАНТЫ")
	fmt.Println("-----------------------Первый-----------------------")
	msgSend1 := protocol.Send(InputsArray[0], "")
	fmt.Println(len(msgSend1))
	msgSend1 = protocol.Invert(msgSend1, 317)
	msgRecieve1 := protocol.Recieve(msgSend1)
	fmt.Println(msgRecieve1.String())

	fmt.Println("-----------------------Второй-----------------------")
	msgSend2 := protocol.Send(InputsArray[1], "")
	fmt.Println(len(msgSend2))
	msgRecieve2 := protocol.Recieve(msgSend2)
	fmt.Println(msgRecieve2.String())

	fmt.Println("-----------------------Четвертый-----------------------")
	msgSend4 := protocol.Send(InputsArray[3], "")
	fmt.Println(len(msgSend4))
	msgSend4 = protocol.Invert(msgSend4, 12)
	msgRecieve4 := protocol.Recieve(msgSend4)
	fmt.Println(msgRecieve4.String())

	fmt.Println("-----------------------Пятый-----------------------")
	msgSend6 := protocol.Send(InputsArray[5], "")
	fmt.Println(len(msgSend6))
	msgRecieve6 := protocol.Recieve(msgSend6)
	fmt.Println(msgRecieve6.String())

}

func readFileByLines(filepath string, processor func(string)) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processor(scanner.Text())
	}
	return scanner.Err()
}

func ADPrepare(ad, extra string) authEncryptionProtocol.AssData {
	a := [4]string(strings.Fields(strings.ReplaceAll(ad, `"`, "")))

	return authEncryptionProtocol.NewAssData(a[0], a[1], a[2], a[3], extra)
}
