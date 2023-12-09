package dto

type EncryptRequest struct {
	Payload   string `json:"payload"`
	SecretKey string `json:"secretKey"`
	PublicKey string `json:"publicKey"`
}

type EncryptResponse struct {
	EncryptedPayload string `json:"encryptedPayload"`
	ApiEncryptionKey string `json:"apiEncryptionKey"`
}
