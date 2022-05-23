package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/rand"
)

func Verify(num string, addr string, signedMsg string) bool {
	msg := CreateMsg(num)
	address := ConvertAddr(addr)
	signedMessage := ConvertSignedMsg(signedMsg)

	recoveredPub, _ := crypto.Ecrecover(msg, signedMessage)
	pubKey, _ := crypto.UnmarshalPubkey(recoveredPub)
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return address == recoveredAddr
}

// create a digest from the dummy random integer
func CreateMsg(num string) []byte {
	hash := []byte(num)
	return crypto.Keccak256(hash)
}

// convert the string addr to the Address object
func ConvertAddr(addr string) common.Address {
	return common.HexToAddress(addr)
}

// convert the string signedMsg to the hex bytes
func ConvertSignedMsg(signedMsg string) []byte {
	return common.Hex2Bytes(signedMsg)
}

// generate a dummy random integer
func genRandomInt() int {
	return rand.Int()
}
