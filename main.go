package main

import "fmt"

func main() {
	// 创建链，奖励给 "MyAddress"
	bc := NewBlockChain("MyAddress")

	banlance := bc.GetBalance("MyAddress")

	fmt.Println("My address balance is ", banlance)

	// 打包一笔交易（这里先用 coinbase 模拟）
	bc.AddBlock([]*Transaction{NewTrasaction("MyAddress", "Bob", 30, bc),
		NewCoinBase("MyAddress", ""),
	})

	banlance = bc.GetBalance("MyAddress")
	banlance1 := bc.GetBalance("Bob")
	fmt.Println("My address balance is ", banlance, "The balance of Bob is ", banlance1)

}
