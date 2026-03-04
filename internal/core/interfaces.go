package core

type BlockTransformer interface {
	Transform(block *CBlock, data []byte) *CBlock
}
