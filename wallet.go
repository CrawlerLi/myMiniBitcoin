package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	Publickey  []byte
	Address    []byte
}

func NewWallet() *Wallet {

	var wallet *Wallet
	private, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubkey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	address := GetAddress(pubkey)

	wallet = &Wallet{
		PrivateKey: private,
		Publickey:  pubkey,
		Address:    address,
	}

	return wallet
}

func GetAddress(pubkey []byte) (address []byte) {
	Hashpub := sha256.Sum256(pubkey)
	ripemd160Hasher := ripemd160.New()
	_, err := ripemd160Hasher.Write(Hashpub[:])
	if err != nil {
		log.Panic(err)
	}

	pubkeyHash := ripemd160Hasher.Sum(nil)
	versionPayload := append([]byte{0x00}, pubkeyHash...)

	FirstSHA := sha256.Sum256(versionPayload)
	SecondSHA := sha256.Sum256(FirstSHA[:])
	checksum := SecondSHA[:4]

	fullPayload := append(versionPayload, checksum...)

	address = Base58encode(fullPayload)

	return address
}

func Base58encode(fullPayload []byte) []byte {
	const Base58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	var result []byte

	x := big.NewInt(0).SetBytes(fullPayload)
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, Base58[mod.Int64()])
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	for b := range fullPayload {
		if b == 0x00 {
			result = append([]byte{Base58[0]}, result...)
		} else {
			break
		}
	}

	return result

}
