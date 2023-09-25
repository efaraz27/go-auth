package models

import (
	"time"

	"github.com/google/uuid"
)

// RoleType is a custom type for Role
type RoleType string

// Enum values for RoleType
const (
	AdminRole RoleType = "admin"
	UserRole  RoleType = "user"
)

type User struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Uuid      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();index" json:"uuid"`
	Email     string    `gorm:"type:varchar(100);not null;unique;index" json:"email"`
	Password  string    `gorm:"type:varchar(100);not null" json:"password"`
	FirstName string    `gorm:"type:varchar(100);not null" json:"firstName"`
	LastName  string    `gorm:"type:varchar(100);not null" json:"lastName"`
	Role      RoleType  `gorm:"type:varchar(100);default:'user'" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
