package main

import (
	"time"
)

// Block represents block within the Blockchain struct
type Block struct {
	Timestamp int64
	Data      []byte
	PrevHash  []byte
	Hash      []byte
	Nonce     int
}

// NewBlock creates new Block struct
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

// NewGenesisBlock creates the first (genesis) block within a Blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
