package service

import (
	"crypto/sha256"
	"encoding/hex"
)

func NewBlockHash(block *Block) string {
	blockInfo := block.Timestamp + block.PrevHash + block.Hash + block.ValidatorAddr
	return NewHash(blockInfo)
}

func NewHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}