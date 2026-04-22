package main

import "fmt"

type BlockChain struct {
	blocks []*Block
}

func (bc *BlockChain) AddBlock(transactions []*Transaction) {
	prevHash := bc.blocks[len(bc.blocks)-1].PrevHash
	newBlock := NewBlock(transactions, prevHash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func NewBlockChain(address string) (bs *BlockChain) {
	coinbase := NewCoinBase(address, "")
	Genesisblock := NewGenesisBlock(coinbase)
	return &BlockChain{[]*Block{Genesisblock}}

}

func (bc *BlockChain) Print() {
	for i, block := range bc.blocks {
		fmt.Printf("========= 区块 %d =========\n", i)
		fmt.Printf("上一个区块哈希: %x\n", block.PrevHash)
		fmt.Printf("当前区块哈希: %x\n", block.Hash)
		fmt.Println()
	}
}
