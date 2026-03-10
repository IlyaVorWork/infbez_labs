package codeRandomGenerator_test

import (
	generator "infbez_labs/internal/codeRandomGenerator"
	"math"
	"math/rand"
	"strings"
	"testing"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	numReplications    = 200
	bitsPerReplication = 4000
)

var (
	lfsr = generator.NewLFSR(alphabet)

	set = [][][]int{
		{
			generator.TapsToBin([]int{19, 18}),
			generator.TapsToBin([]int{18, 7}),
			generator.TapsToBin([]int{17, 3}),
		},
		{
			generator.TapsToBin([]int{19, 18}),
			generator.TapsToBin([]int{18, 7}),
			generator.TapsToBin([]int{16, 14, 13, 11}),
		},
		{
			generator.TapsToBin([]int{19, 18}),
			generator.TapsToBin([]int{18, 7}),
			generator.TapsToBin([]int{15, 13, 12, 10}),
		},
		{
			generator.TapsToBin([]int{19, 18}),
			generator.TapsToBin([]int{18, 7}),
			generator.TapsToBin([]int{14, 5, 3, 1}),
		},
	}

	seed = time.Now().UnixNano()
	rnd  = rand.New(rand.NewSource(seed))
)

func TestNIST_FrequencyMonobit(t *testing.T) {
	var (
		xValues = make([]float64, numReplications)
		passed  = 0
	)

	for rep := 0; rep < numReplications; rep++ {
		repSeed := randomBlock(alphabet.AlphabetArr, 16, rnd)
		bits := generateBits(lfsr, repSeed, set, bitsPerReplication)

		n1 := 0
		for _, b := range bits {
			n1 += b
		}
		n0 := bitsPerReplication - n1
		n := float64(bitsPerReplication)
		p1 := float64(n1) / n
		p0 := float64(n0) / n

		s := math.Sqrt(n*p1*p0) / n
		x := math.Abs(p0-p1) / s

		xValues[rep] = x
		if x < 3.0 {
			passed++
		}
	}

	buildHistogramMonobit(xValues)

	buckets := []struct {
		lo, hi float64
		label  string
	}{
		{0, 0.5, "[0.0, 0.5)"},
		{0.5, 1.0, "[0.5, 1.0)"},
		{1.0, 1.5, "[1.0, 1.5)"},
		{1.5, 2.0, "[1.5, 2.0)"},
		{2.0, 2.5, "[2.0, 2.5)"},
		{2.5, 3.0, "[2.5, 3.0)"},
		{3.0, 1e9, "[3.0, ∞)  "},
	}

	counts := make([]int, len(buckets))
	for _, x := range xValues {
		for i, b := range buckets {
			if x >= b.lo && x < b.hi {
				counts[i]++
				break
			}
		}
	}

	t.Logf("=== NIST Частотный монобитный тест (%d репликаций по %d бит) ===", numReplications, bitsPerReplication)
	t.Logf("Гистограмма значений x = |p0-p1|/s:")
	for i, b := range buckets {
		bar := strings.Repeat("█", counts[i])
		t.Logf("  %s : %3d %s", b.label, counts[i], bar)
	}
	passFrac := float64(passed) / float64(numReplications) * 100.0
	t.Logf("Прошли тест (x < 3): %d / %d = %.1f%%", passed, numReplications, passFrac)

}

func TestNIST_LongestRunOfOnes(t *testing.T) {
	var (
		mValues = make([]int, numReplications)
		passed  = 0
	)

	for rep := 0; rep < numReplications; rep++ {
		repSeed := randomBlock(alphabet.AlphabetArr, 16, rnd)
		bits := generateBits(lfsr, repSeed, set, bitsPerReplication)

		maxRun := 0
		curRun := 0
		for _, b := range bits {
			if b == 1 {
				curRun++
				if curRun > maxRun {
					maxRun = curRun
				}
			} else {
				curRun = 0
			}
		}
		mValues[rep] = maxRun
		if maxRun >= 10 && maxRun <= 15 {
			passed++
		}
	}

	buildHistogramLongestRun(mValues)

	buckets := []struct {
		lo, hi int
		label  string
	}{
		{0, 5, "[0,  5)  "},
		{5, 8, "[5,  8)  "},
		{8, 10, "[8, 10)  "},
		{10, 13, "[10,13)  "},
		{13, 16, "[13,16)✓ "},
		{16, 20, "[16,20)  "},
		{20, 1 << 30, "[20, ∞)  "},
	}

	counts := make([]int, len(buckets))
	for _, m := range mValues {
		for i, b := range buckets {
			if m >= b.lo && m < b.hi {
				counts[i]++
				break
			}
		}
	}

	t.Logf("=== NIST Тест длины серии единиц (%d репликаций по %d бит) ===", numReplications, bitsPerReplication)
	t.Logf("Гистограмма максимальной длины серии единиц m:")
	for i, b := range buckets {
		bar := strings.Repeat("█", counts[i])
		t.Logf("  %s : %3d %s", b.label, counts[i], bar)
	}
	passFrac := float64(passed) / float64(numReplications) * 100.0
	t.Logf("Прошли тест (10 ≤ m ≤ 15): %d / %d = %.1f%%", passed, numReplications, passFrac)
}

func randomBlock(alphabet []rune, length int, rnd *rand.Rand) string {
	r := make([]rune, length)
	for i := 0; i < length; i++ {
		r[i] = alphabet[rnd.Intn(len(alphabet))]
	}
	return string(r)
}

func generateBits(lfsr *generator.LFSR, seed string, tapsSet [][][]int, bitsNeeded int) []int {
	bits := make([]int, 0, bitsNeeded+4)
	var state [][][]int
	flag := "up"

	for len(bits) < bitsNeeded {
		var out string
		out, state = lfsr.WrapCAsLfsrNext(flag, state, seed, tapsSet)
		flag = "down"
		runes := []rune(out)
		for i := 0; i+3 < len(runes); i += 4 {
			block := string(runes[i : i+4])
			binBlock := alphabet.BlockToBin(block)
			bits = append(bits, binBlock...)
		}
	}
	return bits[:bitsNeeded]
}

func buildHistogramMonobit(xValues []float64) {
	values := make(plotter.Values, len(xValues))
	for i, v := range xValues {
		values[i] = v
	}

	p := plot.New()
	p.Title.Text = "NIST: Частотный монобитный тест"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "Частота"

	hist, _ := plotter.NewHist(values, 30)
	hist.FillColor = nil

	p.Add(hist)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, "../../sources/NIST_MonobitFrequency.png"); err != nil {
		panic(err)
	}
}

func buildHistogramLongestRun(mValues []int) {
	values := make(plotter.Values, len(mValues))
	for i, v := range mValues {
		values[i] = float64(v)
	}

	p := plot.New()
	p.Title.Text = "NIST: Длина серии единиц"
	p.X.Label.Text = "m (макс. длина серии единиц)"
	p.Y.Label.Text = "Частота"

	hist, _ := plotter.NewHist(values, 20)
	hist.FillColor = nil

	p.Add(hist)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, "../../sources/NIST_LongestRunOnes.png"); err != nil {
		panic(err)
	}
}
