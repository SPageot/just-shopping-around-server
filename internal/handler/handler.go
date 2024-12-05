package handler

import (
	"encoding/json"
	"just-shopping-around/internal/model"
	"net/http"
)



func GetNews(w *model.NewsResponse, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":"Hello World!",
	})
}