package usecase

import (
	"context"
	"echo-template/db"
	"echo-template/internal/infrastructure/repository"
	"echo-template/internal/models"
	"echo-template/pkg/hash"

	tokenjwt "echo-template/pkg/token_jwt"
)

type ClientService struct {
	clientRepo *repository.ClientRepository
}

func NewClientService(repo *repository.ClientRepository) *ClientService {
	return &ClientService{clientRepo: repo}
}

func (s *ClientService) SignUpClient(ctx context.Context, c *models.ClientSignUp) (*models.SignSuccess, error) {
	pswHsh, err := hash.GenerateHash(c.Password)
	if err != nil {
		return nil, err
	}
	params := db.CreateClientParams{
		PasswordHash: pswHsh,
		Email:        c.Email,
		Name:         c.Name,
	}
	client, err := s.clientRepo.CreateClient(ctx, params)
	if err != nil {
		return &models.SignSuccess{}, err
	}

	token, err := tokenjwt.GenerateJWT(client.ID.String())
	if err != nil {
		return &models.SignSuccess{}, err
	}

	return &models.SignSuccess{
		Token: token,
		ID:    client.ID.String(),
	}, nil
}

func (s *ClientService) SignInClient(ctx context.Context, c *models.ClientSignIn) (*models.SignSuccess, error) {
	client, err := s.clientRepo.GetClientByEmail(ctx, c.Email)
	if err != nil {
		return nil, err
	}
	if err := hash.ComparePassword(c.Password, client.PasswordHash); err != nil {
		return nil, err
	}

	token, err := tokenjwt.GenerateJWT(client.ID.String())
	if err != nil {
		return nil, err
	}

	// log.Println(token)
	return &models.SignSuccess{
		Token: token,
		ID:    client.ID.String(),
	}, nil
}
