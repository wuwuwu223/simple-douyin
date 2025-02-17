package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simple-demo/service"
	"simple-demo/utils"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	id := utils.GetUserIdFromJwtToken(c)
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	//userid to int64
	toUserIdint64, _ := strconv.ParseInt(toUserId, 10, 64)
	if id == toUserIdint64 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "不能关注自己"})
		return
	}
	err := service.RelationAction(id, toUserIdint64, actionType)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "关注操作失败"})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "关注操作成功"})
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	id := utils.GetUserIdFromToken(token)
	userid := c.Query("user_id")
	//userid to int64
	useridInt64, _ := strconv.ParseInt(userid, 10, 64)
	users, err := service.GetFollows(useridInt64)

	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "获取关注列表失败"})
		return
	}
	var userlist []User
	for i := 0; i < len(users); i++ {
		userlist = append(userlist, User{
			Id:            users[i].Id,
			Name:          users[i].Username,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      service.CheckIfFollow(id, users[i].Id),
		})
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userlist,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	id := utils.GetUserIdFromToken(token)
	userid := c.Query("user_id")
	useridInt64, _ := strconv.ParseInt(userid, 10, 64)
	users, err := service.GetFollowers(useridInt64)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "获取粉丝列表失败"})
		return
	}
	var userlist []User
	for i := 0; i < len(users); i++ {
		userlist = append(userlist, User{
			Id:            users[i].Id,
			Name:          users[i].Username,
			FollowCount:   users[i].FollowCount,
			FollowerCount: users[i].FollowerCount,
			IsFollow:      service.CheckIfFollow(id, users[i].Id),
		})
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userlist,
	})
}
