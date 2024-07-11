package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	PhoneNumber string         `json:"phone_number"`
	Name        string         `json:"name"`
	RoleID      uint           `json:"role_id"`
	Role        Role           `json:"role"`
	Token       *string        `json:"token"`
	ExpiredAt   *string        `json:"expired_at"`
}

func (*User) Seed(db *gorm.DB) {
	var adminRole Role
	db.Where("slug = ?", "admin").First(&adminRole)

	user := User{
		PhoneNumber: "6289603288538",
		Name:        "Administrator",
		RoleID:      adminRole.ID,
	}

	db.FirstOrCreate(&User{}, user)
}
