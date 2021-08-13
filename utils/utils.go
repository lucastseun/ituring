package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
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

	publicKeyStream := x509.MarshalPKCS1PublicKey(&publicKey)

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
func RsaEncrypt(plainText []byte, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	fileBuf := make([]byte, fileInfo.Size())
	file.Read(fileBuf)
	block, _ := pem.Decode(fileBuf)
	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, key, plainText)
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

// RSA解密
func RsaDecrypt(cipherText []byte, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	fileBuf := make([]byte, fileInfo.Size())
	file.Read(fileBuf)
	block, _ := pem.Decode(fileBuf)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, key, cipherText)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}
