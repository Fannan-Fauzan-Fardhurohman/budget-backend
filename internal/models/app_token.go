package models

import "time"

type AppTokenModel struct {
	BaseModel
	Token     string    `json:"-" gorm:"type:varchar(255)"`
	TargetId  uint      `json:"target_id" gorm:"index; not null"`
	Type      string    `json:"-" gorm:"index;not null;type:varchar(255)"`
	Used      bool      `json:"-" gorm:"index;not null;type:bool"`
	ExpiresAt time.Time `json:"-" gorm:"index;not null;"`
}

func (AppTokenModel) TableName() string {
	return "app_tokens"
}
