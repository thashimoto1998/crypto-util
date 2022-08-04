package main

import (
	"fmt"
	"encoding/hex"
	ecies "github.com/ecies/go/v2"
)

func main() {
	privKeyBytes, _ := hex.DecodeString("282113c27fe8d46ad636e1024d310d38cb4dfc019136d33a7642e963625f9d75")
	privKey := ecies.NewPrivateKeyFromBytes(privKeyBytes)
	
	plain := "805379df72dcb2ce3dd4de4512a5e178dfd5c8606ea1c1368ed0cb4d7eb7a4c927ccc9b83964d784d42ac56eb7feac409652bd6dee9cb5d2ef783357a91eccdc72a65f84a197d440cddea6e21618103488f"
	fmt.Println("plain data:", plain)

	encryptedBytes, _ := hex.DecodeString("04985cfa2222666b673c84dccb5fdbbb28c5b178d44b43e989d6668f0bc8e43585a9ad381e929af85a5cc4dab0ac0977994aad816c45ca06fe6c37999577a69fa77819d1b2d124d99a80b778c50835ae0cfc5e4a5992a7fdd525daf4fc82f24aafa063fec0a63b85c069f1cf327f57e3a97250ea288bf656216ece622e034166e8d8ac101113fe2f70730fb1919742c5834e3ac3433d8ec073a095d5c3e85ad8b6963c60afa81ba0fc2c5faaa1acaed1db69f80c118d2dbb2298e2bc7b659e9fb5643400913c6ae04f36b4c4619f02613a7f3a6eb727b1a277c1ec359aab7aa62be4ed4616f6fa4ef50553b54feca77577d9f1398be5cdf07d96b8a06ae07570890179b5")
	decryptedtext, err := ecies.Decrypt(privKey, encryptedBytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ciphertext decrypted: %s\n", string(decryptedtext))
	
	fmt.Println("decryption is success:", string(decryptedtext) == plain)
}