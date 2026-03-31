package authEncryptionProtocol

import (
	"errors"
	"infbez_labs/internal/alphabet"
	"infbez_labs/internal/codeRandomGenerator"
	"infbez_labs/internal/compositeCipher"
	"infbez_labs/internal/core"
)

const (
	blockSize = 16
	rounds    = 8
)

type CFB struct {
	alphabet *alphabet.Alphabet
	sBlock   *core.SBlockSTM
	spNet    *compositeCipher.SPNet
}

func NewCFB(alpha *alphabet.Alphabet) *CFB {
	sBlock := core.NewSBlockSTM(alpha)
	pBlock := core.NewPBlock(alpha)
	lfsr := codeRandomGenerator.NewLFSR(alpha)

	return &CFB{
		alphabet: alpha,
		sBlock:   sBlock,
		spNet:    compositeCipher.NewSPNet(alpha, sBlock, pBlock, lfsr),
	}
}

// ForwardCFB Forward macIn: 0 - сообщение без имитоставки;
// 1 - сообщение с имитоставкой;
// -1 - только имитоставка.
// seed - Передавать сид, ключи сами создадутся
func (c *CFB) Forward(message, iv []rune, spNetSeed string, macIn int8) string {
	var (
		ground    = []rune("________________") // 16
		out       []rune
		keyStream []rune
	)
	blockNum := len(message) / blockSize

	for i := 0; i < blockNum; i++ {
		inp := message[16*i : (i+1)*16]
		ground = c.alphabet.BlockXOR(inp, ground)
		keyStream = []rune(c.spNet.FrwSPNet(string(iv), spNetSeed, rounds))
		iv = c.alphabet.BlockXOR(inp, keyStream)
		out = append(out, iv...)
	}

	keyStream = []rune(c.spNet.FrwSPNet(string(iv), spNetSeed, rounds))
	mac := c.alphabet.BlockXOR(ground, keyStream)
	if macIn == 1 {
		out = append(out, mac...)
	} else if macIn == -1 {
		out = mac
	}
	return string(out)
}

func (c *CFB) Invert(message, iv []rune, spNetSeed string, macIn int8) string {
	var (
		ground     = []rune("________________") // 16
		out        []rune
		keyStream  []rune
		msgLen     = len(message) / blockSize
		dataBlocks = msgLen
	)
	if macIn != 0 {
		if msgLen == 0 {
			return errors.New("ошибка входных данных: сообщение должно быть не короче 16 символов").Error()
		}
		dataBlocks = msgLen - 1
	}

	for i := 0; i < dataBlocks; i++ {
		cipherBlock := message[blockSize*i : (i+1)*blockSize]
		keyStream = []rune(c.spNet.FrwSPNet(string(iv), spNetSeed, rounds))
		iv = cipherBlock
		plaintBlock := c.alphabet.BlockXOR(cipherBlock, keyStream)
		ground = c.alphabet.BlockXOR(ground, plaintBlock)
		out = append(out, plaintBlock...)

	}

	if macIn != 0 {
		recvMAC := message[(msgLen-1)*blockSize:]
		keyStream = []rune(c.spNet.FrwSPNet(string(iv), spNetSeed, rounds))
		keyStream = c.alphabet.BlockXOR(recvMAC, keyStream)
		calcMAC := c.alphabet.BlockXOR(ground, keyStream)

		if macIn == 1 {
			out = append(out, calcMAC...)
		} else {
			out = calcMAC
		}
	}
	return string(out)
}
