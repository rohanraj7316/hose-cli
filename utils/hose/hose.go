package hose

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"strings"
)

const (
	KEY_LENGTH     = 32
	NONCE_LENGTH   = 16
	DATA_SEPARATOR = "."
)

type Crypto interface {
	GenerateSecretKey() (string, error)
	Encrypt(payload, secretKey, publicKey string) (string, string, error)
	Decrypt(encryptedPayload, secretKey string) (string, error)
}

type crypto struct{}

func New() Crypto {
	return &crypto{}
}

func derToPemString(key, keyType string) string {
	return fmt.Sprintf("-----BEGIN %s-----\n%s\n-----END %s-----", keyType, key, keyType)
}

func loadPublicKey(publicKey string) (*rsa.PublicKey, error) {
	var parsedKey interface{}
	dPem, _ := pem.Decode([]byte(publicKey))

	if dPem.Type == "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("invalid key type: %s", dPem.Type)
	}

	parsedKey, err := x509.ParsePKIXPublicKey(dPem.Bytes)
	if err != nil {
		return nil, err
	}

	if pubKey, ok := parsedKey.(*rsa.PublicKey); ok {
		return pubKey, nil
	}

	return nil, fmt.Errorf("failed to parse rsa public key")
}

func encryptRsa(secretKey, publicKeyStr string) (string, error) {
	publicKey, err := loadPublicKey(derToPemString(publicKeyStr, "PUBLIC KEY"))
	if err != nil {
		return "", err
	}

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey,
		[]byte(secretKey), nil)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func encryptAes(payload, secretKey string) (string, error) {
	if len(secretKey) < KEY_LENGTH {
		return "", fmt.Errorf("key len should be %d byte", KEY_LENGTH)
	}

	if payload == "" {
		return "", fmt.Errorf("empty payload")
	}

	payloadBytes := []byte(payload)

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCMWithNonceSize(block, NONCE_LENGTH)
	if err != nil {
		return "", err
	}

	// Create a nonce. Nonce should be from GCM
	nonce := make([]byte, NONCE_LENGTH)
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encryptedPayloadBytes := aesGCM.Seal(nil, nonce, payloadBytes, nil)

	nonceStr := base64.StdEncoding.EncodeToString(nonce)

	authTagStart := len(encryptedPayloadBytes) - 16
	authTagStr := base64.StdEncoding.EncodeToString(encryptedPayloadBytes[authTagStart:])

	encryptedPayload := base64.StdEncoding.EncodeToString(encryptedPayloadBytes[:authTagStart])

	fEncryptedPayload := strings.Join([]string{nonceStr, authTagStr, encryptedPayload}, DATA_SEPARATOR)

	return fEncryptedPayload, nil
}

func decryptAes(encryptedPayload, secretKey string) (string, error) {
	if len(secretKey) < KEY_LENGTH {
		return "", fmt.Errorf("key len should be %d byte", KEY_LENGTH)
	}

	encryptedPayloadArr := strings.Split(encryptedPayload, DATA_SEPARATOR)

	if encryptedPayloadArr[0] == "" {
		return "", fmt.Errorf("empty nonce value")
	}

	if encryptedPayloadArr[1] == "" {
		return "", fmt.Errorf("empty auth-tag value")
	}

	if encryptedPayloadArr[2] == "" {
		return "", fmt.Errorf("empty payload value")
	}

	nonceBytes, err := base64.StdEncoding.DecodeString(encryptedPayloadArr[0])
	if err != nil {
		return "", err
	}

	authTagByte, err := base64.StdEncoding.DecodeString(encryptedPayloadArr[1])
	if err != nil {
		return "", err
	}

	encryptedPayloadBytes, err := base64.StdEncoding.DecodeString(encryptedPayloadArr[2])
	if err != nil {
		return "", err
	}

	fEncryptedPayloadBytes := append(encryptedPayloadBytes, authTagByte...)

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCMWithNonceSize(block, NONCE_LENGTH)
	if err != nil {
		return "", err
	}

	// Decrypt the data
	payloadBytes, err := aesGCM.Open(nil, nonceBytes, fEncryptedPayloadBytes, nil)
	if err != nil {
		return "", err
	}

	return string(payloadBytes), nil
}

// start of key encryption
func (c crypto) GenerateSecretKey() (string, error) {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(key), nil
}

func (c crypto) Encrypt(payload, secretKey, publicKeyStr string) (apiEncryptionKey string, encryptedPayload string, err error) {
	apiEncryptionKey, err = encryptRsa(secretKey, publicKeyStr)
	if err != nil {
		return "", "", err
	}

	encryptedPayload, err = encryptAes(payload, secretKey)
	if err != nil {
		return "", "", err
	}

	return apiEncryptionKey, encryptedPayload, nil
}

func (c crypto) Decrypt(encryptedPayload, secretKey string) (payload string, err error) {
	payload, err = decryptAes(encryptedPayload, secretKey)
	if err != nil {
		return "", err
	}

	return payload, nil
}
