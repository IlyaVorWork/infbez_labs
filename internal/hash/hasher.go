package hash

import (
	"infbez_labs/internal/alphabet"
	"infbez_labs/internal/core"
	"strings"
)

type Hasher struct {
	alphabet alphabet.Alphabet
	cBlock   core.CBlock
}

func NewHasher(alphabet alphabet.Alphabet, cBlock core.CBlock) *Hasher {
	return &Hasher{alphabet, cBlock}
}

func (h *Hasher) Hash(message string) string {
	messageLen := len([]rune(message))
	sponge := NewSponge(SpongeStarterState, h.alphabet, h.cBlock)
	var builder strings.Builder

	K := 4 - (messageLen % 4)
	if K < 4 {
		message = concatStrings([]string{message, strings.Repeat("_", K)}, messageLen+K)
	}

	messageArr := []rune(message)
	messageLen = len(messageArr)

	for i := 0; i < messageLen; i += 4 {
		tpm := messageArr[i : i+4]
		sponge.SpongeAbsorb(string(tpm))
	}

	builder.Grow(64)
	for i := 0; i < 16; i++ {
		tpm := sponge.SpongeSqueeze()
		builder.WriteString(tpm)
	}

	return builder.String()
}

//func SpongeHash(message string, alphabet alphabet.Alphabet, cBlock core.CBlock) string {
//	messageLen := len([]rune(message))
//	sponge := NewSponge(SpongeInnerState, alphabet, cBlock)
//	var builder strings.Builder
//
//	K := 4 - (messageLen % 4)
//	if K < 4 {
//		message = concatStrings([]string{message, strings.Repeat("_", K)}, messageLen+K)
//	}
//
//	messageArr := []rune(message)
//	messageLen = len(messageArr)
//
//	for i := 0; i < messageLen; i += 4 {
//		tpm := messageArr[i : i+4]
//		sponge.SpongeAbsorb(string(tpm))
//	}
//
//	builder.Grow(64)
//	for i := 0; i < 16; i++ {
//		tpm := sponge.SpongeSqueeze()
//		builder.WriteString(tpm)
//	}
//
//	return builder.String()
//}
