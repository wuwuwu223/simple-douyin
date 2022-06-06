package dao

import (
	"simple-demo/model"
	"time"
)

//addvideotodb
func AddVideo(video *model.Video) error {
	return db.Create(video).Error
}

//getvideolistbyuserid
func GetVideoListByUserID(userid int64) (videos []*model.Video, err error) {
	err = db.Where("user_id = ?", userid).Find(&videos).Error
	return
}

//getvideolist
func GetVideoList() (videos []*model.Video, err error) {
	err = db.Find(&videos).Error
	return
}

//getvideolist
func GetVideoListAfterTime(t time.Time) (videos []*model.Video, err error) {
	err = db.Where("created_at < ?", t).Find(&videos).Error
	return
}

func FavoriteAction(userid, videoid int64, action_type string) error {
	if action_type == "1" {
		return db.Model(&model.UserFavorite{}).Create(&model.UserFavorite{UserId: userid, VideoID: videoid}).Error
	} else {
		return db.Model(&model.UserFavorite{}).Where("user_id=? and video_id=?", userid, videoid).Delete(&model.UserFavorite{}).Error
	}
}

func GetFavoriteVideoList(userid int64) (videos []*model.Video, err error) {
	var videoids []int64
	err = db.Model(&model.UserFavorite{}).Where("user_id = ?", userid).Pluck("video_id", &videoids).Error
	if err != nil {
		return nil, err
	}
	err = db.Where("id in (?)", videoids).Find(&videos).Error

	return
}

//check if user has favorite video
func CheckIfFavorite(userid, videoid int64) bool {
	tx := db.Model(&model.UserFavorite{}).Where("user_id = ? and video_id = ?", userid, videoid).First(&model.UserFavorite{})
	if tx.RowsAffected == 0 {
		return false
	}
	return true
}

func GetFavoriteCount(videoid int64) (count int64) {
	db.Model(&model.UserFavorite{}).Where("video_id = ?", videoid).Count(&count)
	return
}
