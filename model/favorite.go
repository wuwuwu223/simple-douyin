package model

import (
	"gorm.io/gorm"
	"time"
)

type UserFavorite struct {
	UserId    int64          `json:"user_id,omitempty" gorm:"index:idx_user_favorite"`
	VideoID   int64          `json:"video_id,omitempty" gorm:"index:idx_user_favorite"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index:idx_user_favorite"`
}

func (u *UserFavorite) AfterCreate(tx *gorm.DB) (err error) {
	user := User{}
	if err = tx.Model(&User{}).Where("id = ?", u.UserId).First(&user).Error; err != nil {
		return err
	}
	user.FavoriteCount++
	if err = tx.Model(&User{}).Where("id=?", u.UserId).Update("favorite_count", user.FavoriteCount).Error; err != nil {
		return err
	}

	video := Video{}
	if err = tx.Model(&Video{}).Where("id = ?", u.VideoID).First(&video).Error; err != nil {
		return err
	}
	video.FavoriteCount++
	if err = tx.Model(&Video{}).Where("id=?", u.VideoID).Update("favorite_count", video.FavoriteCount).Error; err != nil {
		return err
	}
	user2 := User{}
	if err = tx.Model(&User{}).Where("id = ?", video.UserID).First(&user2).Error; err != nil {
		return err
	}
	user2.TotalFavorited++
	if err = tx.Model(&User{}).Where("id=?", video.UserID).Update("total_favorited", user2.TotalFavorited).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserFavorite) AfterDelete(tx *gorm.DB) (err error) {
	user := User{}
	if err = tx.Model(&User{}).Where("id = ?", u.UserId).First(&user).Error; err != nil {
		return err
	}
	user.FavoriteCount--
	if err = tx.Model(&User{}).Where("id=?", u.UserId).Update("favorite_count", user.FavoriteCount).Error; err != nil {
		return err
	}

	video := Video{}
	if err = tx.Model(&Video{}).Where("id = ?", u.VideoID).First(&video).Error; err != nil {
		return err
	}
	video.FavoriteCount--
	if err = tx.Model(&Video{}).Where("id=?", u.VideoID).Update("favorite_count", video.FavoriteCount).Error; err != nil {
		return err
	}
	user2 := User{}
	if err = tx.Model(&User{}).Where("id = ?", video.UserID).First(&user2).Error; err != nil {
		return err
	}
	user2.TotalFavorited--
	if err = tx.Model(&User{}).Where("id=?", video.UserID).Update("total_favorited", user2.TotalFavorited).Error; err != nil {
		return err
	}
	return nil
}
