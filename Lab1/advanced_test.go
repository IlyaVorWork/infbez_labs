package trithemius

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")

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
	var delta int
	if n <= 2 {
		delta = 1
	} else {
		if rnd.Intn(2) == 0 {
			delta = -1
		} else {
			delta = 1
		}
	}
	newIdx := (idx + delta + n) % n
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

// Число позиций, где строки (руны) различаются
func diffPositions(a, b string) int {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		if len(ar) > len(br) {
			diff := len(ar) - len(br)
			// сравнить общую часть
			for i := 0; i < len(br); i++ {
				if ar[i] != br[i] {
					diff++
				}
			}
			return diff
		}
		diff := len(br) - len(ar)
		for i := 0; i < len(ar); i++ {
			if ar[i] != br[i] {
				diff++
			}
		}
		return diff
	}
	diff := 0
	for i := 0; i < len(ar); i++ {
		if ar[i] != br[i] {
			diff++
		}
	}
	return diff
}

func randomBlock(alphabet []rune, length int, rnd *rand.Rand) string {
	r := make([]rune, length)
	for i := 0; i < length; i++ {
		r[i] = alphabet[rnd.Intn(len(alphabet))]
	}
	return string(r)
}

// Циклически сдвинуть строку (руны) вправо на сдвиг
func rotateRightRunes(s string, shift int) string {
	rs := []rune(s)
	n := len(rs)
	if n == 0 {
		return s
	}
	k := ((shift % n) + n) % n
	if k == 0 {
		return s
	}
	out := make([]rune, n)
	for i := 0; i < n; i++ {
		out[(i+k)%n] = rs[i]
	}
	return string(out)
}

// Проверяет, состоят ли две строки из одного и того же набора рун
func outputsHaveSameMultiset(a, b string) bool {
	ar := []rune(a)
	br := []rune(b)
	if len(ar) != len(br) {
		return false
	}
	counts := make(map[rune]int)
	for _, r := range ar {
		counts[r]++
	}
	for _, r := range br {
		counts[r]--
	}
	for _, v := range counts {
		if v != 0 {
			return false
		}
	}
	return true
}

// ----------------------------------------------------------------

func Test_STrithemius_Input(t *testing.T) {
	const (
		numInputs = 50
		numKeys   = 40
		blockLen  = 4
		keyLen    = 16
	)
	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(seed))
	trithemius := NewTrithemius(TelegraphAlphabet)

	inputs := make([]string, 0, numInputs)
	keys := make([]string, 0, numKeys)
	for i := 0; i < numInputs; i++ {
		inputs = append(inputs, randomBlock(TelegraphAlphabet, blockLen, rnd))
	}
	for i := 0; i < numKeys; i++ {
		keys = append(keys, randomBlock(TelegraphAlphabet, keyLen, rnd))
	}

	fmt.Println(inputs)
	fmt.Println(keys)

	t.Run("SmallChange", func(t *testing.T) {
		type Result struct {
			input string
			key   string
			c1    string
			c2    string
			diff  int
		}

		var noneChangeList, oneChangeList, twoChangeList, threeChangeList, fourChangeList []Result

		for _, p := range inputs {
			for _, k := range keys {
				pPrime := makeSmallChangeOnBlock(TelegraphAlphabet, p, rnd)
				c1 := trithemius.EncodeSTrithemius(p, k)
				c2 := trithemius.EncodeSTrithemius(pPrime, k)
				diff := diffPositions(c1, c2)

				res := Result{p, k, c1, c2, diff}
				if diff == 4 {
					fourChangeList = append(fourChangeList, res)
				} else if diff == 3 {
					threeChangeList = append(threeChangeList, res)
				} else if diff == 2 {
					twoChangeList = append(twoChangeList, res)
				} else if diff == 1 {
					oneChangeList = append(oneChangeList, res)
				} else {
					noneChangeList = append(noneChangeList, res)
				}
			}
		}

		total := len(inputs) * len(keys)
		fourChangeCount := len(fourChangeList)
		threeChangeCount := len(threeChangeList)
		twoChangeCount := len(twoChangeList)
		oneChangeCount := len(oneChangeList)
		noneChangeCount := len(noneChangeList)

		t.Logf("Всего тестов: %d;\n\t4 изменений=%d (%.2f%%);\n\t3 изменений=%d (%.2f%%);\n\t2 изменений=%d (%.2f%%);\n\t1 изменение=%d (%.2f%%);\n\t0 изменений=%d (%.2f%%)",
			total,
			fourChangeCount, 100.0*float64(fourChangeCount)/float64(total),
			threeChangeCount, 100.0*float64(threeChangeCount)/float64(total),
			twoChangeCount, 100.0*float64(twoChangeCount)/float64(total),
			oneChangeCount, 100.0*float64(oneChangeCount)/float64(total),
			noneChangeCount, 100.0*float64(noneChangeCount)/float64(total),
		)

		printExamples := func(title string, list []Result) {
			if len(list) == 0 {
				return
			}
			t.Logf("Список тестов. %s:", title)
			n := len(list)

			for i := 0; i < n; i++ {
				r := list[i]
				t.Logf("P=%q, P'=%q, K=%q => C1=%q, C2=%q, diff=%d",
					r.input, makeSmallChangeOnBlock(TelegraphAlphabet, r.input, rnd), r.key, r.c1, r.c2, r.diff)
			}
		}
		printExamples("0 символов поменялось", noneChangeList)
		printExamples("1 символ поменялся", oneChangeList)
		printExamples("2 символа поменялось", twoChangeList)
		printExamples("3 символа поменялось", threeChangeList)
		printExamples("4 символов поменялось", fourChangeList)
	})

	t.Run("Rotate", func(t *testing.T) {
		type RResult struct {
			input  string
			rot    int
			key    string
			c1     string
			c2     string
			sameMS bool
		}
		var passList []RResult
		var failList []RResult

		for _, p := range inputs {
			for _, k := range keys {
				shift := 1 + rnd.Intn(blockLen-1)

				pPrime := rotateRightRunes(p, shift)
				c1 := trithemius.EncodeSTrithemius(p, k)
				c2 := trithemius.EncodeSTrithemius(pPrime, k)
				same := outputsHaveSameMultiset(c1, c2)

				res := RResult{p, shift, k, c1, c2, same}
				if same {
					failList = append(failList, res)
				} else {
					passList = append(passList, res)
				}
			}
		}

		total := len(inputs) * len(keys)
		passCount := len(passList)
		failCount := len(failList)

		t.Logf("Rotate test: total=%d, pass=%d (%.2f%%), fail=%d (%.2f%%)",
			total,
			passCount, 100.0*float64(passCount)/float64(total),
			failCount, 100.0*float64(failCount)/float64(total),
		)

		t.Logf("Примеры неудач:")
		n := len(failList)
		for i := 0; i < n; i++ {
			r := failList[i]
			t.Logf("P=%q, shift=%d, P'=%q, K=%q => C1=%q, C2=%q",
				r.input, r.rot, rotateRightRunes(r.input, r.rot), r.key, r.c1, r.c2)
		}
		t.Logf("Примеры удач:")
		n = len(passList)
		for i := 0; i < n; i++ {
			r := passList[i]
			t.Logf("P=%q, shift=%d, P'=%q, K=%q => C1=%q, C2=%q",
				r.input, r.rot, rotateRightRunes(r.input, r.rot), r.key, r.c1, r.c2)
		}

	})

	t.Run("Linear", func(t *testing.T) {
		type LResult struct {
			p1      string
			p2      string
			p3      string
			key     string
			c1      string
			c2      string
			c3      string
			c1plus2 string
			diff    int
		}
		var passList, failList []LResult

		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(inputs); j++ {
				if i == j {
					continue
				}
				p1 := inputs[i]
				p2 := inputs[j]
				p3 := trithemius.Alphabet.AddTxt(p1, p2)
				if p3 == p1 || p3 == p2 {
					continue
				}
				for _, k := range keys {
					c1 := trithemius.EncodeSTrithemius(p1, k)
					c2 := trithemius.EncodeSTrithemius(p2, k)
					c3 := trithemius.EncodeSTrithemius(p3, k)
					c1plus2 := trithemius.Alphabet.AddTxt(c1, c2)
					diff := diffPositions(c1plus2, c3)

					res := LResult{p1, p2, p3, k, c1, c2, c3, c1plus2, diff}
					if diff != 0 {
						passList = append(passList, res)
					} else {
						failList = append(failList, res)
					}
				}
			}
		}

		total := len(passList) + len(failList)
		passCount := len(passList)
		failCount := len(failList)

		t.Logf("Linear test: total=%d, pass=%d (%.2f%%), fail=%d (%.2f%%)",
			total,
			passCount, 100.0*float64(passCount)/float64(total),
			failCount, 100.0*float64(failCount)/float64(total),
		)

		if len(failList) > 0 {
			t.Logf("Примеры FAIL")
			n := min(len(failList), 100)
			for i := 0; i < n; i++ {
				r := failList[i]
				t.Logf("P1=%q, P2=%q, P3=%q, K=%q => C1=%q, C2=%q, C3=%q, C1+C2=%q, diff=%d",
					r.p1, r.p2, r.p3, r.key, r.c1, r.c2, r.c3, r.c1plus2, r.diff)
			}
		}
		if len(passList) > 0 {
			t.Logf("Примеры Pass")
			n := min(len(passList), 100)
			for i := 0; i < n; i++ {
				r := passList[i]
				t.Logf("P1=%q, P2=%q, P3=%q, K=%q => C1=%q, C2=%q, C3=%q, C1+C2=%q, diff=%d",
					r.p1, r.p2, r.p3, r.key, r.c1, r.c2, r.c3, r.c1plus2, r.diff)
			}
		}
	})

}

func Test_STrithemius_Key(t *testing.T) {
	const (
		numInputs = 50
		numKeys   = 40
		blockLen  = 4
		keyLen    = 16
	)
	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(seed))
	alphabet := NewTrithemius(TelegraphAlphabet)

	inputs := make([]string, 0, numInputs)
	keys := make([]string, 0, numKeys)
	for i := 0; i < numInputs; i++ {
		inputs = append(inputs, randomBlock(TelegraphAlphabet, blockLen, rnd))
	}
	for i := 0; i < numKeys; i++ {
		keys = append(keys, randomBlock(TelegraphAlphabet, keyLen, rnd))
	}

	fmt.Println(inputs)
	fmt.Println(keys)

	t.Run("SmallChange", func(t *testing.T) {
		type Result struct {
			input string
			key   string
			c1    string
			c2    string
			diff  int
		}

		var noneChangeList, oneChangeList, twoChangeList, threeChangeList, fourChangeList []Result

		for _, k := range keys {
			for _, p := range inputs {
				kPrime := makeSmallChangeOnBlock(TelegraphAlphabet, k, rnd)
				c1 := alphabet.EncodeSTrithemius(p, k)
				c2 := alphabet.EncodeSTrithemius(p, kPrime)
				diff := diffPositions(c1, c2)

				res := Result{p, k, c1, c2, diff}
				if diff == 4 {
					fourChangeList = append(fourChangeList, res)
				} else if diff == 3 {
					threeChangeList = append(threeChangeList, res)
				} else if diff == 2 {
					twoChangeList = append(twoChangeList, res)
				} else if diff == 1 {
					oneChangeList = append(oneChangeList, res)
				} else {
					noneChangeList = append(noneChangeList, res)
				}
			}
		}

		total := len(inputs) * len(keys)
		fourChangeCount := len(fourChangeList)
		threeChangeCount := len(threeChangeList)
		twoChangeCount := len(twoChangeList)
		oneChangeCount := len(oneChangeList)
		noneChangeCount := len(noneChangeList)

		t.Logf("Всего тестов: %d;\n\t4 изменений=%d (%.2f%%);\n\t3 изменений=%d (%.2f%%);\n\t2 изменений=%d (%.2f%%);\n\t1 изменение=%d (%.2f%%);\n\t0 изменений=%d (%.2f%%)",
			total,
			fourChangeCount, 100.0*float64(fourChangeCount)/float64(total),
			threeChangeCount, 100.0*float64(threeChangeCount)/float64(total),
			twoChangeCount, 100.0*float64(twoChangeCount)/float64(total),
			oneChangeCount, 100.0*float64(oneChangeCount)/float64(total),
			noneChangeCount, 100.0*float64(noneChangeCount)/float64(total),
		)

		printExamples := func(title string, list []Result) {
			if len(list) == 0 {
				return
			}
			t.Logf("Список тестов. %s:", title)
			n := len(list)

			for i := 0; i < n; i++ {
				r := list[i]
				t.Logf("P=%q, K=%q, K'=%q => C1=%q, C2=%q, diff=%d",
					r.input, r.key, makeSmallChangeOnBlock(TelegraphAlphabet, r.key, rnd), r.c1, r.c2, r.diff)
			}
		}
		printExamples("0 символов поменялось", noneChangeList)
		printExamples("1 символ поменялся", oneChangeList)
		printExamples("2 символа поменялось", twoChangeList)
		printExamples("3 символа поменялось", threeChangeList)
		printExamples("4 символов поменялось", fourChangeList)
	})

	t.Run("Rotate", func(t *testing.T) {
		type RResult struct {
			input  string
			rot    int
			key    string
			c1     string
			c2     string
			sameMS bool
		}
		var passList []RResult
		var failList []RResult

		for _, k := range keys {
			for _, p := range inputs {
				shift := 1 + rnd.Intn(blockLen-1)

				kPrime := rotateRightRunes(k, shift)
				c1 := alphabet.EncodeSTrithemius(p, k)
				c2 := alphabet.EncodeSTrithemius(p, kPrime)
				same := outputsHaveSameMultiset(c1, c2)

				res := RResult{p, shift, k, c1, c2, same}
				if same {
					failList = append(failList, res)
				} else {
					passList = append(passList, res)
				}
			}
		}

		total := len(inputs) * len(keys)
		passCount := len(passList)
		failCount := len(failList)

		t.Logf("Rotate test: total=%d, pass=%d (%.2f%%), fail=%d (%.2f%%)",
			total,
			passCount, 100.0*float64(passCount)/float64(total),
			failCount, 100.0*float64(failCount)/float64(total),
		)

		t.Logf("Примеры неудач:")
		n := len(failList)
		for i := 0; i < n; i++ {
			r := failList[i]
			t.Logf("P=%q, shift=%d, K=%q, K'=%q => C1=%q, C2=%q",
				r.input, r.rot, r.key, rotateRightRunes(r.key, r.rot), r.c1, r.c2)
		}
		t.Logf("Примеры удач:")
		n = len(passList)
		for i := 0; i < n; i++ {
			r := passList[i]
			t.Logf("P=%q, shift=%d, K=%q, K'=%q => C1=%q, C2=%q",
				r.input, r.rot, r.key, rotateRightRunes(r.key, r.rot), r.c1, r.c2)
		}

	})

	t.Run("Linear", func(t *testing.T) {
		type LResult struct {
			p       string
			k1      string
			k2      string
			k3      string
			c1      string
			c2      string
			c3      string
			c1plus2 string
			diff    int
		}
		var passList, failList []LResult

		for i := 0; i < len(keys); i++ {
			for j := 0; j < len(keys); j++ {
				if i == j {
					continue
				}
				k1 := keys[i]
				k2 := keys[j]
				k3 := alphabet.Alphabet.AddTxt(k1, k2)
				if k3 == k1 || k3 == k2 {
					continue
				}
				for _, p := range inputs {
					c1 := alphabet.EncodeSTrithemius(p, k1)
					c2 := alphabet.EncodeSTrithemius(p, k2)
					c3 := alphabet.EncodeSTrithemius(p, k3)
					c1plus2 := alphabet.Alphabet.AddTxt(c1, c2)
					diff := diffPositions(c1plus2, c3)

					res := LResult{p, k1, k2, k3, c1, c2, c3, c1plus2, diff}
					if diff != 0 {
						passList = append(passList, res)
					} else {
						failList = append(failList, res)
					}
				}
			}
		}

		total := len(passList) + len(failList)
		passCount := len(passList)
		failCount := len(failList)

		t.Logf("Linear test: total=%d, pass=%d (%.2f%%), fail=%d (%.2f%%)",
			total,
			passCount, 100.0*float64(passCount)/float64(total),
			failCount, 100.0*float64(failCount)/float64(total),
		)

		if len(failList) > 0 {
			t.Logf("Примеры FAIL")
			n := min(len(failList), 100)
			for i := 0; i < n; i++ {
				r := failList[i]
				t.Logf("P=%q, K1=%q, K2=%q, K3=%q => C1=%q, C2=%q, C3=%q, C1+C2=%q, diff=%d",
					r.p, r.k1, r.k2, r.k3, r.c1, r.c2, r.c3, r.c1plus2, r.diff)
			}
		}
		if len(passList) > 0 {
			t.Logf("Примеры Pass")
			n := min(len(passList), 100)
			for i := 0; i < n; i++ {
				r := passList[i]
				t.Logf("P1=%q, P2=%q, P3=%q, K=%q => C1=%q, C2=%q, C3=%q, C1+C2=%q, diff=%d",
					r.p, r.k1, r.k2, r.k3, r.c1, r.c2, r.c3, r.c1plus2, r.diff)
			}
		}
	})
}

func Test_MergeBlock_Input(t *testing.T) {
	const (
		numInputs = 50
		numKeys   = 40
		blockLen  = 4
		keyLen    = 16
	)
	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(seed))
	alphabet := NewTrithemius(TelegraphAlphabet)

	inputs := make([]string, 0, numInputs)
	keys := make([]string, 0, numKeys)
	for i := 0; i < numInputs; i++ {
		inputs = append(inputs, randomBlock(TelegraphAlphabet, blockLen, rnd))
	}
	for i := 0; i < numKeys; i++ {
		keys = append(keys, randomBlock(TelegraphAlphabet, keyLen, rnd))
	}

	fmt.Println(inputs)
	fmt.Println(keys)

	t.Run("SmallChange", func(t *testing.T) {
		type Result struct {
			input string
			key   string
			c1    string
			c2    string
			diff  int
		}

		var noneChangeList, oneChangeList, twoChangeList, threeChangeList, fourChangeList []Result

		for _, p := range inputs {
			for _, k := range keys {
				pPrime := makeSmallChangeOnBlock(TelegraphAlphabet, p, rnd)
				c1 := alphabet.EncodeMergeBlock(p, k)
				c2 := alphabet.EncodeMergeBlock(pPrime, k)
				diff := diffPositions(c1, c2)

				res := Result{p, k, c1, c2, diff}
				if diff == 4 {
					fourChangeList = append(fourChangeList, res)
				} else if diff == 3 {
					threeChangeList = append(threeChangeList, res)
				} else if diff == 2 {
					twoChangeList = append(twoChangeList, res)
				} else if diff == 1 {
					oneChangeList = append(oneChangeList, res)
				} else {
					noneChangeList = append(noneChangeList, res)
				}
			}
		}

		total := len(inputs) * len(keys)
		fourChangeCount := len(fourChangeList)
		threeChangeCount := len(threeChangeList)
		twoChangeCount := len(twoChangeList)
		oneChangeCount := len(oneChangeList)
		noneChangeCount := len(noneChangeList)

		t.Logf("Всего тестов: %d;\n\t4 изменений=%d (%.2f%%);\n\t3 изменений=%d (%.2f%%);\n\t2 изменений=%d (%.2f%%);\n\t1 изменение=%d (%.2f%%);\n\t0 изменений=%d (%.2f%%)",
			total,
			fourChangeCount, 100.0*float64(fourChangeCount)/float64(total),
			threeChangeCount, 100.0*float64(threeChangeCount)/float64(total),
			twoChangeCount, 100.0*float64(twoChangeCount)/float64(total),
			oneChangeCount, 100.0*float64(oneChangeCount)/float64(total),
			noneChangeCount, 100.0*float64(noneChangeCount)/float64(total),
		)

		printExamples := func(title string, list []Result) {
			if len(list) == 0 {
				return
			}
			t.Logf("Список тестов. %s:", title)
			n := len(list)

			for i := 0; i < n; i++ {
				r := list[i]
				t.Logf("P=%q, P'=%q, K=%q => C1=%q, C2=%q, diff=%d",
					r.input, makeSmallChangeOnBlock(TelegraphAlphabet, r.input, rnd), r.key, r.c1, r.c2, r.diff)
			}
		}
		printExamples("0 символов поменялось", noneChangeList)
		printExamples("1 символ поменялся", oneChangeList)
		printExamples("2 символа поменялось", twoChangeList)
		printExamples("3 символа поменялось", threeChangeList)
		printExamples("4 символов поменялось", fourChangeList)
	})

	t.Run("Rotate", func(t *testing.T) {
		type RResult struct {
			input  string
			rot    int
			key    string
			c1     string
			c2     string
			sameMS bool
		}
		var passList []RResult
		var failList []RResult

		for _, p := range inputs {
			for _, k := range keys {
				shift := 1 + rnd.Intn(blockLen-1)

				pPrime := rotateRightRunes(p, shift)
				c1 := alphabet.EncodeMergeBlock(p, k)
				c2 := alphabet.EncodeMergeBlock(pPrime, k)
				same := outputsHaveSameMultiset(c1, c2)

				res := RResult{p, shift, k, c1, c2, same}
				if same {
					failList = append(failList, res)
				} else {
					passList = append(passList, res)
				}
			}
		}

		total := len(inputs) * len(keys)
		passCount := len(passList)
		failCount := len(failList)

		t.Logf("Rotate test: total=%d, pass=%d (%.2f%%), fail=%d (%.2f%%)",
			total,
			passCount, 100.0*float64(passCount)/float64(total),
			failCount, 100.0*float64(failCount)/float64(total),
		)

		t.Logf("Примеры неудач:")
		n := len(failList)
		for i := 0; i < n; i++ {
			r := failList[i]
			t.Logf("P=%q, shift=%d, P'=%q, K=%q => C1=%q, C2=%q",
				r.input, r.rot, rotateRightRunes(r.input, r.rot), r.key, r.c1, r.c2)
		}
		t.Logf("Примеры удач:")
		n = len(passList)
		for i := 0; i < n; i++ {
			r := passList[i]
			t.Logf("P=%q, shift=%d, P'=%q, K=%q => C1=%q, C2=%q",
				r.input, r.rot, rotateRightRunes(r.input, r.rot), r.key, r.c1, r.c2)
		}

	})

	t.Run("Linear", func(t *testing.T) {
		type LResult struct {
			p1      string
			p2      string
			p3      string
			key     string
			c1      string
			c2      string
			c3      string
			c1plus2 string
			diff    int
		}
		var passList, failList []LResult

		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(inputs); j++ {
				if i == j {
					continue
				}
				p1 := inputs[i]
				p2 := inputs[j]
				p3 := alphabet.Alphabet.AddTxt(p1, p2)
				if p3 == p1 || p3 == p2 {
					continue
				}
				for _, k := range keys {
					c1 := alphabet.EncodeMergeBlock(p1, k)
					c2 := alphabet.EncodeMergeBlock(p2, k)
					c3 := alphabet.EncodeMergeBlock(p3, k)
					c1plus2 := alphabet.Alphabet.AddTxt(c1, c2)
					diff := diffPositions(c1plus2, c3)

					res := LResult{p1, p2, p3, k, c1, c2, c3, c1plus2, diff}
					if diff != 0 {
						passList = append(passList, res)
					} else {
						failList = append(failList, res)
					}
				}
			}
		}

		total := len(passList) + len(failList)
		passCount := len(passList)
		failCount := len(failList)

		t.Logf("Linear test: total=%d, pass=%d (%.2f%%), fail=%d (%.2f%%)",
			total,
			passCount, 100.0*float64(passCount)/float64(total),
			failCount, 100.0*float64(failCount)/float64(total),
		)

		if len(failList) > 0 {
			t.Logf("Примеры FAIL")
			n := min(len(failList), 100)
			for i := 0; i < n; i++ {
				r := failList[i]
				t.Logf("P1=%q, P2=%q, P3=%q, K=%q => C1=%q, C2=%q, C3=%q, C1+C2=%q, diff=%d",
					r.p1, r.p2, r.p3, r.key, r.c1, r.c2, r.c3, r.c1plus2, r.diff)
			}
		}
		if len(passList) > 0 {
			t.Logf("Примеры Pass")
			n := min(len(passList), 100)
			for i := 0; i < n; i++ {
				r := passList[i]
				t.Logf("P1=%q, P2=%q, P3=%q, K=%q => C1=%q, C2=%q, C3=%q, C1+C2=%q, diff=%d",
					r.p1, r.p2, r.p3, r.key, r.c1, r.c2, r.c3, r.c1plus2, r.diff)
			}
		}
	})

}

func Test_MergeBlock_Key(t *testing.T) {
	const (
		numInputs = 50
		numKeys   = 40
		blockLen  = 4
		keyLen    = 16
	)
	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(seed))
	alphabet := NewTrithemius(TelegraphAlphabet)

	inputs := make([]string, 0, numInputs)
	keys := make([]string, 0, numKeys)
	for i := 0; i < numInputs; i++ {
		inputs = append(inputs, randomBlock(TelegraphAlphabet, blockLen, rnd))
	}
	for i := 0; i < numKeys; i++ {
		keys = append(keys, randomBlock(TelegraphAlphabet, keyLen, rnd))
	}

	fmt.Println(inputs)
	fmt.Println(keys)

	t.Run("SmallChange", func(t *testing.T) {
		type Result struct {
			input string
			key   string
			c1    string
			c2    string
			diff  int
		}

		var noneChangeList, oneChangeList, twoChangeList, threeChangeList, fourChangeList []Result

		for _, k := range keys {
			for _, p := range inputs {
				kPrime := makeSmallChangeOnBlock(TelegraphAlphabet, k, rnd)
				c1 := alphabet.EncodeMergeBlock(p, k)
				c2 := alphabet.EncodeMergeBlock(p, kPrime)
				diff := diffPositions(c1, c2)

				res := Result{p, k, c1, c2, diff}
				if diff == 4 {
					fourChangeList = append(fourChangeList, res)
				} else if diff == 3 {
					threeChangeList = append(threeChangeList, res)
				} else if diff == 2 {
					twoChangeList = append(twoChangeList, res)
				} else if diff == 1 {
					oneChangeList = append(oneChangeList, res)
				} else {
					noneChangeList = append(noneChangeList, res)
				}
			}
		}

		total := len(inputs) * len(keys)
		fourChangeCount := len(fourChangeList)
		threeChangeCount := len(threeChangeList)
		twoChangeCount := len(twoChangeList)
		oneChangeCount := len(oneChangeList)
		noneChangeCount := len(noneChangeList)

		t.Logf("Всего тестов: %d;\n\t4 изменений=%d (%.2f%%);\n\t3 изменений=%d (%.2f%%);\n\t2 изменений=%d (%.2f%%);\n\t1 изменение=%d (%.2f%%);\n\t0 изменений=%d (%.2f%%)",
			total,
			fourChangeCount, 100.0*float64(fourChangeCount)/float64(total),
			threeChangeCount, 100.0*float64(threeChangeCount)/float64(total),
			twoChangeCount, 100.0*float64(twoChangeCount)/float64(total),
			oneChangeCount, 100.0*float64(oneChangeCount)/float64(total),
			noneChangeCount, 100.0*float64(noneChangeCount)/float64(total),
		)

		printExamples := func(title string, list []Result) {
			if len(list) == 0 {
				return
			}
			t.Logf("Список тестов. %s:", title)
			n := len(list)

			for i := 0; i < n; i++ {
				r := list[i]
				t.Logf("P=%q, K=%q, K'=%q => C1=%q, C2=%q, diff=%d",
					r.input, r.key, makeSmallChangeOnBlock(TelegraphAlphabet, r.key, rnd), r.c1, r.c2, r.diff)
			}
		}
		printExamples("0 символов поменялось", noneChangeList)
		printExamples("1 символ поменялся", oneChangeList)
		printExamples("2 символа поменялось", twoChangeList)
		printExamples("3 символа поменялось", threeChangeList)
		printExamples("4 символов поменялось", fourChangeList)
	})

	t.Run("Rotate", func(t *testing.T) {
		type RResult struct {
			input  string
			rot    int
			key    string
			c1     string
			c2     string
			sameMS bool
		}
		var passList []RResult
		var failList []RResult

		for _, k := range keys {
			for _, p := range inputs {
				shift := 1 + rnd.Intn(blockLen-1)

				kPrime := rotateRightRunes(k, shift)
				c1 := alphabet.EncodeMergeBlock(p, k)
				c2 := alphabet.EncodeMergeBlock(p, kPrime)
				same := outputsHaveSameMultiset(c1, c2)

				res := RResult{p, shift, k, c1, c2, same}
				if same {
					failList = append(failList, res)
				} else {
					passList = append(passList, res)
				}
			}
		}

		total := len(inputs) * len(keys)
		passCount := len(passList)
		failCount := len(failList)

		t.Logf("Rotate test: total=%d, pass=%d (%.2f%%), fail=%d (%.2f%%)",
			total,
			passCount, 100.0*float64(passCount)/float64(total),
			failCount, 100.0*float64(failCount)/float64(total),
		)

		t.Logf("Примеры неудач:")
		n := len(failList)
		for i := 0; i < n; i++ {
			r := failList[i]
			t.Logf("P=%q, shift=%d, K=%q, K'=%q => C1=%q, C2=%q",
				r.input, r.rot, r.key, rotateRightRunes(r.key, r.rot), r.c1, r.c2)
		}
		t.Logf("Примеры удач:")
		n = len(passList)
		for i := 0; i < n; i++ {
			r := passList[i]
			t.Logf("P=%q, shift=%d, K=%q, K'=%q => C1=%q, C2=%q",
				r.input, r.rot, r.key, rotateRightRunes(r.key, r.rot), r.c1, r.c2)
		}

	})

	t.Run("Linear", func(t *testing.T) {
		type LResult struct {
			p       string
			k1      string
			k2      string
			k3      string
			c1      string
			c2      string
			c3      string
			c1plus2 string
			diff    int
		}
		var passList, failList []LResult

		for i := 0; i < len(keys); i++ {
			for j := 0; j < len(keys); j++ {
				if i == j {
					continue
				}
				k1 := keys[i]
				k2 := keys[j]
				k3 := alphabet.Alphabet.AddTxt(k1, k2)
				if k3 == k1 || k3 == k2 {
					continue
				}
				for _, p := range inputs {
					c1 := alphabet.EncodeMergeBlock(p, k1)
					c2 := alphabet.EncodeMergeBlock(p, k2)
					c3 := alphabet.EncodeMergeBlock(p, k3)
					c1plus2 := alphabet.Alphabet.AddTxt(c1, c2)
					diff := diffPositions(c1plus2, c3)

					res := LResult{p, k1, k2, k3, c1, c2, c3, c1plus2, diff}
					if diff != 0 {
						passList = append(passList, res)
					} else {
						failList = append(failList, res)
					}
				}
			}
		}

		total := len(passList) + len(failList)
		passCount := len(passList)
		failCount := len(failList)

		t.Logf("Linear test: total=%d, pass=%d (%.2f%%), fail=%d (%.2f%%)",
			total,
			passCount, 100.0*float64(passCount)/float64(total),
			failCount, 100.0*float64(failCount)/float64(total),
		)

		if len(failList) > 0 {
			t.Logf("Примеры FAIL")
			n := min(len(failList), 100)
			for i := 0; i < n; i++ {
				r := failList[i]
				t.Logf("P=%q, K1=%q, K2=%q, K3=%q => C1=%q, C2=%q, C3=%q, C1+C2=%q, diff=%d",
					r.p, r.k1, r.k2, r.k3, r.c1, r.c2, r.c3, r.c1plus2, r.diff)
			}
		}
		if len(passList) > 0 {
			t.Logf("Примеры Pass")
			n := min(len(passList), 100)
			for i := 0; i < n; i++ {
				r := passList[i]
				t.Logf("P1=%q, P2=%q, P3=%q, K=%q => C1=%q, C2=%q, C3=%q, C1+C2=%q, diff=%d",
					r.p, r.k1, r.k2, r.k3, r.c1, r.c2, r.c3, r.c1plus2, r.diff)
			}
		}
	})
}

func Test_STrithemiusM_Input(t *testing.T) {
	const (
		numInputs = 50
		numKeys   = 40
		blockLen  = 4
		keyLen    = 16
	)
	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(seed))
	alphabet := NewTrithemius(TelegraphAlphabet)

	inputs := make([]string, 0, numInputs)
	keys := make([]string, 0, numKeys)
	for i := 0; i < numInputs; i++ {
		inputs = append(inputs, randomBlock(TelegraphAlphabet, blockLen, rnd))
	}
	for i := 0; i < numKeys; i++ {
		keys = append(keys, randomBlock(TelegraphAlphabet, keyLen, rnd))
	}

	fmt.Println(inputs)
	fmt.Println(keys)

	t.Run("SmallChange", func(t *testing.T) {
		type Result struct {
			input string
			key   string
			c1    string
			c2    string
			diff  int
		}

		var noneChangeList, oneChangeList, twoChangeList, threeChangeList, fourChangeList []Result

		for _, p := range inputs {
			for _, k := range keys {
				pPrime := makeSmallChangeOnBlock(TelegraphAlphabet, p, rnd)
				c1 := alphabet.EncodeSTrithemiusM(p, k)
				c2 := alphabet.EncodeSTrithemiusM(pPrime, k)
				diff := diffPositions(c1, c2)

				res := Result{p, k, c1, c2, diff}
				if diff == 4 {
					fourChangeList = append(fourChangeList, res)
				} else if diff == 3 {
					threeChangeList = append(threeChangeList, res)
				} else if diff == 2 {
					twoChangeList = append(twoChangeList, res)
				} else if diff == 1 {
					oneChangeList = append(oneChangeList, res)
				} else {
					noneChangeList = append(noneChangeList, res)
				}
			}
		}

		total := len(inputs) * len(keys)
		fourChangeCount := len(fourChangeList)
		threeChangeCount := len(threeChangeList)
		twoChangeCount := len(twoChangeList)
		oneChangeCount := len(oneChangeList)
		noneChangeCount := len(noneChangeList)

		t.Logf("Всего тестов: %d;\n\t4 изменений=%d (%.2f%%);\n\t3 изменений=%d (%.2f%%);\n\t2 изменений=%d (%.2f%%);\n\t1 изменение=%d (%.2f%%);\n\t0 изменений=%d (%.2f%%)",
			total,
			fourChangeCount, 100.0*float64(fourChangeCount)/float64(total),
			threeChangeCount, 100.0*float64(threeChangeCount)/float64(total),
			twoChangeCount, 100.0*float64(twoChangeCount)/float64(total),
			oneChangeCount, 100.0*float64(oneChangeCount)/float64(total),
			noneChangeCount, 100.0*float64(noneChangeCount)/float64(total),
		)

		printExamples := func(title string, list []Result) {
			if len(list) == 0 {
				return
			}
			t.Logf("Список тестов. %s:", title)
			n := len(list)

			for i := 0; i < n; i++ {
				r := list[i]
				t.Logf("P=%q, P'=%q, K=%q => C1=%q, C2=%q, diff=%d",
					r.input, makeSmallChangeOnBlock(TelegraphAlphabet, r.input, rnd), r.key, r.c1, r.c2, r.diff)
			}
		}
		printExamples("0 символов поменялось", noneChangeList)
		printExamples("1 символ поменялся", oneChangeList)
		printExamples("2 символа поменялось", twoChangeList)
		printExamples("3 символа поменялось", threeChangeList)
		printExamples("4 символов поменялось", fourChangeList)
	})

	t.Run("Rotate", func(t *testing.T) {
		type RResult struct {
			input  string
			rot    int
			key    string
			c1     string
			c2     string
			sameMS bool
		}
		var passList []RResult
		var failList []RResult

		for _, p := range inputs {
			for _, k := range keys {
				shift := 1 + rnd.Intn(blockLen-1)

				pPrime := rotateRightRunes(p, shift)
				c1 := alphabet.EncodeSTrithemiusM(p, k)
				c2 := alphabet.EncodeSTrithemiusM(pPrime, k)
				same := outputsHaveSameMultiset(c1, c2)

				res := RResult{p, shift, k, c1, c2, same}
				if same {
					failList = append(failList, res)
				} else {
					passList = append(passList, res)
				}
			}
		}

		total := len(inputs) * len(keys)
		passCount := len(passList)
		failCount := len(failList)

		t.Logf("Rotate test: total=%d, pass=%d (%.2f%%), fail=%d (%.2f%%)",
			total,
			passCount, 100.0*float64(passCount)/float64(total),
			failCount, 100.0*float64(failCount)/float64(total),
		)

		t.Logf("Примеры неудач:")
		n := len(failList)
		for i := 0; i < n; i++ {
			r := failList[i]
			t.Logf("P=%q, shift=%d, P'=%q, K=%q => C1=%q, C2=%q",
				r.input, r.rot, rotateRightRunes(r.input, r.rot), r.key, r.c1, r.c2)
		}
		t.Logf("Примеры удач:")
		n = len(passList)
		for i := 0; i < n; i++ {
			r := passList[i]
			t.Logf("P=%q, shift=%d, P'=%q, K=%q => C1=%q, C2=%q",
				r.input, r.rot, rotateRightRunes(r.input, r.rot), r.key, r.c1, r.c2)
		}

	})

	t.Run("Linear", func(t *testing.T) {
		type LResult struct {
			p1      string
			p2      string
			p3      string
			key     string
			c1      string
			c2      string
			c3      string
			c1plus2 string
			diff    int
		}
		var passList, failList []LResult

		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(inputs); j++ {
				if i == j {
					continue
				}
				p1 := inputs[i]
				p2 := inputs[j]
				p3 := alphabet.Alphabet.AddTxt(p1, p2)
				if p3 == p1 || p3 == p2 {
					continue
				}
				for _, k := range keys {
					c1 := alphabet.EncodeSTrithemiusM(p1, k)
					c2 := alphabet.EncodeSTrithemiusM(p2, k)
					c3 := alphabet.EncodeSTrithemiusM(p3, k)
					c1plus2 := alphabet.Alphabet.AddTxt(c1, c2)
					diff := diffPositions(c1plus2, c3)

					res := LResult{p1, p2, p3, k, c1, c2, c3, c1plus2, diff}
					if diff != 0 {
						passList = append(passList, res)
					} else {
						failList = append(failList, res)
					}
				}
			}
		}

		total := len(passList) + len(failList)
		passCount := len(passList)
		failCount := len(failList)

		t.Logf("Linear test: total=%d, pass=%d (%.2f%%), fail=%d (%.2f%%)",
			total,
			passCount, 100.0*float64(passCount)/float64(total),
			failCount, 100.0*float64(failCount)/float64(total),
		)

		if len(failList) > 0 {
			t.Logf("Примеры FAIL")
			n := min(len(failList), 100)
			for i := 0; i < n; i++ {
				r := failList[i]
				t.Logf("P1=%q, P2=%q, P3=%q, K=%q => C1=%q, C2=%q, C3=%q, C1+C2=%q, diff=%d",
					r.p1, r.p2, r.p3, r.key, r.c1, r.c2, r.c3, r.c1plus2, r.diff)
			}
		}
		if len(passList) > 0 {
			t.Logf("Примеры Pass")
			n := min(len(passList), 100)
			for i := 0; i < n; i++ {
				r := passList[i]
				t.Logf("P1=%q, P2=%q, P3=%q, K=%q => C1=%q, C2=%q, C3=%q, C1+C2=%q, diff=%d",
					r.p1, r.p2, r.p3, r.key, r.c1, r.c2, r.c3, r.c1plus2, r.diff)
			}
		}
	})

}

func Test_STrithemiusM_Key(t *testing.T) {
	const (
		numInputs = 50
		numKeys   = 40
		blockLen  = 4
		keyLen    = 16
	)
	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(seed))
	alphabet := NewTrithemius(TelegraphAlphabet)

	inputs := make([]string, 0, numInputs)
	keys := make([]string, 0, numKeys)
	for i := 0; i < numInputs; i++ {
		inputs = append(inputs, randomBlock(TelegraphAlphabet, blockLen, rnd))
	}
	for i := 0; i < numKeys; i++ {
		keys = append(keys, randomBlock(TelegraphAlphabet, keyLen, rnd))
	}

	fmt.Println(inputs)
	fmt.Println(keys)

	t.Run("SmallChange", func(t *testing.T) {
		type Result struct {
			input string
			key   string
			c1    string
			c2    string
			diff  int
		}

		var noneChangeList, oneChangeList, twoChangeList, threeChangeList, fourChangeList []Result

		for _, k := range keys {
			for _, p := range inputs {
				kPrime := makeSmallChangeOnBlock(TelegraphAlphabet, k, rnd)
				c1 := alphabet.EncodeSTrithemiusM(p, k)
				c2 := alphabet.EncodeSTrithemiusM(p, kPrime)
				diff := diffPositions(c1, c2)

				res := Result{p, k, c1, c2, diff}
				if diff == 4 {
					fourChangeList = append(fourChangeList, res)
				} else if diff == 3 {
					threeChangeList = append(threeChangeList, res)
				} else if diff == 2 {
					twoChangeList = append(twoChangeList, res)
				} else if diff == 1 {
					oneChangeList = append(oneChangeList, res)
				} else {
					noneChangeList = append(noneChangeList, res)
				}
			}
		}

		total := len(inputs) * len(keys)
		fourChangeCount := len(fourChangeList)
		threeChangeCount := len(threeChangeList)
		twoChangeCount := len(twoChangeList)
		oneChangeCount := len(oneChangeList)
		noneChangeCount := len(noneChangeList)

		t.Logf("Всего тестов: %d;\n\t4 изменений=%d (%.2f%%);\n\t3 изменений=%d (%.2f%%);\n\t2 изменений=%d (%.2f%%);\n\t1 изменение=%d (%.2f%%);\n\t0 изменений=%d (%.2f%%)",
			total,
			fourChangeCount, 100.0*float64(fourChangeCount)/float64(total),
			threeChangeCount, 100.0*float64(threeChangeCount)/float64(total),
			twoChangeCount, 100.0*float64(twoChangeCount)/float64(total),
			oneChangeCount, 100.0*float64(oneChangeCount)/float64(total),
			noneChangeCount, 100.0*float64(noneChangeCount)/float64(total),
		)

		printExamples := func(title string, list []Result) {
			if len(list) == 0 {
				return
			}
			t.Logf("Список тестов. %s:", title)
			n := len(list)

			for i := 0; i < n; i++ {
				r := list[i]
				t.Logf("P=%q, K=%q, K'=%q => C1=%q, C2=%q, diff=%d",
					r.input, r.key, makeSmallChangeOnBlock(TelegraphAlphabet, r.key, rnd), r.c1, r.c2, r.diff)
			}
		}
		printExamples("0 символов поменялось", noneChangeList)
		printExamples("1 символ поменялся", oneChangeList)
		printExamples("2 символа поменялось", twoChangeList)
		printExamples("3 символа поменялось", threeChangeList)
		printExamples("4 символов поменялось", fourChangeList)
	})

	t.Run("Rotate", func(t *testing.T) {
		type RResult struct {
			input  string
			rot    int
			key    string
			c1     string
			c2     string
			sameMS bool
		}
		var passList []RResult
		var failList []RResult

		for _, k := range keys {
			for _, p := range inputs {
				shift := 1 + rnd.Intn(blockLen-1)

				kPrime := rotateRightRunes(k, shift)
				c1 := alphabet.EncodeSTrithemiusM(p, k)
				c2 := alphabet.EncodeSTrithemiusM(p, kPrime)
				same := outputsHaveSameMultiset(c1, c2)

				res := RResult{p, shift, k, c1, c2, same}
				if same {
					failList = append(failList, res)
				} else {
					passList = append(passList, res)
				}
			}
		}

		total := len(inputs) * len(keys)
		passCount := len(passList)
		failCount := len(failList)

		t.Logf("Rotate test: total=%d, pass=%d (%.2f%%), fail=%d (%.2f%%)",
			total,
			passCount, 100.0*float64(passCount)/float64(total),
			failCount, 100.0*float64(failCount)/float64(total),
		)

		t.Logf("Примеры неудач:")
		n := len(failList)
		for i := 0; i < n; i++ {
			r := failList[i]
			t.Logf("P=%q, shift=%d, K=%q, K'=%q => C1=%q, C2=%q",
				r.input, r.rot, r.key, rotateRightRunes(r.key, r.rot), r.c1, r.c2)
		}
		t.Logf("Примеры удач:")
		n = len(passList)
		for i := 0; i < n; i++ {
			r := passList[i]
			t.Logf("P=%q, shift=%d, K=%q, K'=%q => C1=%q, C2=%q",
				r.input, r.rot, r.key, rotateRightRunes(r.key, r.rot), r.c1, r.c2)
		}

	})

	t.Run("Linear", func(t *testing.T) {
		type LResult struct {
			p       string
			k1      string
			k2      string
			k3      string
			c1      string
			c2      string
			c3      string
			c1plus2 string
			diff    int
		}
		var passList, failList []LResult

		for i := 0; i < len(keys); i++ {
			for j := 0; j < len(keys); j++ {
				if i == j {
					continue
				}
				k1 := keys[i]
				k2 := keys[j]
				k3 := alphabet.Alphabet.AddTxt(k1, k2)
				if k3 == k1 || k3 == k2 {
					continue
				}
				for _, p := range inputs {
					c1 := alphabet.EncodeSTrithemiusM(p, k1)
					c2 := alphabet.EncodeSTrithemiusM(p, k2)
					c3 := alphabet.EncodeSTrithemiusM(p, k3)
					c1plus2 := alphabet.Alphabet.AddTxt(c1, c2)
					diff := diffPositions(c1plus2, c3)

					res := LResult{p, k1, k2, k3, c1, c2, c3, c1plus2, diff}
					if diff != 0 {
						passList = append(passList, res)
					} else {
						failList = append(failList, res)
					}
				}
			}
		}

		total := len(passList) + len(failList)
		passCount := len(passList)
		failCount := len(failList)

		t.Logf("Linear test: total=%d, pass=%d (%.2f%%), fail=%d (%.2f%%)",
			total,
			passCount, 100.0*float64(passCount)/float64(total),
			failCount, 100.0*float64(failCount)/float64(total),
		)

		if len(failList) > 0 {
			t.Logf("Примеры FAIL")
			n := min(len(failList), 100)
			for i := 0; i < n; i++ {
				r := failList[i]
				t.Logf("P=%q, K1=%q, K2=%q, K3=%q => C1=%q, C2=%q, C3=%q, C1+C2=%q, diff=%d",
					r.p, r.k1, r.k2, r.k3, r.c1, r.c2, r.c3, r.c1plus2, r.diff)
			}
		}
		if len(passList) > 0 {
			t.Logf("Примеры Pass")
			n := min(len(passList), 100)
			for i := 0; i < n; i++ {
				r := passList[i]
				t.Logf("P1=%q, P2=%q, P3=%q, K=%q => C1=%q, C2=%q, C3=%q, C1+C2=%q, diff=%d",
					r.p, r.k1, r.k2, r.k3, r.c1, r.c2, r.c3, r.c1plus2, r.diff)
			}
		}
	})
}
