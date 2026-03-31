package authEncryptionProtocol

import (
	"infbez_labs/internal/alphabet"
	"infbez_labs/internal/hash"
)

type KeyDerivation struct {
	alphabet *alphabet.Alphabet
	hasher   *hash.Hasher
}

func NewKeyDerivation(hasher *hash.Hasher, alpha *alphabet.Alphabet) *KeyDerivation {
	return &KeyDerivation{
		hasher:   hasher,
		alphabet: alpha,
	}
}

//func (k KeyDerivation) Run(input, salt string, context []string, outputSizes []int, rounds int) []string {
//	var (
//		builder, tempBuilder, resultBulder strings.Builder
//		buffer, prk, result                string
//		out                                []string
//	)
//	builder.Grow(len(input) + len(salt))
//	builder.WriteString(input)
//	builder.WriteString(salt)
//
//	buffer = builder.String()
//
//	for i := 0; i < rounds; i++ {
//		newHash := k.hasher.Hash(buffer)
//
//		tempBuilder.Reset()
//		tempBuilder.Grow(len(newHash) + len(buffer))
//		tempBuilder.WriteString(newHash)
//		tempBuilder.WriteString(buffer)
//
//		buffer = tempBuilder.String()
//	}
//	tempBuilder.Reset()
//	prk = buffer
//
//	for i := 0; i < len(outputSizes); i++ {
//		q := (outputSizes[i] - (outputSizes[i] % 64)) / 64
//		rem := i
//		resultBulder.Reset()
//		resultBulder.Grow(32)
//
//		for rem > 0 {
//			blockSize := rem % 32
//			a := k.alphabet.GetCharByKey(blockSize)
//			resultBulder.WriteString(a)
//			rem = (rem - blockSize) / 32
//		}
//		result = resultBulder.String()
//
//		if q > 0 {
//			currentHash := prk
//
//			for j := 0; j <= q; j++ {
//				var b strings.Builder
//				b.Grow(len(currentHash) + len(context[i]) + len(prk))
//
//				b.WriteString(currentHash)
//				b.WriteString(context[i])
//				b.WriteString(prk)
//
//				currentHash = k.hasher.Hash(b.String())
//
//				result = currentHash + result
//			}
//		} else {
//			var b strings.Builder
//			b.Grow(len(prk) + len(context[i]) + len(prk))
//
//			b.WriteString(prk)
//			b.WriteString(context[i])
//			b.WriteString(prk)
//
//			result = k.hasher.Hash(b.String())
//		}
//
//		out = append(out, string([]rune(result)[:outputSizes[i]]))
//	}
//	return out
//}

func (k KeyDerivation) Run(input, salt string, context []string, outputSizes []int, rounds int) []string {
	tmp := input + salt

	for i := 0; i <= rounds; i++ {
		ext := k.hasher.Hash(tmp)
		tmp = ext + tmp
	}

	prk := tmp

	var out []string

	for i := 0; i < len(outputSizes); i++ {
		q := (outputSizes[i] - (outputSizes[i] % 64)) / 64

		rem := i
		res := ""

		for rem > 0 {
			h := rem % 32
			res = res + string(k.alphabet.GetCharByKey(h))
			rem = (rem - h) / 32
		}

		if q > 0 {
			hash2 := prk

			for j := 0; j <= q; j++ {
				tmp = hash2 + context[i] + prk
				hash2 = k.hasher.Hash(tmp)
				res = hash2 + res
			}
		} else {
			tmp = prk + context[i] + prk
			res = k.hasher.Hash(tmp)
		}

		out = append(out, string([]rune(res)[:outputSizes[i]]))
	}

	return out
}
