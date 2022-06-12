package model

import "time"

type Video struct {
	Id            int64      `json:"id,omitempty" gorm:"primaryKey"`
	CreatedAt     time.Time  `json:"created_at,omitempty"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
	UpdatedAt     time.Time  `json:"updated_at,omitempty"`
	PlayUrl       string     `json:"play_url,omitempty"`
	CoverUrl      string     `json:"cover_url,omitempty"`
	UserID        int64      `json:"user_id,omitempty"`
	Title         string     `json:"title,omitempty"`
	FavoriteCount int64      `json:"favorite_count,omitempty"`
	CommentCount  int64      `json:"comment_count,omitempty"`
	Comments      []Comment  `json:"comments,omitempty"`
}
