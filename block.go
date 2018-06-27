package main

import (
  "crypto/sha256"
  "strconv"
  "bytes"
  "time"
)

type Block struct {
  Timestamp int64
  Data []byte
  PrevHash []byte
  Hash []byte
}


func (b *Block) setHash() {
  timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
  headers := bytes.Join([][]byte{b.PrevHash, b.Data, timestamp}, []byte{})
  hash := sha256.Sum256(headers)
  b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
  block := &Block{time.Now().Unix, []byte(data), prevBlockHash, []byte{}}
  block.setHash()
  return block
}
