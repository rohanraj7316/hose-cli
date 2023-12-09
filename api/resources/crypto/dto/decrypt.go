package dto

type DecryptRequest struct {
	EncryptedPayload string `json:"encryptedPayload"`
	SecretKey        string `json:"secretKey"`
}

type DecryptResponse struct {
	Payload string `json:"payload"`
}
