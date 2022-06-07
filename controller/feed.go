package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/model"
	"simple-demo/service"
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
	latestTime := c.Query("latest_time")

	//parse latest_time to int64
	latestTimeInt64, _ := strconv.ParseInt(latestTime, 10, 64)

	//parse latest_time to time.Time
	latestTimeTime := time.UnixMilli(latestTimeInt64)

	userId := utils.GetUserIdFromToken(token)
	var videos []*model.Video
	if latestTime != "" {
		videos, _ = service.GetVideoListAfterTime(latestTimeTime)
	} else {
		videos, _ = service.GetVideoList()
	}
	var videoList []Video
	for i := range videos {
		var video Video
		video.Id = videos[i].Id
		video.PlayUrl = videos[i].PlayUrl
		video.CoverUrl = videos[i].CoverUrl
		user, _ := service.GetUserByID(videos[i].UserID)
		video.Author = User{
			Id:             user.Id,
			Name:           user.Username,
			FollowCount:    user.FollowCount,
			FollowerCount:  user.FollowerCount,
			IsFollow:       service.CheckIfFollow(userId, user.Id),
			FavoriteCount:  user.FavoriteCount,
			TotalFavorited: user.TotalFavorited,
		}
		video.FavoriteCount = videos[i].FavoriteCount
		video.IsFavorite = service.CheckIfFavorite(userId, videos[i].Id)
		video.CommentCount = videos[i].CommentCount
		video.Title = videos[i].Title
		videoList = append(videoList, video)
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().UnixMilli(),
	})
}
