package dao

import (
	"fmt"
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
	err = db.Model(&model.User{Id: userid}).Association("Follows").Find(&users)
	fmt.Println(users)
	return
}

func GetFollowers(userid int64) (users []*model.User, err error) {
	var followers []*model.User
	var ids []int64
	err = db.Table("user_follows").Where("follow_id = ?", userid).Pluck("user_id", &ids).Error
	if err != nil {
		return nil, err
	}
	db.Where("id in (?)", ids).Find(&followers)
	return followers, err
}

func RelationAction(id, toid int64, action_type string) (err error) {
	if action_type == "1" {
		err = db.Model(&model.User{Id: id}).Association("Follows").Append(&model.User{Id: toid})
	} else {
		err = db.Model(&model.User{Id: id}).Association("Follows").Delete(&model.User{Id: toid})
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
