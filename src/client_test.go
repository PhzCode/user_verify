package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestVerify(t *testing.T) {
	var testAddrHex = "970e8128ab834e8eac17ab8e3812f010678cf791"
	var testPrivHex = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"
	key, _ := crypto.HexToECDSA(testPrivHex)
	num := "5577006791947779410"
	hash := []byte(num)
	msg := crypto.Keccak256(hash)
	sig, _ := crypto.Sign(msg, key)
	signedMsg := common.Bytes2Hex(sig)
	t.Log(signedMsg)
	res := Verify(num, testAddrHex, signedMsg)
	if res {
		t.Log(res)
	} else {
		t.Error(res)
	}
}
