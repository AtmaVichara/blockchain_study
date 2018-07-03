package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlock(t *testing.T) {
	block := NewBlock("some data", []byte{})
	blockData := []byte("some data")

	assert.NotNil(t, block)
	assert.Equal(t, block.Data, blockData, "They should equal")
}

func TestNewGenesisBlock(t *testing.T) {
	genBlock := NewGenesisBlock()
	genData := []byte("Genesis Block")

	assert.NotNil(t, genBlock)
	assert.Equal(t, genBlock.Data, genData, "They should equal")
}
