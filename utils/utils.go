package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"ituring/models"
	"os"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"golang.org/x/crypto/bcrypt"
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

// 生成哈希值
func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// 校验密码是否匹配
func ValidatePassword(password string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}

// 生成nanoid
func GenerateNanoId(l ...int) (string, error) {
	var (
		defaultAlphabet = []rune("_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		defaultSize     = 21
		size            int
	)

	switch {
	case len(l) == 0:
		size = defaultSize
	case len(l) == 1:
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

const (
	accessTokenMaxAge  = 10 * time.Minute
	refreshTokenMaxAge = time.Hour
)

var (
	privateKey, publicKey = jwt.MustLoadRSA("./utils/rsa_private_key.pem", "./utils/rsa_public_key.pem")

	signer   = jwt.NewSigner(jwt.RS256, privateKey, accessTokenMaxAge)
	verifier = jwt.NewVerifier(jwt.RS256, publicKey)
)

// 生成token
func GenerateToken(claims models.UserClaims) string {
	token, err := signer.Sign(claims)
	if err != nil {
		return err.Error()
	}
	return string(token)
}

func VerifyMiddleware() iris.Handler {
	return verifier.Verify(func() interface{} {
		return new(models.UserClaims)
	})
}
