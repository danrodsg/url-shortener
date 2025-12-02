package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
)

var (
	urlStore =  make (map[string]string)
	secretKey = []byte("secretaeskey12345678901234567890")
)

func encrypt(originalUrl string) string {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	plainText := []byte(originalUrl)
	cipherText := make ([]byte, aes.BlockSize+len(plainText))

	iv := cipherText[:aes.BlockSize]

	if _, err := rand.Read(iv); err != nil{
		log.Fatal(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return hex.EncodeToString(cipherText)


}

func shortenUrl(w http.ResponseWriter, r *http.Request) {
	originalUrl := r.URL.Query().Get("url")
	if originalUrl == ""{
		http.Error(w, "Parametro URL no quary é obrigatório", http.StatusBadRequest)
		return
	}

	encryptedUrl := encrypt(originalUrl)



}