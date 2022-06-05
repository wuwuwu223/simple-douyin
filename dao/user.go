package dao

import (
	"simple-demo/model"
)

func GetUserByUsername(username string) (user *model.User, err error) {
	user = &model.User{}
	err = db.Where("username = ?", username).First(&user).Error
	return
}

func GetUserByID(id int64) (user *model.User, err error) {
	user = &model.User{}
	err = db.Where("id = ?", id).First(&user).Error
	db.Model(&model.UserFollow{}).Where("user_id=?", id).Count(&user.FollowCount)
	db.Model(&model.UserFollow{}).Where("follow_id=?", id).Count(&user.FollowerCount)
	return
}

func CreateUser(user model.User) (err error) {
	err = db.Create(&user).Error
	return
}

func GetFollows(userid int64) (users []*model.User, err error) {
	var ids []int64
	var userFollows []model.UserFollow
	err = db.Model(&model.UserFollow{}).Where("user_id = ?", userid).Find(&userFollows).Error
	if err != nil {
		return nil, err
	}
	for _, userFollow := range userFollows {
		ids = append(ids, userFollow.FollowId)
	}
	db.Where("id in (?)", ids).Find(&users)
	return
}

func GetFollowers(userid int64) (users []*model.User, err error) {
	var followers []*model.User
	var ids []int64
	var userFollows []model.UserFollow
	err = db.Model(&model.UserFollow{}).Where("follow_id = ?", userid).Find(&userFollows).Error
	if err != nil {
		return nil, err
	}
	for _, userFollow := range userFollows {
		ids = append(ids, userFollow.UserId)
	}
	db.Where("id in (?)", ids).Find(&followers)
	return followers, err
}

func RelationAction(id, toid int64, action_type string) (err error) {
	if action_type == "1" {
		err = db.Model(&model.UserFollow{}).Create(&model.UserFollow{UserId: id, FollowId: toid}).Error
	} else {
		err = db.Model(&model.UserFollow{}).Where("user_id = ? and follow_id = ?", id, toid).Delete(&model.UserFollow{}).Error
	}
	return nil
}

func CheckIfFollow(id, toid int64) bool {
	var userFollow model.UserFollow
	err := db.Where("user_id = ? and follow_id = ?", id, toid).First(&userFollow).Error
	if err != nil {
		return false
	}
	return true
}
