// archivo radio.go
package models

import (
	"gorm.io/gorm"
)

type Radio struct {
	gorm.Model
	SerialNumber string `gorm:"not null"` // Número de serie
	Brand        string `gorm:"not null"` // Marca
	ModelRadio   string `gorm:"not null"` // Modelo
	RfsiRadio    int    `gorm:"not null"` // ID de usuario asignado
}

func (r *Radio) CreateRadio(db *gorm.DB) error {
	return db.Create(&r).Error
}

func (r *Radio) GetRadioByID(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).First(&r).Error
}

func (r *Radio) UpdateRadio(db *gorm.DB) error {
	return db.Save(&r).Error
}

func (r *Radio) DeleteRadio(db *gorm.DB) error {
	return db.Delete(&r).Error
}

func GetRadios(db *gorm.DB, radios *[]Radio) error {
	return db.Find(&radios).Error
}

func GetRadiosByUserID(db *gorm.DB, rfsiRadio uint, radios *[]Radio) error {
	return db.Where("rfsi_radio = ?", rfsiRadio).Find(&radios).Error
}
