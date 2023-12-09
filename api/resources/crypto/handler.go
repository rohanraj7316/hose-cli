package cr

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rohanraj7316/hose/api/resources/crypto/dto"
)

type Handler interface {
	GetSecret(ctx *fiber.Ctx) error
	Encrypt(*fiber.Ctx) error
	Decrypt(*fiber.Ctx) error
}

type handler struct {
	model Model
}

func NewHandler(
	model Model,
) Handler {
	return &handler{
		model: model,
	}
}

func (h *handler) GetSecret(ctx *fiber.Ctx) error {
	secretKey, err := h.model.GetSecret(ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusOK).JSON(&dto.Err{Err: err.Error()})
	}

	response := &dto.GetSecretResponse{SecretKey: secretKey}

	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *handler) Encrypt(ctx *fiber.Ctx) error {
	request := &dto.EncryptRequest{}

	err := ctx.BodyParser(request)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&dto.Err{Err: err.Error()})
	}

	apiEncryptionKey, encryptedPayload, err := h.model.Encrypt(ctx.UserContext(), request.Payload, request.SecretKey, request.PublicKey)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&dto.Err{Err: err.Error()})
	}

	response := &dto.EncryptResponse{EncryptedPayload: encryptedPayload, ApiEncryptionKey: apiEncryptionKey}

	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *handler) Decrypt(ctx *fiber.Ctx) error {
	request := &dto.DecryptRequest{}

	err := ctx.BodyParser(request)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&dto.Err{Err: err.Error()})
	}

	payload, err := h.model.Decrypt(ctx.UserContext(), request.EncryptedPayload, request.SecretKey)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&dto.Err{Err: err.Error()})
	}

	response := &dto.DecryptResponse{Payload: payload}

	return ctx.Status(http.StatusOK).JSON(response)
}
