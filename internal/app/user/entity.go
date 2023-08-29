package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey;not null;type:uuid" json:"id"`
	Name      string         `gorm:"index" json:"name"`
	Role      string         `gorm:"not null" json:"role"`
	Email     string         `gorm:"unique;not null; index" json:"email"`
	Password  string         `gorm:"not null" json:"password"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
