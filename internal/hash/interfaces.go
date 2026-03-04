package hash

type Hasher interface {
	Hash(data []byte) string
}
