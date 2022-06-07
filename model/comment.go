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

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	video := Video{}
	if err = tx.Model(&Video{}).Where("id = ?", c.VideoID).First(&video).Error; err != nil {
		return err
	}
	video.CommentCount++
	if err = tx.Model(&Video{}).Where("id=?", c.VideoID).Update("comment_count", video.CommentCount).Error; err != nil {
		return err
	}
	return nil
}

func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	video := Video{}
	if err = tx.Model(&Video{}).Where("id = ?", c.VideoID).First(&video).Error; err != nil {
		return err
	}
	video.CommentCount++
	if err = tx.Model(&Video{}).Where("id=?", c.VideoID).Update("comment_count", video.CommentCount).Error; err != nil {
		return err
	}
	return nil
}
