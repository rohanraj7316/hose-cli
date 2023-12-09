package cr

import (
	"context"

	"github.com/rohanraj7316/hose/utils/hose"
)

type Model interface {
	GetSecret(ctx context.Context) (string, error)
	Encrypt(ctx context.Context, payload, secretKey, publicKey string) (string, string, error)
	Decrypt(ctx context.Context, encryptedPayload, secretKey string) (string, error)
}

type model struct {
	cr hose.Crypto
}

func NewModel(
	cr hose.Crypto,
) Model {
	return &model{
		cr: cr,
	}
}

func (m *model) GetSecret(ctx context.Context) (string, error) {
	return m.cr.GenerateSecretKey()
}

func (m *model) Encrypt(ctx context.Context, payload, secretKey, publicKey string) (string, string, error) {
	return m.cr.Encrypt(payload, secretKey, publicKey)
}

func (m *model) Decrypt(ctx context.Context, encryptedPayload, secretKey string) (string, error) {
	return m.cr.Decrypt(encryptedPayload, secretKey)
}
