package authEncryptionProtocol_test

import (
	"bufio"
	"fmt"
	"infbez_labs/internal/alphabet"
	"infbez_labs/internal/authEncryptionProtocol"
	"os"
	"strings"
	"testing"
)

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

func TestProtocol(t *testing.T) {

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

	telegraphAlphabet := alphabet.NewAlphabet(alphabet.TelegraphAlphabet)
	//cBlock := core.NewCBlock(telegraphAlphabet)
	//hasher := hash.NewHasher(telegraphAlphabet, cBlock)
	protocol := authEncryptionProtocol.NewProtocol(telegraphAlphabet)

	t.Run(fmt.Sprintf("Целое число блоков"), func(t *testing.T) {
		IN := InputsArray[0]
		if len(IN) != 368 {
			t.Errorf("Failed Protocol test. Wrong input initial length want=%d, got=%d", 368, len(IN))
			return
		}

		BIN1, _ := telegraphAlphabet.MessageToBin([]rune(IN))
		if len(BIN1) != 1840 {
			t.Errorf("Failed Protocol test. Wrong bin message length want=%d, got=%d", 1840, len(BIN1))
			return
		}

		INTER := protocol.PadMessage([]rune(IN))
		if len(INTER) != 368 {
			t.Errorf("Failed Protocol test. Wrong input with pad length want=%d, got=%d", 368, len(INTER))
			return
		}

		BIN2, _ := telegraphAlphabet.MessageToBin(INTER)
		if len(BIN2) != 1840 {
			t.Errorf("Failed Protocol test. Wrong bin message with pad length want=%d, got=%d", 1840, len(BIN2))
			return
		}

		OUT := protocol.UnpadMessage(INTER)
		if len(OUT) != 368 {
			t.Errorf("Failed Protocol test. Wrong input after unpad length want=%d, got=%d", 368, len(OUT))
			return
		}

		BIN3, _ := telegraphAlphabet.MessageToBin(OUT)
		if len(BIN3) != 1840 {
			t.Errorf("Failed Protocol test. Wrong bin message after unpad length want=%d, got=%d", 1840, len(BIN3))
			return
		}
	})

	t.Run(fmt.Sprintf("НЕ целое число блоков"), func(t *testing.T) {
		IN := InputsArray[1]
		if len(IN) != 3079 {
			t.Errorf("Failed Protocol test. Wrong input initial length want=%d, got=%d", 3079, len(IN))
			return
		}

		BIN1, _ := telegraphAlphabet.MessageToBin([]rune(IN))
		if len(BIN1) != 15395 {
			t.Errorf("Failed Protocol test. Wrong bin message length want=%d, got=%d", 15395, len(BIN1))
			return
		}

		hasPad, numBlocks, padLength := protocol.CheckPadding(BIN1)
		if hasPad || numBlocks != 0 || padLength != 0 {
			t.Errorf("Failed Protocol test. Wrong padding want=%t,%d,%d , got=%t,%d,%d", false, 0, 0, hasPad, numBlocks, padLength)
			return
		}

		INTER := protocol.PadMessage([]rune(IN))
		if len(INTER) != 3088 {
			t.Errorf("Failed Protocol test. Wrong input with pad length want=%d, got=%d", 3088, len(INTER))
			return
		}

		BIN2, _ := telegraphAlphabet.MessageToBin(INTER)
		if len(BIN2) != 15440 {
			t.Errorf("Failed Protocol test. Wrong bin message with pad length want=%d, got=%d", 15440, len(BIN2))
			return
		}

		hasPad, numBlocks, padLength = protocol.CheckPadding(BIN2)
		if !hasPad || numBlocks != 0 || padLength != 0 {
			t.Errorf("Failed Protocol test. Wrong padding want=%t,%d,%d , got=%t,%d,%d", true, 193, 45, hasPad, numBlocks, padLength)
			return
		}

		OUT := protocol.UnpadMessage(INTER)
		if len(OUT) != 3079 {
			t.Errorf("Failed Protocol test. Wrong input after unpad length want=%d, got=%d", 3079, len(OUT))
			return
		}

		BIN3, _ := telegraphAlphabet.MessageToBin(OUT)
		if len(BIN3) != 15395 {
			t.Errorf("Failed Protocol test. Wrong bin message after unpad length want=%d, got=%d", 15395, len(BIN3))
			return
		}
	})

	t.Run(fmt.Sprintf("Целое число блоков, заканчивается на подложку"), func(t *testing.T) {
		IN := InputsArray[2]
		if len(IN) != 519 {
			t.Errorf("Failed Protocol test. Wrong input initial length want=%d, got=%d", 519, len(IN))
			return
		}

		BIN1, _ := telegraphAlphabet.MessageToBin([]rune(IN))
		if len(BIN1) != 2595 {
			t.Errorf("Failed Protocol test. Wrong bin message length want=%d, got=%d", 2595, len(BIN1))
			return
		}

		IN2 := protocol.PadMessage([]rune(InputsArray[2]))
		if len(IN2) != 528 {
			t.Errorf("Failed Protocol test. Wrong input initial length want=%d, got=%d", 528, len(IN))
			return
		}

		BIN2, _ := telegraphAlphabet.MessageToBin(IN2)
		if len(BIN2) != 2640 {
			t.Errorf("Failed Protocol test. Wrong bin message length want=%d, got=%d", 2640, len(BIN2))
			return
		}

		hasPad, numBlocks, padLength := protocol.CheckPadding(BIN2)
		if !hasPad || numBlocks != 0 || padLength != 0 {
			t.Errorf("Failed Protocol test. Wrong padding want=%t,%d,%d , got=%t,%d,%d", true, 33, 45, hasPad, numBlocks, padLength)
			return
		}

		INTER := protocol.PadMessage(IN2)
		if len(INTER) != 544 {
			t.Errorf("Failed Protocol test. Wrong input with pad length want=%d, got=%d", 544, len(INTER))
			return
		}

		BIN3, _ := telegraphAlphabet.MessageToBin(INTER)
		if len(BIN3) != 2720 {
			t.Errorf("Failed Protocol test. Wrong bin message with pad length want=%d, got=%d", 2720, len(BIN3))
			return
		}

		hasPad, numBlocks, padLength = protocol.CheckPadding(BIN3)
		if !hasPad || numBlocks != 0 || padLength != 0 {
			t.Errorf("Failed Protocol test. Wrong padding want=%t,%d,%d , got=%t,%d,%d", true, 34, 80, hasPad, numBlocks, padLength)
			return
		}

		OUT := protocol.UnpadMessage(INTER)
		if len(OUT) != 528 {
			t.Errorf("Failed Protocol test. Wrong input after unpad length want=%d, got=%d", 528, len(OUT))
			return
		}

		BIN4, _ := telegraphAlphabet.MessageToBin(OUT)
		if len(BIN4) != 2640 {
			t.Errorf("Failed Protocol test. Wrong bin message after unpad length want=%d, got=%d", 2640, len(BIN4))
			return
		}

		if string(OUT) != string(IN2) {
			t.Errorf("Failed Protocol test. Wrong unpadded message")
			return
		}

		if string(protocol.UnpadMessage(OUT)) != IN {
			t.Errorf("Failed Protocol test. Wrong unpadded message")
			return
		}
	})

	t.Run(fmt.Sprintf("Оставшиеся тесты"), func(t *testing.T) {
		IN := InputsArray[3]
		if len(IN) != 122 {
			t.Errorf("Failed Protocol test. Wrong input initial length want=%d, got=%d", 122, len(IN))
			return
		}

		BIN1, _ := telegraphAlphabet.MessageToBin([]rune(IN))
		if len(BIN1) != 138 {
			t.Errorf("Failed Protocol test. Wrong bin message length want=%d, got=%d", 138, len(BIN1))
			return
		}

		INTER := protocol.PadMessage([]rune(IN))
		if len(INTER) != 48 {
			t.Errorf("Failed Protocol test. Wrong input with pad length want=%d, got=%d", 48, len(INTER))
			return
		}

		BIN3, _ := telegraphAlphabet.MessageToBin(INTER)
		if len(BIN3) != 240 {
			t.Errorf("Failed Protocol test. Wrong bin message with pad length want=%d, got=%d", 240, len(BIN3))
			return
		}

		hasPad, numBlocks, padLength := protocol.CheckPadding(BIN3)
		if !hasPad || numBlocks != 0 || padLength != 0 {
			t.Errorf("Failed Protocol test. Wrong padding want=%t,%d,%d , got=%t,%d,%d", true, 3, 102, hasPad, numBlocks, padLength)
			return
		}

		OUT := protocol.UnpadMessage(INTER)
		if len(OUT) != 528 {
			t.Errorf("Failed Protocol test. Wrong input after unpad length want=%d, got=%d", 528, len(OUT))
			return
		}

		if string(OUT) != IN {
			t.Errorf("Failed Protocol test. Wrong unpadded message")
			return
		}
	})

	t.Run(fmt.Sprintf("Передача пакетов"), func(t *testing.T) {
		XTST := protocol.PreparePacket(ADPrepare(AssocdataArray[1], ""), "КОЛЕСО", InputsArray[1])
		YTST := protocol.Recieve(protocol.Transmit(XTST))
		if XTST.String() != YTST.String() {
			t.Errorf("Failed Protocol test. Packets not match")
			return
		}
	})
}
