package main

import (
	"fmt"
	"strconv"
)

type BlockChain struct {
	blocks []*Block
}

func (bc *BlockChain) AddBlock(transactions []*Transaction) {
	prevHash := bc.blocks[len(bc.blocks)-1].PrevHash
	newBlock := NewBlock(transactions, prevHash)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *BlockChain) FindAllUTXO() map[string]TxOutput {
	utxo := make(map[string]TxOutput)
	spentTxos := make(map[string]bool)

	for _, block := range bc.blocks {
		for _, tx := range block.Transactions {
			txid := tx.ID
			for i, txo := range tx.Vout {
				key := string(txid) + strconv.Itoa(i)
				if !spentTxos[key] {
					utxo[key] = txo
				}
			}

			if !IsCoinBase(tx) {
				for _, txi := range tx.Vin {
					key := string(txi.Txid) + strconv.Itoa(txi.OutIndex)
					spentTxos[key] = true
					delete(utxo, key)
				}

			}

		}
	}

	return utxo
}

func (bc *BlockChain) GetBalance(address string) int {
	var balance int

	utxos := bc.FindAllUTXO()
	for _, utxo := range utxos {
		if string(utxo.ScriptPubkey) == address {
			balance += utxo.Value
		}
	}

	return balance
}

func (bc *BlockChain) FindSpendableUTXOS(amount int, address string) (map[string][]int, int) {

	payable := make(map[string][]int)
	acc := 0

	utxos := bc.FindAllUTXO()
	for key, output := range utxos {
		if string(output.ScriptPubkey) == address {

			txid := key[:len(key)-1]
			outidx := key[len(key)-1]

			acc += output.Value
			payable[txid] = append(payable[txid], int(outidx))

			if acc >= amount {
				break
			}
		}
	}

	return payable, acc

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
