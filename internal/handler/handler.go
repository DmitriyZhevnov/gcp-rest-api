package handler

import (
	"net/http"

	"github.com/DmitriyZhevnov/rest-api/internal/apperror"
	"github.com/DmitriyZhevnov/rest-api/internal/service"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	Register(router *httprouter.Router)
}

type handler struct {
	services *service.Service
	logger   *logging.Logger
}

func NewHandler(services *service.Service, logger *logging.Logger) Handler {
	return &handler{
		services: services,
		logger:   logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	// user routes
	router.HandlerFunc(http.MethodGet, usersURL, apperror.MiddleWare(h.GetAllUsers))
	router.HandlerFunc(http.MethodGet, userURL, apperror.MiddleWare(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.MiddleWare(h.CreateUser))
	router.HandlerFunc(http.MethodPut, userURL, apperror.MiddleWare(h.UpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.MiddleWare(h.DeleteUser))

	// authors routes
	router.HandlerFunc(http.MethodGet, authorsURL, apperror.MiddleWare(h.GetAllAuthors))
	router.HandlerFunc(http.MethodGet, authorURL, apperror.MiddleWare(h.GetAuthorByUUID))
	router.HandlerFunc(http.MethodPost, authorsURL, apperror.MiddleWare(h.CreateAuthor))
	router.HandlerFunc(http.MethodPut, authorURL, apperror.MiddleWare(h.UpdateAuthor))
	router.HandlerFunc(http.MethodDelete, authorURL, apperror.MiddleWare(h.DeleteAuthor))
}
