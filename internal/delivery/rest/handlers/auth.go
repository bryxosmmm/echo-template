package handlers

import (
	"echo-template/internal/infrastructure/repository"
	"echo-template/internal/models"
	"echo-template/internal/utils"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"

	usecase "echo-template/internal/use_case"
)

type AuthHandler struct {
	clientService *usecase.ClientService
	validate      *validator.Validate
}

func NewAuthHandler(clientService *usecase.ClientService) *AuthHandler {
	return &AuthHandler{
		clientService: clientService,
		validate:      validator.New(),
	}
}

func (h *AuthHandler) SignUpClient(c echo.Context) error {
	var client models.ClientSignUp
	if err := c.Bind(&client); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}

	if err := h.validate.Struct(client); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}

	sign, err := h.clientService.SignUpClient(c.Request().Context(), &client)
	if errors.Is(err, &pgconn.PgError{Code: "23505"}) {
		return c.JSON(http.StatusConflict, utils.Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, sign)
}

func (h *AuthHandler) SignInClient(c echo.Context) error {
	var client models.ClientSignIn
	if err := c.Bind(&client); err != nil {
		return c.JSON(utils.BadRequestError())
	}
	if err := h.validate.Struct(client); err != nil {
		return c.JSON(utils.BadRequestError())
	}
	sign, err := h.clientService.SignInClient(c.Request().Context(), &client)
	if errors.Is(err, repository.ErrClientNotFound) {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, sign)
}
