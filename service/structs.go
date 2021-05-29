package service

import "fmt"

type PoSNetwork struct {
	Blockchain     []*Block
	BlockchainHead *Block
	Validators     []*Node
}

type Node struct {
	Stake   int
	Address string
}

type Block struct {
	Timestamp     string
	PrevHash      string
	Hash          string
	ValidatorAddr string
}

func (n PoSNetwork) PrintBlockchainInfo() {
	for _, v := range n.Validators {
		fmt.Printf("Address: %v - Stake: %v\n", v.Address, v.Stake)
	}
}
