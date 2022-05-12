package controllers

import (
	"github.com/douglira/alura-golang-api-rest.git/database"
	"github.com/douglira/alura-golang-api-rest.git/models"
	"github.com/gorilla/mux"

	"encoding/json"
	"net/http"
)

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality
	database.DB.Find(&personalities)
	json.NewEncoder(w).Encode(personalities)
}

func FindPersonality(w http.ResponseWriter, r *http.Request) {
	personalityId := mux.Vars(r)["id"]
	var personality models.Personality
	result := database.DB.First(&personality, personalityId)
	if result.RowsAffected == 0 {
		json.NewEncoder(w)
		return
	}
	json.NewEncoder(w).Encode(personality)
}
