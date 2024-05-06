package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/i101dev/rss-aggregator/util"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type User struct {
	ID       uint   `gorm:"primary key;autoIncrement" json:"id"`
	Name     string `json:"name"`
	Age      uint   `json:"age"`
	Location string `json:"location"`
	UUID     string `json:"uuid"`
}

var db *gorm.DB

func InitUsers(database *gorm.DB) {

	if err := database.AutoMigrate(&User{}); err != nil {
		fmt.Println("Error initializing [models/users.go]")
		log.Fatal(err)
	}

	db = database

	fmt.Println("\n** >>> Successfully initialized [models/users.go]")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Creating new user")

	newUser := User{
		UUID: uuid.New().String(),
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newUser); err != nil {
		util.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	db.Create(&newUser)

	util.RespondWithJSON(w, 200, "New user created successfully")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	allUsers := &[]User{}

	result := db.Find(&allUsers)

	if result.Error != nil {
		fmt.Printf("Error fetching users: %s\n", result.Error)
		util.RespondWithError(w, http.StatusInternalServerError, "Error fetching users")
		return
	}

	if result.RowsAffected == 0 {
		util.RespondWithJSON(w, http.StatusOK, []User{})
		return
	}

	util.RespondWithJSON(w, http.StatusOK, allUsers)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {

	userUUID := chi.URLParam(r, "id")
	if userUUID == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Fetch user -----------------------------------------------------------
	userDat := User{}
	if err := db.Where("uuid = ?", userUUID).First(&userDat).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			util.RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		util.RespondWithError(w, http.StatusInternalServerError, "Error retrieving user")
		return
	}

	util.RespondWithJSON(w, http.StatusOK, userDat)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	userUUID := chi.URLParam(r, "id")
	if userUUID == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Fetch user -----------------------------------------------------------
	userDat := User{}
	if err := db.Where("uuid = ?", userUUID).First(&userDat).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			util.RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		util.RespondWithError(w, http.StatusInternalServerError, "Error retrieving user")
		return
	}

	updatedUser := User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedUser); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if updatedUser.Name != "" {
		userDat.Name = updatedUser.Name
	}
	if updatedUser.Location != "" {
		userDat.Location = updatedUser.Location
	}

	if err := db.Save(&userDat).Error; err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Error updating user")
		return
	}

	util.RespondWithJSON(w, http.StatusOK, userDat)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	userUUID := chi.URLParam(r, "id")
	if userUUID == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Fetch user -----------------------------------------------------------
	userDat := User{}
	if err := db.Where("uuid = ?", userUUID).First(&userDat).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			util.RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		util.RespondWithError(w, http.StatusInternalServerError, "Error retrieving user")
		return
	}

	// Delete the user
	if err := db.Delete(&userDat).Error; err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Error deleting user")
		return
	}

	// Return success message
	util.RespondWithJSON(w, http.StatusOK, "user successfully deleted")
}
