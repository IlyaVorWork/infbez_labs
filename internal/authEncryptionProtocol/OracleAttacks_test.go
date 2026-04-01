package authEncryptionProtocol_test

import (
	"encoding/binary"
	"fmt"
	"infbez_labs/internal/alphabet"
	"infbez_labs/internal/authEncryptionProtocol"
	"math/rand"
	"testing"
)

func TestConsistancyOracleAttacks(t *testing.T) {

	PT := "ЯТАКБОЛЬШЕНЕМОГУ"
	AssocdataArray := make([]string, 0)

	err := readFileByLines("sources/ad.txt", func(line string) {
		AssocdataArray = append(AssocdataArray, line)
	})
	if err != nil {
		return
	}

	telegraphAlphabet := alphabet.NewAlphabet(alphabet.TelegraphAlphabet)
	protocol := authEncryptionProtocol.NewProtocol(telegraphAlphabet)
	protocol.NewConnection(ADPrepare(AssocdataArray[1], "ЭКСТР"), "СЕАНСОВЫЙ_КЛЮЧИК", "СЕМИХАТОВ_КВАНТЫ")

	t.Run(fmt.Sprintf("Нарушение целостности"), func(t *testing.T) {

		XTST := protocol.Send(PT, "")

		FlipRandomBits(XTST, 2)

		protocol.Recieve(XTST)
	})
}

func TestReplayOracleAttacks(t *testing.T) {

	PT := "ЯТАКБОЛЬШЕНЕМОГУ"
	AssocdataArray := make([]string, 0)

	err := readFileByLines("sources/ad.txt", func(line string) {
		AssocdataArray = append(AssocdataArray, line)
	})
	if err != nil {
		return
	}

	telegraphAlphabet := alphabet.NewAlphabet(alphabet.TelegraphAlphabet)
	protocol := authEncryptionProtocol.NewProtocol(telegraphAlphabet)
	protocol.NewConnection(ADPrepare(AssocdataArray[1], "ЭКСТР"), "СЕАНСОВЫЙ_КЛЮЧИК", "СЕМИХАТОВ_КВАНТЫ")

	t.Run(fmt.Sprintf("Replay-атака"), func(t *testing.T) {

		XTST := protocol.Send(PT, "")

		t.Logf("Первая отправка")
		protocol.Recieve(XTST)

		t.Logf("Вторая отправка")
		protocol.Recieve(XTST)
	})
}

func TestConfidentialityOracleAttacks(t *testing.T) {

	PT1 := "ЯТАКБОЛЬШЕНЕМОГУ"
	PT2 := "УЖЕМОГУХИХИХАХАХ"
	AssocdataArray := make([]string, 0)

	err := readFileByLines("sources/ad.txt", func(line string) {
		AssocdataArray = append(AssocdataArray, line)
	})
	if err != nil {
		return
	}

	telegraphAlphabet := alphabet.NewAlphabet(alphabet.TelegraphAlphabet)
	protocol := authEncryptionProtocol.NewProtocol(telegraphAlphabet)
	protocol.NewConnection(ADPrepare(AssocdataArray[1], "ЭКСТР"), "СЕАНСОВЫЙ_КЛЮЧИК", "СЕМИХАТОВ_КВАНТЫ")

	t.Run(fmt.Sprintf("IV-Replay-атака"), func(t *testing.T) {

		victimBits := protocol.Send(PT1, "ОЧЕНЬСЕКРЕТНЫЙИВ")
		victimPacket := protocol.Response(victimBits)

		attackBits := protocol.Send(PT2, "ОЧЕНЬСЕКРЕТНЫЙИВ")
		attackPacket := protocol.Response(attackBits)
		
		pt1XorPt2 := telegraphAlphabet.BlockXOR(victimPacket.Message, attackPacket.Message)
		pt2 := telegraphAlphabet.BlockXOR(pt1XorPt2, []rune(PT1))

		t.Log(PT2, string(pt2))
	})
}

func FlipRandomBits(data []byte, n int) []byte {
	if len(data) == 0 || n <= 0 {
		return nil
	}

	totalBits := len(data) * 8

	for i := 0; i < n; i++ {
		bitIndex, err := cryptoRandInt(totalBits)
		if err != nil {
			return nil
		}

		byteIndex := bitIndex / 8
		bitOffset := bitIndex % 8

		// Переворачиваем бит через XOR
		data[byteIndex] ^= 1 << bitOffset
	}

	return nil
}

func cryptoRandInt(max int) (int, error) {
	if max <= 0 {
		return 0, nil
	}

	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}

	return int(binary.LittleEndian.Uint64(b[:]) % uint64(max)), nil
}
