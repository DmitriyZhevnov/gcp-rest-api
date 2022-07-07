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
	authorsURL = "/authors"
	authorURL  = "/authors/:uuid"
)

func (h *handler) GetAllAuthors(w http.ResponseWriter, r *http.Request) error {
	authors, err := h.services.Author.FindAll(r.Context())
	if err != nil {
		return err
	}

	response.SendResponse(w, 200, authors)
	return nil
}

func (h *handler) GetAuthorByUUID(w http.ResponseWriter, r *http.Request) error {
	authorID := httprouter.ParamsFromContext(r.Context()).ByName("uuid")

	author, err := h.services.FindAuthor(r.Context(), authorID)
	if err != nil {
		return err
	}

	response.SendResponse(w, 200, author)
	return nil
}

func (h *handler) CreateAuthor(w http.ResponseWriter, r *http.Request) error {
	request := model.CreateAuthorDTO{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		return apperror.NewUnprocessableEntityError(err.Error(), "34536453")
	}

	id, err := h.services.Author.Create(r.Context(), request)
	if err != nil {
		return err
	}

	response.SendResponse(w, 201, id)
	return nil
}

func (h *handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) error {
	request := model.UpdateAuthorDTO{}

	authorID := httprouter.ParamsFromContext(r.Context()).ByName("uuid")

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		return apperror.NewUnprocessableEntityError(err.Error(), "235364576")
	}

	if err = h.services.Author.Update(r.Context(), authorID, request); err != nil {
		return err
	}

	response.SendResponse(w, 204, nil)
	return nil
}

func (h *handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) error {
	authorID := httprouter.ParamsFromContext(r.Context()).ByName("uuid")

	if err := h.services.Author.Delete(r.Context(), authorID); err != nil {
		return err
	}

	response.SendResponse(w, 204, nil)
	return nil
}
