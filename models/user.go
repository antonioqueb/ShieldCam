package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName      string    `gorm:"not null"`
	LastName       string    `gorm:"not null"`
	DependencyName string    `gorm:"not null"`
	Position       string    `gorm:"not null"`
	Degreer        string    `gorm:"not null"`
	Department     string    `gorm:"not null"`
	Area           string    `gorm:"not null"`
	BirthDate      time.Time `gorm:"not null"`
	RfsiRadios     []Radio   `gorm:"foreignKey:RfsiRadio"`
}

func (u *User) CreateUser(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u *User) GetUserByID(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).First(&u).Error
}

func (u *User) UpdateUser(db *gorm.DB) error {
	return db.Save(&u).Error
}

func (u *User) DeleteUser(db *gorm.DB) error {
	return db.Delete(&u).Error
}

func GetUsers(db *gorm.DB, users *[]User) error {
	return db.Find(&users).Error
}
