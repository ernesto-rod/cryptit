package main

import (
	"fmt"

	"github.com/ernesto-rod/cryptit/decrypt"
	"github.com/ernesto-rod/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("KodeKloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
