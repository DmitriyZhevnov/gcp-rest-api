package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DmitriyZhevnov/rest-api/internal/apperror"
	"github.com/DmitriyZhevnov/rest-api/internal/model"
	"github.com/DmitriyZhevnov/rest-api/pkg/response"
	"github.com/julienschmidt/httprouter"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

func (h *handler) GetAllUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := h.services.User.FindAll(r.Context())
	if err != nil {
		return err
	}

	response.SendResponse(w, 200, users)
	return nil
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	userID := httprouter.ParamsFromContext(r.Context()).ByName("uuid")

	user, err := h.services.FindUser(r.Context(), userID)
	if err != nil {
		return err
	}

	response.SendResponse(w, 200, user)
	return nil
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	request := model.CreateUserDTO{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		return apperror.NewUnprocessableEntityError(err.Error(), "234346543")
	}

	id, err := h.services.User.Create(r.Context(), request)
	if err != nil {
		return err
	}

	response.SendResponse(w, 201, id)
	return nil
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	request := model.UpdateUserDTO{}

	userID := httprouter.ParamsFromContext(r.Context()).ByName("uuid")

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		return apperror.NewUnprocessableEntityError(err.Error(), "23425365")
	}

	if err = h.services.User.Update(r.Context(), userID, request); err != nil {
		return err
	}

	response.SendResponse(w, 204, nil)
	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	userID := httprouter.ParamsFromContext(r.Context()).ByName("uuid")

	if err := h.services.User.Delete(r.Context(), userID); err != nil {
		return err
	}

	response.SendResponse(w, 204, nil)
	return nil
}
