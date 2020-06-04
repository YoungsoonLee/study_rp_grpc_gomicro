package main

import (
	"log"

	"github.com/YoungsoonLee/restful_go/ch11/encryptString/utils"
)

func main() {
	key := "A&iwSdvdJtJn1FP2!e5bz#C5j7v3@!%I"
	message := "I am A message"
	log.Println("Original mmessage: ", message)

	encryptedString := utils.EncryptString(key, message)
	log.Println("Encrypted mmessage: ", encryptedString)

	decryptedString := utils.DecryptString(key, encryptedString)
	log.Println("Decrypted mmessage: ", decryptedString)

}
