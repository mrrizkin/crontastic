package models

import (
	"time"

	"gorm.io/gorm"
)

type Database struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      *string        `json:"name"`
	Desc      *string        `json:"desc"`
	Host      string         `json:"host"`
	Port      string         `json:"port"`
	Username  *string        `json:"username"`
	Password  *string        `json:"password"`
	Driver    string         `json:"driver"`
}
