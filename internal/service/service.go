package service

import (
	"context"

	"github.com/DmitriyZhevnov/rest-api/internal/model"
	"github.com/DmitriyZhevnov/rest-api/internal/repository"
	"github.com/DmitriyZhevnov/rest-api/pkg/hash"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
)

type Service struct {
	User
	Author
}

type User interface {
	Create(ctx context.Context, dto model.CreateUserDTO) (string, error)
	FindUser(ctx context.Context, id string) (model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, id string, user model.UpdateUserDTO) error
	Delete(ctx context.Context, id string) error
}

type Author interface {
	FindAll(ctx context.Context) ([]model.Author, error)
	FindAuthor(ctx context.Context, id string) (model.Author, error)
	Create(ctx context.Context, dto model.CreateAuthorDTO) (string, error)
	Update(ctx context.Context, id string, dto model.UpdateAuthorDTO) error
	Delete(ctx context.Context, id string) error
}

func NewService(hasher hash.PasswordHasher, repos *repository.Repository, logger *logging.Logger) *Service {
	return &Service{
		User:   NewUserSevice(hasher, repos.User, logger),
		Author: NewAuthorSevice(repos.Author, logger),
	}
}
