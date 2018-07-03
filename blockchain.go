package main

// Blockchain represents the chain of blocks, i.e. 'blockchain'
type Blockchain struct {
	Blocks []*Block
}

// AddBlock inserts new Block struct into Blockchain.Blocks
func (bc *Blockchain) AddBlock(data string) {
	newBlock := NewBlock(data, bc.Blocks[len(bc.Blocks)-1].Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockChain creates new Blockchain struct
func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
