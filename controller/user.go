package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/dao"
	"simple-demo/model"
	"simple-demo/utils"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	ok := utils.ValidateEmail(username)
	if !ok {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "邮箱格式错误"},
		})
		return
	}

	user := &model.User{Username: username, Password: password}
	err := dao.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "创建用户失败"},
		})
		return
	}
	//time.Sleep(time.Second)
	token, _ := utils.GenerateJwtToken(user.Id)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   user.Id,
		Token:    token,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//get user form db
	user, err := dao.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用户不存在"},
		})
		return
	}

	//check password
	if user.Password != password {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "密码错误"},
		})
		return
	}

	//generate token
	token, _ := utils.GenerateJwtToken(user.Id)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   user.Id,
		Token:    token,
	})
}

func UserInfo(c *gin.Context) {
	id, _ := c.Get("userId")
	user, err := dao.GetUserByID(int64(id.(float64)))
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用户不存在"},
		})
		return
	}
	userinfo := &User{
		Id:            user.Id,
		Name:          user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     *userinfo,
	})
}
