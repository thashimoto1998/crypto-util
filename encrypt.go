package main

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	ecies "github.com/ecies/go/v2"
)

func main() {
	priv := secp256k1.GenPrivKey()
	privByte := priv.Bytes()
	pubByte := priv.PubKey().Bytes()
	fmt.Println("private key:", priv)
	fmt.Println("private key bytes:", privByte)
	fmt.Println("pub key bytes:", pubByte)

	privKey := ecies.NewPrivateKeyFromBytes(privByte)
	fmt.Println("priv key:", privKey)

	pubKey, _ := ecies.NewPublicKeyFromBytes(pubByte)
	fmt.Println("pub key:", pubKey)

	ciphertext, err := ecies.Encrypt(pubKey, []byte("THIS IS THE TEST"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("plaintext encrypted: %v\n", ciphertext)

	plaintext, err := ecies.Decrypt(privKey, ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ciphertext decrypted: %s\n", string(plaintext))
}