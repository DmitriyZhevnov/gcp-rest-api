package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/DmitriyZhevnov/rest-api/internal/model"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
)

type userFirestore struct {
	client firestore.Client
	logger *logging.Logger
}

func NewUserFirestore(client firestore.Client, logger *logging.Logger) *userFirestore {
	return &userFirestore{
		client: client,
		logger: logger,
	}
}

// TODO: Implement method
func (uf *userFirestore) Create(ctx context.Context, user model.User) (string, error) {
	return "", nil
}

// TODO: Implement method
func (uf *userFirestore) FindAll(ctx context.Context) (u []model.User, err error) {
	return u, nil
}

// TODO: Implement method
func (uf *userFirestore) FindOne(ctx context.Context, id string) (u model.User, err error) {
	return u, nil
}

// TODO: Implement method
func (uf *userFirestore) Update(ctx context.Context, user model.User) error {
	return nil
}

// TODO: Implement method
func (uf *userFirestore) Delete(ctx context.Context, id string) error {
	return nil
}
