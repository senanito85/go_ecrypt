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
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//file ecrypter
func encryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}

//filr decrypter
func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}

func main() {

	//this will read the lines of the file
	lines, err := readLines("message.in.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	//convert it to byteSlice
	stringByte := strings.Join(lines, " ")
	byteArray := []byte(stringByte)

	//Get users passphrase
	iinput := consolline()
	hashedpassfraze := createHash(iinput)

	println("Ths is hash of your passphrase: ", hashedpassfraze)
	//println("This is your data: ", string([]byte(stringByte)))

	//Encryption in console Part
	//encrresult := encrypt(byteArray, hashedpassfraze)
	//println("This is encrypted text: ", string([]byte(encrresult)))

	//write the encrypted text to a new output filename
	outputfile := "message.ecr.txt"
	encryptFile(outputfile, byteArray, hashedpassfraze)

	//Decryption in console Part
	//decryptedtext := decrypt(encrresult, hashedpassfraze)
	//println("This is Dencrypted text: ", string([]byte(decryptedtext)))

	//read the encrypted file, decrypt and print
	fmt.Println("This is the decrypted data:")
	fmt.Println(string(decryptFile(outputfile, hashedpassfraze)))

}

//function to read user's passphrase input
func consolline() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the passPhrase: ")
	text, _ := reader.ReadString('\n')
	return text
}

//Function to create hash
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

//function to encrypt
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

//Function to decrypt
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

//Function to read the text file
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
