package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlockChain(t *testing.T) {
	bc := NewBlockChain()
	genData := []byte("Genesis Block")

	assert.NotNil(t, bc)
	assert.Equal(t, bc.Blocks[0].Data, genData, "they should equal")
}

func TestAddBlock(t *testing.T) {
	bc := NewBlockChain()
	bc.AddBlock("some data")
	newBlockData := []byte("some data")

	assert.NotNil(t, bc.Blocks[1])
	assert.Equal(t, bc.Blocks[1].Data, newBlockData, "they should equal")
}
