package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"

	"golang.org/x/crypto/scrypt"
)

func genKey() []byte {
	salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}

	dk, err := scrypt.Key([]byte("mYpasSwoRd"), salt, 32768, 8, 1, 32)
	if err != nil {
		log.Println(err)
	}
	return dk
}

func MakeIV(s string) []byte {
	payload := []byte(s)
	key := genKey()

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(payload))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Println(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], payload)

	return ciphertext
}
