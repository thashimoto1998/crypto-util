package main

import (
	"fmt"
	"bytes"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/libp2p/go-libp2p-core/crypto"
)

func main() {
	priv := secp256k1.GenPrivKey()
	privByte := priv.Bytes()
	pubByte := priv.PubKey().Bytes()
	fmt.Println("private key:", priv)
	fmt.Println("private key bytes:", priv.Bytes())
	fmt.Println("pub key:", priv.PubKey())
	fmt.Println("pub key bytes:", priv.PubKey().Bytes())

	fmt.Println("Create libp2p private key instance from private key bytes of cosmos sdk")
	priv2, _ := crypto.UnmarshalSecp256k1PrivateKey(priv.Bytes())
	priv2Byte, _ := priv2.Raw()	
	pub2Byte, _ := priv2.GetPublic().Raw()
	fmt.Println("private key:", priv2)
	fmt.Println("private key bytes:", priv2Byte)
	fmt.Println("public key:", priv2.GetPublic())
	fmt.Println("pub key byte:", pub2Byte)

	fmt.Println("privByte == priv2Byte:", bytes.Compare(privByte, priv2Byte) == 0)
	fmt.Println("pubByte == pub2Byte:", bytes.Compare(pubByte, pub2Byte) == 0)
}