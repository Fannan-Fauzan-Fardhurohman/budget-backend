package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type DeletedAt sql.NullTime

type BaseModel struct {
	ID        uint64         `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
