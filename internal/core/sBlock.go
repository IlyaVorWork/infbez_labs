package core

import (
	"infbez_labs/internal/alphabet"
	"infbez_labs/internal/cipher"
)

type SBlock struct {
	Alphabet *alphabet.Alphabet
	cipher   *cipher.Trithemius
}

func NewSBlock(alphabet *alphabet.Alphabet) *SBlock {
	return &SBlock{
		Alphabet: alphabet,
		cipher:   cipher.NewTrithemius(alphabet),
	}
}

func (s *SBlock) FrwRun(inPrime, inAux string) string {
	if (len([]rune(inPrime)) != 16) || (len([]rune(inAux)) != 16) {
		panic("неподходящая длина входных данных")
	}
	return s.cipher.EncodePolyTrithemius(inAux, inPrime)
}

func (s *SBlock) InvRun(inPrime, inAux string) string {
	if (len([]rune(inPrime)) != 16) || (len([]rune(inAux)) != 16) {
		panic("неподходящая длина входных данных")
	}
	return s.cipher.DecodePolyTrithemius(inAux, inPrime)
}
