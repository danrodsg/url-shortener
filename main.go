package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"sync"
)

var (
	urlStore    = make(map[string]string)
	mu          sync.Mutex
	secretKey   = []byte("secretaeskey12345678901234567890")
	lettersRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func encrypt(originalUrl string) string {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	plainText := []byte(originalUrl)
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	iv := cipherText[:aes.BlockSize]

	if _, err := rand.Read(iv); err != nil {
		log.Fatal(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return hex.EncodeToString(cipherText)
}

func decrypt(encryptedUrl string) string {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	cipherText, err := hex.DecodeString(encryptedUrl)
	if err != nil {
		log.Fatal(err)
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)

}

func generateShortId() string {
	b := make([]rune, 6)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(lettersRune))))
		if err != nil {
			log.Fatal(err)
		}

		b[i] = lettersRune[num.Int64()]
	}

	return string(b)

}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortId := r.URL.Path[1:]

	mu.Lock()
	encryptedUrl, ok := urlStore[shortId]
	mu.Unlock()

	if !ok {
		http.Error(w, "Está URL não existe no projeto", http.StatusNotFound)
		return
	}

	decryptedUrl := decrypt(encryptedUrl)
	http.Redirect(w, r, decryptedUrl, http.StatusNotFound)

}

func shortenUrl(w http.ResponseWriter, r *http.Request) {
	originalUrl := r.URL.Query().Get("url")
	if originalUrl == "" {
		http.Error(w, "Parametro URL no quary é obrigatório", http.StatusBadRequest)
		return
	}

	encryptedUrl := encrypt(originalUrl)
	shortId := generateShortId()

	mu.Lock()
	urlStore[shortId] = encryptedUrl
	mu.Unlock()

	shortUrl := fmt.Sprintf("http://localhost:8081/%s", shortId)
	fmt.Fprintf(w, "A URL encurtada desta url original é: %s", shortUrl)
}

func main() {
	http.HandleFunc("/shorten", shortenUrl)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Projeto iniciando e rodando na porta 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
