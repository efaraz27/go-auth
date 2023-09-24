package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int       `gorm:"primaryKey" json:"id"`
	Uuid      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();index" json:"uuid"`
	FirstName string    `gorm:"type:varchar(100);not null" json:"firstName"`
	LastName  string    `gorm:"type:varchar(100);not null" json:"lastName"`
	Email     string    `gorm:"type:varchar(100);not null;unique;index" json:"email"`
	Password  string    `gorm:"type:varchar(100);not null" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
