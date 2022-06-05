package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id         int64          `json:"id,omitempty"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty"`
	UpdatedAt  time.Time      `json:"updated_at,omitempty"`
	UserID     int64          `json:"user_id,omitempty"`
	VideoID    int64          `json:"video_id,omitempty"`
	Content    string         `json:"content,omitempty"`
	CreateDate string         `json:"create_date,omitempty"`
}
