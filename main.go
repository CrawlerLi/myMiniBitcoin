package main

import "fmt"

func main() {
	// 创建链，奖励给 "MyAddress"
	bc := NewBlockChain("MyAddress")

	// 打包一笔交易（这里先用 coinbase 模拟）
	bc.AddBlock([]*Transaction{
		NewCoinBase("MyAddress", ""),
	})

	// 再挖一个块
	bc.AddBlock([]*Transaction{
		NewCoinBase("MyAddress", ""),
	})

	// 打印
	for i, block := range bc.blocks {
		fmt.Printf("===== Block %d =====\n", i)
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Hash:     %x\n", block.Hash)

		for _, tx := range block.Transactions {
			fmt.Printf("TxID: %x\n", tx.ID)
			for _, out := range tx.Vout {
				fmt.Printf("  Output: %s => %d\n", string(out.ScriptPubkey), out.Value)
			}
		}

	}
}
