package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/dao"
	"simple-demo/utils"
	"strconv"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	id := utils.GetUserIdFromJwtToken(c)
	actionType := c.Query("action_type")
	videoId := c.Query("video_id")
	videoIdInt64, _ := strconv.ParseInt(videoId, 10, 64)
	user, _ := dao.GetUserByID(id)
	if actionType == "1" {
		text := c.Query("comment_text")
		err := dao.AddComment(id, videoIdInt64, text)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
			Comment: Comment{
				Id: 1,
				User: User{
					Id:            user.Id,
					Name:          user.Username,
					FollowCount:   user.FollowCount,
					FollowerCount: user.FollowerCount,
					IsFollow:      user.IsFollow,
					Avatar:        user.Avatar,
				},
				Content:    text,
				CreateDate: time.Now().Format("01-02"),
			}})
		return
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoId := c.Query("video_id")
	videoIdInt64, _ := strconv.ParseInt(videoId, 10, 64)
	comments, err := dao.GetComments(videoIdInt64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	var commentList []Comment
	for i := 0; i < len(comments); i++ {
		user, _ := dao.GetUserByID(comments[i].UserID)
		commentList = append(commentList, Comment{
			Id: comments[i].Id,
			User: User{
				Id:            user.Id,
				Name:          user.Username,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      user.IsFollow,
				Avatar:        user.Avatar,
			},
			Content:    comments[i].Content,
			CreateDate: comments[i].CreateDate,
		})
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: commentList,
	})
}
