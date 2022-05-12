package controllers

import (
	"strconv"

	"github.com/douglira/alura-golang-api-rest.git/database"
	"github.com/douglira/alura-golang-api-rest.git/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"

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

func RegisterPersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Select("Name", "History").Create(&personality)
	json.NewEncoder(w).Encode(personality)
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	personalityId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err.Error())
	}
	json.NewDecoder(r.Body).Decode(&personality)
	personality.Id = personalityId
	database.DB.Clauses(clause.Returning{}).Updates(&personality)
	json.NewEncoder(w).Encode(personality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	personalityId := mux.Vars(r)["id"]
	var personality models.Personality
	database.DB.Clauses(clause.Returning{}).Delete(&personality, personalityId)
	json.NewEncoder(w).Encode(personality)
}
