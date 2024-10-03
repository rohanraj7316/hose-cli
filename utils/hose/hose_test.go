package hose_test

import (
	"testing"

	"github.com/rohanraj7316/hose/utils/hose"
)

func Test_Encrypt(t *testing.T) {
	payload := "this is me rohan raj"
	secretKey := "15760b7b91427cb951011634a426e3c7"
	publicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqsalSBFQaNMCeMrEGkhDfRRfHJAAGWEu+sx+DuOKeXIB21AgbNxvv0Qm3jxVUPlRbr0wCLs+tsA67oj2dx6GNFoRznT9fEKuBvXHzqiDejjP5HmgqFgVnJgXH+2++1VUtuRcU6fHtZoWddvnlDKL3RGLLDl13ObVgsrG2nlC2a+++xvdavASnaz6TbbqLbn511U+05nnkX+vuso5GGYAMhqUf0QyDAiR0BEgZy2VX4MBngKfYpvIRwNNog7DQvm4OH9524PLz0rfxlkZT0xC403kPqd9sNHHdvJ4qnjHlPQG6aQQkAR6Potk67mGWNyDvctobPppTUsF2BYpCMhPEQIDAQAB"

	_, _, err := hose.New().Encrypt(payload, secretKey, publicKey)
	if err != nil {
		t.Errorf("error in encryption: %s", err)
	}

}

func Test_Decrypt(t *testing.T) {
	payload := "this is me rohan raj"
	secretKey := "15760b7b91427cb951011634a426e3c7"
	publicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqsalSBFQaNMCeMrEGkhDfRRfHJAAGWEu+sx+DuOKeXIB21AgbNxvv0Qm3jxVUPlRbr0wCLs+tsA67oj2dx6GNFoRznT9fEKuBvXHzqiDejjP5HmgqFgVnJgXH+2++1VUtuRcU6fHtZoWddvnlDKL3RGLLDl13ObVgsrG2nlC2a+++xvdavASnaz6TbbqLbn511U+05nnkX+vuso5GGYAMhqUf0QyDAiR0BEgZy2VX4MBngKfYpvIRwNNog7DQvm4OH9524PLz0rfxlkZT0xC403kPqd9sNHHdvJ4qnjHlPQG6aQQkAR6Potk67mGWNyDvctobPppTUsF2BYpCMhPEQIDAQAB"

	_, encryptedPayload, err := hose.New().Encrypt(payload, secretKey, publicKey)
	if err != nil {
		t.Errorf("error in encryption: %s", err)
	}

	decryptedPayload, err := hose.New().Decrypt(encryptedPayload, secretKey)
	if err != nil {
		t.Errorf("error in decryption: %s", err)
	}

	if decryptedPayload != payload {
		t.Errorf("decrypted payload does not match original payload")
	}
}
