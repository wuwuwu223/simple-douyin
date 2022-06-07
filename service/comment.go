package service

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

func DeleteComment(userid, commentid int64) error {
	comment := model.Comment{}
	db.Model(&model.Comment{}).Where("id = ? and user_id = ?", commentid, userid).First(&comment)
	db.Model(&model.Comment{}).Delete(&comment)
	return nil
}

//GetComments
func GetComments(videoId int64) ([]model.Comment, error) {
	var comments []model.Comment
	db.Where("video_id = ?", videoId).Find(&comments)
	return comments, nil
}
