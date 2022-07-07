package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/DmitriyZhevnov/rest-api/internal/model"
	"github.com/DmitriyZhevnov/rest-api/pkg/client/postgresql"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
)

type Repository struct {
	User
	Author
}

type User interface {
	Create(ctx context.Context, user model.User) (string, error)
	FindOne(ctx context.Context, id string) (model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, id string) error
}

type Author interface {
	FindAll(ctx context.Context) (u []model.Author, err error)
	FindOne(ctx context.Context, id string) (model.Author, error)
	Create(ctx context.Context, aurhor model.Author) (string, error)
	Update(ctx context.Context, user model.Author) error
	Delete(ctx context.Context, id string) error
}

func NewRepository(postgresClient postgresql.Client, firestoreClient firestore.Client, logger *logging.Logger) *Repository {
	return &Repository{
		User:   NewUserFirestore(firestoreClient, logger),
		Author: NewAuthorPostgres(postgresClient, logger),
	}
}
