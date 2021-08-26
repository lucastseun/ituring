package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"os"
)

func GenRsaKey(bits int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		return err
	}

	privateKeyStream := x509.MarshalPKCS1PrivateKey(privateKey)

	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyStream,
	}

	privateKeyFile, err := os.Create("./utils/rsa_private_key.pem")

	if err != nil {
		return err
	}

	defer privateKeyFile.Close()

	err = pem.Encode(privateKeyFile, &block)

	if err != nil {
		return err
	}

	publicKey := privateKey.PublicKey

	publicKeyStream, err := x509.MarshalPKIXPublicKey(&publicKey)

	if err != nil {
		return err
	}

	block = pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyStream,
	}

	publicKeyFile, err := os.Create("./utils/rsa_public_key.pem")

	if err != nil {
		return err
	}

	defer publicKeyFile.Close()

	err = pem.Encode(publicKeyFile, &block)

	if err != nil {
		return err
	}
	return nil
}

// RSA加密
func RsaEncrypt(plainText string) string {
	file, err := os.Open("./utils/rsa_public_key.pem")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileInfo, err := os.Stat("./utils/rsa_public_key.pem")

	if err != nil {
		panic(err)
	}

	fileBuf := make([]byte, fileInfo.Size())
	file.Read(fileBuf)
	block, _ := pem.Decode(fileBuf)
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		panic(err)
	}

	key := publicKeyInterface.(*rsa.PublicKey)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, key, []byte(plainText))

	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(cipherText)
}

// RSA解密
func RsaDecrypt(cipherText string) string {
	b, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}

	file, err := os.Open("./utils/rsa_private_key.pem")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileInfo, err := os.Stat("./utils/rsa_private_key.pem")

	if err != nil {
		panic(err)
	}

	fileBuf := make([]byte, fileInfo.Size())
	file.Read(fileBuf)
	block, _ := pem.Decode(fileBuf)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		panic(err)
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, key, []byte(b))
	if err != nil {
		panic(err)
	}
	return string(plainText)
}

func NanoId(size int) (string, error) {
	var (
		defaultAlphabet = []rune("_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		defaultSize     = 21
	)

	switch {
	case size == 0:
		size = defaultSize
	case size != 0:
		if size < 0 {
			return "", errors.New("negative id length")
		}
	default:
		return "", errors.New("unexpected parameter")
	}
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	id := make([]rune, size)
	for i := 0; i < size; i++ {
		id[i] = defaultAlphabet[bytes[i]&63]
	}
	return string(id[:size]), nil
}
