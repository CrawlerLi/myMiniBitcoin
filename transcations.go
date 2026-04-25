package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

type TxInput struct {
	Txid      []byte
	OutIndex  int
	Signature []byte
	Pubkey    []byte
}

type TxOutput struct {
	Value        int
	ScriptPubkey []byte
}

type Transaction struct {
	ID   []byte
	Vin  []TxInput
	Vout []TxOutput
}

func (tx *Transaction) Hash() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(tx)
	if err != nil {
		fmt.Println(err)
	}

	hash := sha256.Sum256(buf.Bytes())
	return hash[:]
}

func NewCoinBase(toAdress string, data string) *Transaction {
	if data == "" {
		data = "Reward to" + toAdress
	}

	txin := TxInput{
		Txid:      []byte{},
		OutIndex:  -1,
		Signature: []byte(data),
	}

	txout := TxOutput{
		Value:        50,
		ScriptPubkey: []byte(toAdress),
	}

	tx := &Transaction{
		Vin:  []TxInput{txin},
		Vout: []TxOutput{txout},
	}

	tx.ID = tx.Hash()

	return tx
}

func NewTrasaction(from string, to string, amount int, bc *BlockChain) *Transaction {
	var tx *Transaction
	payable, acc := bc.FindSpendableUTXOS(amount, from)
	if acc < amount {
		fmt.Println("balance of the address is not enough")
	}

	var Vin []TxInput

	for txID, idxs := range payable {
		for _, idx := range idxs {
			txin := TxInput{
				Txid:      []byte(txID),
				OutIndex:  idx,
				Signature: []byte(from),
			}
			Vin = append(Vin, txin)
		}
	}

	txout := TxOutput{
		Value:        amount,
		ScriptPubkey: []byte(to),
	}

	Vout := []TxOutput{txout}

	if acc > amount {
		Vout = append(Vout, TxOutput{amount, []byte(from)})
	}

	tx = &Transaction{
		Vin:  Vin,
		Vout: Vout,
	}

	tx.ID = tx.Hash()

	return tx
}

func IsCoinBase(tx *Transaction) bool {
	return len(tx.Vin[0].Txid) == 0 && tx.Vin[0].OutIndex == -1
}
