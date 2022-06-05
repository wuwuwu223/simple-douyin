package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id            int64          `json:"id,omitempty"`
	CreatedAt     time.Time      `json:"created_at,omitempty"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at,omitempty"`
	UpdatedAt     time.Time      `json:"updated_at,omitempty"`
	Username      string         `json:"username,omitempty" gorm:"index:idx_name,unique"`
	Password      string         `json:"password,omitempty"`
	Follows       []User         `json:"follows,omitempty" gorm:"many2many:user_follows;"`
	Videos        []Video        `json:"videos,omitempty"`
	FollowerCount int64          `json:"follower_count,omitempty" gorm:"-"`
	IsFollow      bool           `json:"is_follow,omitempty" gorm:"-"`
	FollowCount   int64          `json:"follow_count,omitempty" gorm:"-"`
	Avatar        string         `json:"avatar,omitempty"`
	Favorites     []Video        `json:"favorites,omitempty" gorm:"many2many:user_favorites;"`
}

type UserFollow struct {
	UserId    int64          `json:"user_id,omitempty" gorm:"index:idx_user_follow"`
	FollowId  int64          `json:"follow_id,omitempty" gorm:"index:idx_user_follow"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index:idx_user_follow"`
}

type UserFavorite struct {
	UserId    int64          `json:"user_id,omitempty" gorm:"index:idx_user_favorite"`
	VideoID   int64          `json:"video_id,omitempty" gorm:"index:idx_user_favorite"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index:idx_user_favorite"`
}
