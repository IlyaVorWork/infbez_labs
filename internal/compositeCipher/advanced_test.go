package compositeCipher_test

import (
	"crypto/rand"
	"fmt"
	"infbez_labs/internal/alphabet"
	"infbez_labs/internal/codeRandomGenerator"
	"infbez_labs/internal/compositeCipher"
	"infbez_labs/internal/core"
	"testing"
)

const (
	tests = 100
	bits  = 80
	key   = "САМЙ_ЛУЧШИЙ_КЛЮЧ"
)

func hammingDistance(a, b []int) int {
	dist := 0

	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}

	return dist
}

func flipBit(bits []int, pos int) []int {
	out := make([]int, len(bits))
	copy(out, bits)

	if out[pos] == 0 {
		out[pos] = 1
	} else {
		out[pos] = 0
	}

	return out
}

func randomBits(size int) []int {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	bits := make([]int, size)
	for i := 0; i < size; i++ {
		bits[i] = int(bytes[i] & 1)
	}
	return bits
}

func printHistogram(sums []int) {
	maxval := 0
	for _, v := range sums {
		if v > maxval {
			maxval = v
		}
	}

	for i, v := range sums {
		avg := float64(v) / float64(tests)
		barLen := int(avg)

		fmt.Printf("%2d | ", i)
		for j := 0; j < barLen; j++ {
			fmt.Print("#")
		}
		fmt.Printf(" %.2f\n", avg)
	}
}

func TestCompositeCipher_Avalanche(t *testing.T) {

	sums := make([]int, bits)

	Alphabet := alphabet.NewAlphabet(alphabet.TelegraphAlphabet)
	Sblock := core.NewSBlockSTM(Alphabet)
	Pblock := core.NewPBlock(Alphabet)
	LFSR := codeRandomGenerator.NewLFSR(Alphabet)
	SPNet := compositeCipher.NewSPNet(Alphabet, Sblock, Pblock, LFSR)

	for i := 0; i < tests; i++ {

		plainBits := randomBits(bits)
		plainText := Pblock.Text80BitTo16Char(plainBits)
		cipherText := SPNet.FrwSPNet(plainText, key, 8)

		for j := 0; j < bits; j++ {

			plainBitsFlipped := flipBit(plainBits, j)
			plainTextFlipped := Pblock.Text80BitTo16Char(plainBitsFlipped)
			cipherTextFlipped := SPNet.FrwSPNet(plainTextFlipped, key, 8)
			d := hammingDistance(Pblock.Text16CharTo80Bit(cipherText), Pblock.Text16CharTo80Bit(cipherTextFlipped))

			sums[j] += d
		}
	}

	printHistogram(sums)
}
