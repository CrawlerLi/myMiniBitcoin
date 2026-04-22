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
