package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/dao"
	"simple-demo/model"
	"simple-demo/utils"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//get user id from jwt token
	token := c.Query("token")
	latest_time := c.Query("latest_time")

	//parse latest_time to int64
	latest_time_int64, _ := strconv.ParseInt(latest_time, 10, 64)

	//parse latest_time to time.Time
	latest_time_time := time.Unix(latest_time_int64, 0)

	userId := utils.GetUserIdFromToken(token)
	var videos []*model.Video
	if latest_time != "" {
		videos, _ = dao.GetVideoListAfterTime(latest_time_time)
	} else {
		videos, _ = dao.GetVideoList()
	}
	var videoList []Video
	for i := range videos {
		var video Video
		video.Id = videos[i].Id
		video.PlayUrl = videos[i].PlayUrl
		video.CoverUrl = videos[i].CoverUrl
		user, _ := dao.GetUserByID(videos[i].UserID)
		video.Author = User{
			Id:            user.Id,
			Name:          user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      dao.CheckIfFollow(userId, user.Id),
			Avatar:        user.Avatar,
		}
		video.FavoriteCount = dao.GetFavoriteCount(videos[i].Id)
		video.IsFavorite = dao.CheckIfFavorite(userId, videos[i].Id)
		video.CommentCount = dao.GetCommentCount(videos[i].Id)
		video.Title = videos[i].Title
		videoList = append(videoList, video)
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
