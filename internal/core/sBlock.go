package core

import (
	"infbez_labs/internal/alphabet"
	"infbez_labs/internal/cipher"
)

type SBlockPoly struct {
	Alphabet *alphabet.Alphabet
	cipher   *cipher.Trithemius
}

func NewSBlockPloy(alphabet *alphabet.Alphabet) *SBlockPoly {
	t := cipher.NewTrithemius(alphabet)
	return &SBlockPoly{
		Alphabet: alphabet,
		cipher:   t,
	}
}

func (s *SBlockPoly) FrwRun(inPrime, inAux string) string {
	if (len([]rune(inPrime)) != 16) || (len([]rune(inAux)) != 16) {
		panic("неподходящая длина входных данных")
	}
	return s.cipher.EncodePolyTrithemius(inAux, inPrime)
}

func (s *SBlockPoly) InvRun(inPrime, inAux string) string {
	if (len([]rune(inPrime)) != 16) || (len([]rune(inAux)) != 16) {
		panic("неподходящая длина входных данных")
	}
	return s.cipher.DecodePolyTrithemius(inAux, inPrime)
}

type SBlockSTM struct {
	Alphabet *alphabet.Alphabet
	cipher   *cipher.Trithemius
}

func NewSBlockSTM(alphabet *alphabet.Alphabet) *SBlockSTM {
	t := cipher.NewTrithemius(alphabet)
	return &SBlockSTM{
		Alphabet: alphabet,
		cipher:   t,
	}
}

func (s *SBlockSTM) FrwRun(inPrime, inAux string) string {
	if (len([]rune(inPrime)) != 4) || (len([]rune(inAux)) != 16) {
		panic("неподходящая длина входных данных")
	}
	return s.cipher.EncodeSTrithemiusM(inPrime, inAux)
}

func (s *SBlockSTM) InvRun(inPrime, inAux string) string {
	if (len([]rune(inPrime)) != 4) || (len([]rune(inAux)) != 16) {
		panic("неподходящая длина входных данных")
	}
	return s.cipher.DecodeSTrithemiusM(inPrime, inAux)
}
