package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	stringSlice := []string{"Hello", "my", "dear", "frieds", "I've", "come", "to", "meet", "you"}
	stringByte := strings.Join(stringSlice, " ")
	byteArray := []byte(stringByte)

	iinput := consolline()
	hashedpassfraze := createHash(iinput)

	println("Ths is hash of your passphrase: ", hashedpassfraze)
	println("This is your data: ", string([]byte(stringByte)))

	encrresult := encrypt(byteArray, hashedpassfraze)
	println("This is encrypted text: ", string([]byte(encrresult)))

	decryptedtext := decrypt(encrresult, hashedpassfraze)
	println("This is Dencrypted text: ", string([]byte(decryptedtext)))

}

func consolline() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the passPhrase: ")
	text, _ := reader.ReadString('\n')
	return text
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
