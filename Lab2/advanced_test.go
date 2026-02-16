package sponge

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	trithemius "infbez_labs/Lab1"
	"math/rand"
	"testing"
	"time"
)

var TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")

func randomBlock(alphabet []rune, length int, rnd *rand.Rand) string {
	r := make([]rune, length)
	for i := 0; i < length; i++ {
		r[i] = alphabet[rnd.Intn(len(alphabet))]
	}
	return string(r)
}

func indexOf(runes []rune, r rune) int {
	for i, v := range runes {
		if v == r {
			return i
		}
	}
	return -1
}

func smallChangeRune(alpha []rune, r rune, rnd *rand.Rand) rune {
	idx := indexOf(alpha, r)
	if idx == -1 {
		return r
	}
	n := len(alpha)
	newIdx := (idx + 1 + n) % n
	return alpha[newIdx]
}

func makeSmallChangeOnBlock(alpha []rune, block string, rnd *rand.Rand) string {
	runes := []rune(block)
	if len(runes) == 0 {
		return block
	}
	i := rnd.Intn(len(runes))
	runes[i] = smallChangeRune(alpha, runes[i], rnd)
	return string(runes)
}

func Test_SpongeHash_Random(t *testing.T) {
	const (
		numInputs = 500
		blockLen  = 16
	)
	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(seed))
	alphabet := trithemius.NewAlphabet(TelegraphAlphabet)

	inputs := make([]string, 0, numInputs)
	for i := 0; i < numInputs; i++ {
		inputs = append(inputs, randomBlock(TelegraphAlphabet, blockLen, rnd))
	}

	fmt.Println(inputs)

	frequencies := make([]int, len(TelegraphAlphabet))

	t.Run("SmallChange", func(t *testing.T) {
		for _, p := range inputs {
			hash1 := SpongeHash(p, *alphabet)
			changedInput := makeSmallChangeOnBlock(TelegraphAlphabet, p, rnd)
			hash2 := SpongeHash(changedInput, *alphabet)

			diff := alphabet.SubTxt(hash1, hash2)

			for _, v := range []rune(diff) {
				frequencies[alphabet.GetKeyByChar(string(v))] += 1
			}

			t.Logf("Small Change test: input=%s, changedInput=%s, hash1=%s, hash2=%s, hashDiff=%s", p, changedInput, hash1, hash2, diff)
		}

		t.Logf("Small Change test: frequencies=%v", frequencies)

		BuildHistogram(frequencies, TelegraphAlphabet)
	})
}

func BuildHistogram(frequencies []int, alphabet []rune) {
	values := make(plotter.Values, len(frequencies))
	for i, v := range frequencies {
		values[i] = float64(v)
	}

	p := plot.New()
	p.Title.Text = "Столбчатая гистограмма частот вхождения символов алфавита в разности хэшей"

	barWidth := vg.Points(10)

	bars, err := plotter.NewBarChart(values, barWidth)
	if err != nil {
		panic(err)
	}

	p.Add(bars)

	labels := make([]string, len(alphabet))
	for i, r := range alphabet {
		labels[i] = string(r)
	}
	labels = append(labels[len(labels)-1:], labels[:len(labels)-1]...)

	p.NominalX(labels...)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, "barchart.png"); err != nil {
		panic(err)
	}
}
