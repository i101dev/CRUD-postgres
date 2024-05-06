package controllers

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type User struct {
	ID       uint    `gorm:"primary key;autoIncrement" json:"id"`
	Name     *string `json:"name"`
	Age      *uint   `json:"age"`
	Location *string `json:"location"`
}

func InitUsers(db *gorm.DB) {

	if err := db.AutoMigrate(&User{}); err != nil {
		fmt.Println("Error initializing [models/users.go]")
		log.Fatal(err)
	}

	fmt.Println("\n** >>> Successfully initialized [models/users.go]")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating new user")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all users")
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting user by id")
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating user")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting user")
}
