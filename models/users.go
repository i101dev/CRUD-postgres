package models

import "gorm.io/gorm"

type Users struct {
	ID       uint    `gorm:"primary key;autoIncrement" json:"id"`
	Name     *string `json:"name"`
	Age      *uint   `json:"age"`
	Location *string `json:"location"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err
}
