package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/dao"
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
	to_user_id := c.Query("to_user_id")
	action_type := c.Query("action_type")
	//userid to int64
	to_user_idInt64, _ := strconv.ParseInt(to_user_id, 10, 64)
	if id == to_user_idInt64 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Cannot follow yourself"})
		return
	}
	err := dao.RelationAction(id, to_user_idInt64, action_type)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Relation Action Error"})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Relation Action Success"})
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	//id:=utils.GetUserIdFromJwtToken(c)
	userid := c.Query("user_id")
	//userid to int64
	useridInt64, _ := strconv.ParseInt(userid, 10, 64)
	users, err := dao.GetFollows(useridInt64)

	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Get Follow List Error"})
		return
	}
	var userlist []User
	for i := 0; i < len(users); i++ {
		userlist = append(userlist, User{
			Id:            users[i].Id,
			Name:          users[i].Username,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      true,
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
	userid := c.Query("user_id")
	//userid to int64
	useridInt64, _ := strconv.ParseInt(userid, 10, 64)
	users, err := dao.GetFollowers(useridInt64)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Get Follow List Error"})
		return
	}
	var userlist []User
	for i := 0; i < len(users); i++ {
		userlist = append(userlist, User{
			Id:            users[i].Id,
			Name:          users[i].Username,
			FollowCount:   users[i].FollowCount,
			FollowerCount: users[i].FollowerCount,
			IsFollow:      dao.CheckIfFollow(useridInt64, users[i].Id),
		})
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userlist,
	})
}
