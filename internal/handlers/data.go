package api

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"sing-song/internal/models"
	"sing-song/pkg/consts"
	"sing-song/pkg/utils"
	"strconv"
)

// @Summary getSongs
// @Tags Data
// @Description getSongs
// @ID getSongs
// @Accept json
// @Param page query string true "Введите данные"
// @Param count query string true "Введите данные"
// @Param group query string false "Введите данные"
// @Param date_from query string true "Введите данные"
// @Param date_to query string true "Введите данные"
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /lib/all [get]
func (h *Handler) getSongs(w http.ResponseWriter, r *http.Request) {
	pageStr := mux.Vars(r)["page"]
	countStr := mux.Vars(r)["count"]
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InvalidDigitalParams), 400, 0)
		return
	}
	count, err := strconv.Atoi(countStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InvalidDigitalParams), 400, 0)
		return
	}
	groupStr := r.URL.Query().Get("group")
	dateFrom := r.URL.Query().Get("date_from")
	dateTo := r.URL.Query().Get("date_to")
	songs, err := h.service.GetSongs(page, count, groupStr, dateFrom, dateTo)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
		return
	}
	utils.Response(w, songs)
}

// @Summary getSong
// @Tags Data
// @Description getSong
// @ID getSong
// @Accept json
// @Param page query string true "Введите данные"
// @Param couplet query string true "Введите данные"
// @Param id query string true "Введите данные"
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /lib/song [get]
func (h *Handler) getSong(w http.ResponseWriter, r *http.Request) {
	pageStr := mux.Vars(r)["page"]
	countStr := mux.Vars(r)["couplet"]
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InvalidDigitalParams), 400, 0)
		return
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InvalidDigitalParams), 400, 0)
		return
	}
	couplet, err := strconv.Atoi(countStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InvalidDigitalParams), 400, 0)
		return
	}
	song, err := h.service.GetSong(id, page, couplet)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
		return
	}
	utils.Response(w, song)
}

// @Summary removeSong
// @Tags Data
// @Description removeSong
// @ID removeSong
// @Accept json
// @Param id query string true "Введите данные"
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /lib/remove [delete]
func (h *Handler) removeSong(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InvalidDigitalParams), 400, 0)
		return
	}
	h.logger.Info(id)
	err = h.service.RemoveSong(id)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
		return
	}
	h.logger.Info(consts.Success)
	utils.Response(w, consts.Success)
}

// @Summary editInfo
// @Tags Data
// @Description editInfo
// @ID editInfo
// @Accept json
// @Param input body models.Songs true "Введите данные"
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /lib/edit [put]
func (h *Handler) editInfo(w http.ResponseWriter, r *http.Request) {
	var song *models.Songs
	err := json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InvalidRequestData), 400, 0)
		return
	}
	h.logger.Info(song)
	err = h.service.EditInfo(song)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
		return
	}
	h.logger.Info(consts.Success)
	utils.Response(w, consts.Success)
}

// @Summary addNewSong
// @Tags Data
// @Description addNewSong
// @ID addNewSong
// @Accept json
// @Param input body models.Songs true "Введите данные"
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /lib/add [post]
func (h *Handler) addNewSong(w http.ResponseWriter, r *http.Request) {
	var newSong *models.Songs
	err := json.NewDecoder(r.Body).Decode(&newSong)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InvalidRequestData), 400, 0)
		return
	}
	h.logger.Info(newSong)
	err = h.service.Data.AddNewSong(newSong)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
		return
	}
	h.logger.Info(consts.Success)
	utils.Response(w, consts.Success)
}
