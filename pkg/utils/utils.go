package utils

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"sing-song/internal/models"
	"strings"
)

type response struct {
	Message interface{} `json:"message"`
}

func InitConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func ErrorResponse(w http.ResponseWriter, err error, statusCode, errorCode int) {
	message := models.ErrorResponse{Error: err.Error(), ErrorCode: errorCode}
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(data)
	return
}

func Response(w http.ResponseWriter, data interface{}) {
	result := &response{Message: data}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LyricHelper(lyric string, page, couplets int) string {
	var end int
	text := strings.Split(lyric, "\n")
	start := (page * couplets) - couplets
	if start+couplets < len(text) {
		end = start + couplets
	} else {
		end = len(text) - (start + couplets)
	}
	lyric = ""
	for _, value := range text[start:end] {
		lyric += value + "\n"
	}
	return lyric
}
