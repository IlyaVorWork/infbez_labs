package authEncryptionProtocol_test

import (
	"crypto/rand"
	"fmt"
	"infbez_labs/internal/alphabet"
	"infbez_labs/internal/authEncryptionProtocol"
	"math/big"
	"testing"
)

func TestOracleAttacks(t *testing.T) {

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

	t.Run(fmt.Sprintf("Нарушение целостности"), func(t *testing.T) {
		XTST := protocol.PreparePacket(ADPrepare(AssocdataArray[1], ""), "УБЕЙТЕМЕНЯ", PT)

		BITS := protocol.Transmit(XTST)
		NEWBITS, _ := FlipRandomBitsCopy(BITS, 2)

		YTST := protocol.Recieve(NEWBITS)

		if XTST.String() == YTST.String() {
			t.Errorf("Failed Protocol test. Packets not match")
			return
		}
	})
}

func FlipRandomBitsCopy(data []byte, n int) ([]byte, error) {
	if len(data) == 0 || n <= 0 {
		out := make([]byte, len(data))
		copy(out, data)
		return out, nil
	}

	// Копируем исходный массив, чтобы не менять оригинал
	out := make([]byte, len(data))
	copy(out, data)

	totalBits := len(out) * 8

	for i := 0; i < n; i++ {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(totalBits)))
		if err != nil {
			return nil, err
		}

		bitIndex := int(r.Int64())
		byteIndex := bitIndex / 8
		bitOffset := bitIndex % 8

		// Переворот бита
		out[byteIndex] ^= 1 << bitOffset
	}

	return out, nil
}
