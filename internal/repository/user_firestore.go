package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/DmitriyZhevnov/rest-api/internal/apperror"
	"github.com/DmitriyZhevnov/rest-api/internal/model"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"google.golang.org/api/iterator"
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

func (uf *userFirestore) Create(ctx context.Context, user model.User) (string, error) {
	res, _, err := uf.client.Collection("users").Add(ctx, map[string]interface{}{
		"email":    user.Email,
		"username": user.Username,
		"password": user.PasswordHash,
	})
	if err != nil {
		return "", apperror.NewInternalServerError(fmt.Sprintf("failed to create user due to error: %v", err), "45645234")
	}

	return res.ID, nil
}

func (uf *userFirestore) FindAll(ctx context.Context) (u []model.User, err error) {
	query := uf.client.Collection("users")
	iter := query.Documents(ctx)
	var users []model.User
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return users, apperror.NewInternalServerError(fmt.Sprintf("failed to find all users due to error: %v", err), "2234234")
		}
		user := model.User{}
		user.ID = doc.Ref.ID
		err = doc.DataTo(&user)
		if err != nil {
			return u, apperror.NewInternalServerError(fmt.Sprintf("failed to read all documents from cursor. error: %v", err), "245646")
		}
		users = append(users, user)
	}

	return users, nil
}

func (uf *userFirestore) FindOne(ctx context.Context, id string) (u model.User, err error) {
	doc := uf.client.Collection("users").Doc(id)
	var user model.User
	res, err := doc.Get(ctx)
	if err != nil {
		return user, apperror.NewErrNotFound("user not exists", "2546461")
	}

	if err = res.DataTo(&user); err != nil {
		return u, apperror.NewInternalServerError(fmt.Sprintf("failed to decode user(id: %s) from DB due to error: %v", id, err), "2345676543")
	}

	user.ID = doc.ID
	return user, nil
}

func (uf *userFirestore) Update(ctx context.Context, user model.User) error {
	u := uf.client.Collection("users").Doc(user.ID)
	_, err := u.Set(ctx, map[string]interface{}{
		"email":    user.Email,
		"username": user.Username,
		"password": user.PasswordHash,
	})
	if err != nil {
		return apperror.NewInternalServerError(fmt.Sprintf("failed to execute update user query. error: %v", err), "94530904858")
	}

	return nil
}

func (uf *userFirestore) Delete(ctx context.Context, id string) error {
	user := uf.client.Collection("users").Doc(id)
	_, err := user.Delete(ctx)
	if err != nil {
		return apperror.NewInternalServerError(fmt.Sprintf("failed to delete user. error: %v", err), "9495854932")
	}

	return nil
}
