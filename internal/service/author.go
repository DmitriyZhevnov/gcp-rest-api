package service

import (
	"context"
	"fmt"

	"github.com/DmitriyZhevnov/rest-api/internal/apperror"
	"github.com/DmitriyZhevnov/rest-api/internal/model"
	"github.com/DmitriyZhevnov/rest-api/internal/repository"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"github.com/DmitriyZhevnov/rest-api/pkg/utils"
)

type authorService struct {
	repository repository.Author
	logger     *logging.Logger
}

func NewAuthorSevice(repos repository.Author, logger *logging.Logger) *authorService {
	return &authorService{
		repository: repos,
		logger:     logger,
	}
}

func (s *authorService) FindAll(ctx context.Context) ([]model.Author, error) {
	return s.repository.FindAll(ctx)
}

func (s *authorService) FindAuthor(ctx context.Context, id string) (model.Author, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *authorService) Create(ctx context.Context, dto model.CreateAuthorDTO) (string, error) {
	author := model.Author{
		Name: dto.Name,
		Age:  dto.Age,
	}

	return s.repository.Create(ctx, author)
}

func (s *authorService) Update(ctx context.Context, id string, dto model.UpdateAuthorDTO) error {
	if !utils.IsValidUUID(id) {
		return apperror.NewErrNotFound(fmt.Sprintf("failed to convert authorID to UUID. ID=%s", id), "23423424")
	}

	if dto.Name == "" && dto.Age == 0 {
		return apperror.NewBadRequestError("invalid body", "2340584930")
	}

	aurhor := model.Author{
		ID:   id,
		Name: dto.Name,
		Age:  dto.Age,
	}

	return s.repository.Update(ctx, aurhor)
}

func (s *authorService) Delete(ctx context.Context, id string) error {
	if !utils.IsValidUUID(id) {
		return apperror.NewErrNotFound(fmt.Sprintf("failed to convert authorID to UUID. ID=%s", id), "23423424")
	}

	return s.repository.Delete(ctx, id)
}
