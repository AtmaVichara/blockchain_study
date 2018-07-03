package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Sending Sending, Testing, Testing")
	bc.AddBlock("Testing Again!!!")

	for _, block := range bc.Blocks {
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}

}
