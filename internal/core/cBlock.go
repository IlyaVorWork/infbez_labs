package core

import (
	alpha "infbez_labs/internal/alphabet"
)

var cBlockInnerState = [4]string{
	"________________",
	"ПРОЖЕКТОР_ЧЕПУХИ",
	"КОЛЫХАТЬ_ПАРОДИЮ",
	"КАРМАННЫЙ_АТАМАН",
}

type CBlock struct {
	state    [4]string
	SBlock   *SBlock
	alphabet *alpha.Alphabet
}

func NewCBlock(alpha alpha.Alphabet) *CBlock {
	return &CBlock{
		state:    cBlockInnerState,
		SBlock:   NewSBlock(alpha),
		alphabet: &alpha,
	}
}

//// RunCblock Обертка. Функция для быстрого вызова CBlock
//func RunCblock(alpha alpha.Alphabet, inArr []string, outSize int) string {
//	cBlock := NewCBlock(alpha)
//	return cBlock.Run(inArr, outSize)
//}

func (c *CBlock) Run(inArr []string, outSize int) string {
	var CountRow = len(inArr)
	c.resetState()

	for i := 0; i < CountRow; i++ {
		if len([]rune(inArr[i])) == 16 {
			c.state[i] = c.alphabet.AddTxt(c.state[i], inArr[i])
		} else {
			panic("ошибка входных данных: строка должна быть длиной 16 символов")
		}
	}
	c.state = c.mixInnerStateCBlock(c.state)
	TMP1 := c.SBlock.Run(c.state[0], c.state[2])
	TMP2 := c.SBlock.Run(c.state[3], c.state[1])
	TMP3 := c.Confuse(TMP1, TMP2)

	result := c.SBlock.Run(TMP3, TMP1)
	result = c.Compress(result, outSize)
	return result

}

func (c *CBlock) resetState() {
	c.state = cBlockInnerState
}

func (c *CBlock) Confuse(In1, In2 string) string {
	arr1 := c.alphabet.TextToArray(In1)
	arr2 := c.alphabet.TextToArray(In2)
	for i := 0; i < 16; i++ {
		arr1[i] = (max(arr1[i], arr2[i]) + i) % c.alphabet.AlphabetLength
	}
	result := c.alphabet.ArrayToText(arr1)

	return c.alphabet.AddTxt(c.alphabet.AddTxt(result, In1), In2)
}

func (c *CBlock) Compress(In16 string, outN int) string {
	if outN == 16 {
		return In16
	}
	var (
		In16Arr = []rune(In16)
		a1      = In16Arr[0:4]
		a2      = In16Arr[4:8]
		a3      = In16Arr[8:12]
		a4      = In16Arr[12:16]
	)
	if outN == 8 {
		a13 := append([]rune(nil), a1...)
		a13 = append(a13, a3...)

		a24 := append([]rune(nil), a2...)
		a24 = append(a24, a4...)

		return c.alphabet.AddTxt(string(a13), string(a24))
	} else if outN == 4 {
		var (
			a13 = c.alphabet.SubTxt(string(a1), string(a3))
			a24 = c.alphabet.SubTxt(string(a2), string(a4))
		)
		return c.alphabet.AddTxt(a13, a24)
	} else {
		panic("недопустимый размер выходных данных")
	}
}

func (c *CBlock) mixInnerStateCBlock(In [4]string) [4]string {
	var newInnerState = [4]string(make([]string, 4))

	newInnerState[0] = c.alphabet.AddTxt(In[0], In[1])
	newInnerState[1] = c.alphabet.SubTxt(In[0], In[1])
	newInnerState[2] = c.alphabet.AddTxt(newInnerState[1], c.alphabet.AddTxt(In[2], In[3]))
	newInnerState[3] = c.alphabet.AddTxt(newInnerState[0], c.alphabet.SubTxt(In[2], In[3]))

	return newInnerState
}
