package core

import "infbez_labs/internal/alphabet"

type PBlock struct {
	alphabet alphabet.Alphabet
}

func NewPBlock(alphabet alphabet.Alphabet) *PBlock {
	return &PBlock{
		alphabet: alphabet,
	}
}

func (pb *PBlock) FrwRound(block string, roundNum int) string {
	return ""
}

func (pb *PBlock) InvRound(block string, roundNum int) string {
	return ""
}
