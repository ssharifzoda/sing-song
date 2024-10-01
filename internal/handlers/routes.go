package api

import (
	"github.com/gorilla/mux"
	_ "github.com/swaggo/http-swagger"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"sing-song/internal/service"
	"sing-song/pkg/logging"
)

type Handler struct {
	service *service.Service
	logger  logging.Logger
}

func NewHandler(s *service.Service, logger logging.Logger) *Handler {
	return &Handler{service: s, logger: logger}
}

func (h *Handler) InitRoutes() *mux.Router {
	main := mux.NewRouter()
	router := main.PathPrefix("/api/v1").Subrouter()
	router.Use(CORS, RecoverAllPanic)
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	libGr := router.PathPrefix("/lib").Subrouter()
	libGr.HandleFunc("/all", h.getSongs).Methods(http.MethodGet, http.MethodOptions)
	libGr.HandleFunc("/song", h.getSong).Methods(http.MethodGet, http.MethodOptions)
	libGr.HandleFunc("/remove", h.removeSong).Methods(http.MethodDelete, http.MethodOptions)
	libGr.HandleFunc("/edit", h.editInfo).Methods(http.MethodPut, http.MethodOptions)
	libGr.HandleFunc("/add", h.addNewSong).Methods(http.MethodPost, http.MethodOptions)
	return router
}
