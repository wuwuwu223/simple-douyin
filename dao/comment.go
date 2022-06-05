package dao

import (
	"simple-demo/model"
	"time"
)

func AddComment(userId int64, videoId int64, content string) error {
	db.Create(&model.Comment{
		UserID:     userId,
		VideoID:    videoId,
		Content:    content,
		CreateDate: time.Now().Format("01-02"),
	})
	return nil
}

//GetComments
func GetComments(videoId int64) ([]model.Comment, error) {
	var comments []model.Comment
	db.Where("video_id = ?", videoId).Find(&comments)
	return comments, nil
}

func GetCommentCount(videoId int64) int64 {
	var count int64
	db.Model(&model.Comment{}).Where("video_id = ?", videoId).Count(&count)
	return count
}
