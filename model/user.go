package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id             int64          `json:"id,omitempty"`
	CreatedAt      time.Time      `json:"created_at,omitempty"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at,omitempty"`
	UpdatedAt      time.Time      `json:"updated_at,omitempty"`
	Username       string         `json:"username,omitempty" gorm:"index:idx_name,unique"`
	Password       string         `json:"password,omitempty"`
	Follows        []User         `json:"follows,omitempty" gorm:"many2many:user_follows;"`
	Videos         []Video        `json:"videos,omitempty"`
	FollowerCount  int64          `json:"follower_count,omitempty"`
	IsFollow       bool           `json:"is_follow,omitempty" gorm:"-"`
	FollowCount    int64          `json:"follow_count,omitempty"`
	Favorites      []Video        `json:"favorites,omitempty" gorm:"many2many:user_favorites;"`
	TotalFavorited int64          `json:"total_favorited,omitempty"`
	FavoriteCount  int64          `json:"favorite_count,omitempty"`
}

type UserFollow struct {
	UserId    int64          `json:"user_id,omitempty" gorm:"index:idx_user_follow"`
	FollowId  int64          `json:"follow_id,omitempty" gorm:"index:idx_user_follow"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index:idx_user_follow"`
}

func (u *UserFollow) AfterCreate(tx *gorm.DB) (err error) {
	user := User{}
	if err = tx.Model(&User{}).Where("id = ?", u.FollowId).First(&user).Error; err != nil {
		return err
	}
	user.FollowerCount++
	if err = tx.Model(&User{}).Where("id=?", u.FollowId).Update("follower_count", user.FollowerCount).Error; err != nil {
		return err
	}

	user2 := User{}
	if err = tx.Model(&User{}).Where("id = ?", u.UserId).First(&user2).Error; err != nil {
		return err
	}
	user2.FollowCount++
	if err = tx.Model(&User{}).Where("id=?", u.UserId).Update("follow_count", user2.FollowCount).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserFollow) AfterDelete(tx *gorm.DB) (err error) {
	user := User{}
	if err = tx.Where("id = ?", u.FollowId).First(&user).Error; err != nil {
		return err
	}
	user.FollowerCount--
	if err = tx.Model(&User{}).Where("id=?", u.FollowId).Update("follower_count", user.FollowerCount).Error; err != nil {
		return err
	}

	user2 := User{}
	if err = tx.Where("id = ?", u.UserId).First(&user2).Error; err != nil {
		return err
	}
	user2.FollowCount--
	if err = tx.Model(&User{}).Where("id=?", u.UserId).Update("follow_count", user2.FollowCount).Error; err != nil {
		return err
	}
	return nil
}
