package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	message := "Hello VWBL"
	var messageLen int = len(message)
	data := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(messageLen) + message)
	hash := crypto.Keccak256Hash(data)
	fmt.Println("message:", message)
	fmt.Println("message hash:", hash.Hex())  

	sig, err := hexutil.Decode("0x99600a6d7df6232d119b73609d0aaffb94cb2f7774d50caa4a4ae5d0ff81ea070dafcd2e6affae8fef46f8d81177484faa828dadf1d6d6f0ddd93269a6d9b9da1c")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("address: 0x9490C613a200D6Cc6fBeA95199A3b5f77eB5caDA")
	fmt.Println("sig:", hexutil.Encode(sig))

	sig[crypto.RecoveryIDOffset] -= 27
	

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("recovered public key:", sigPublicKeyECDSA)
	

	addr := crypto.PubkeyToAddress(*sigPublicKeyECDSA)
	
	fmt.Println("recovered address:", addr.String())
	fmt.Println("recovered address is valid:", addr.String() == "0x9490C613a200D6Cc6fBeA95199A3b5f77eB5caDA")
}