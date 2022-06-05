package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/dao"
	"simple-demo/utils"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	id := utils.GetUserIdFromJwtToken(c)
	videoid := c.Query("video_id")
	//videoid to int64
	videoidInt64, _ := strconv.ParseInt(videoid, 10, 64)
	action_type := c.Query("action_type")
	err := dao.FavoriteAction(id, videoidInt64, action_type)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
	})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	id := utils.GetUserIdFromJwtToken(c)
	videos, err := dao.GetFavoriteVideoList(id)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	var videolist []Video
	for i := range videos {
		videolist = append(videolist, Video{
			Id:            videos[i].Id,
			Title:         videos[i].Title,
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: dao.GetFavoriteCount(videos[i].Id),
			IsFavorite:    dao.CheckIfFavorite(id, videos[i].Id),
		})
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videolist,
	})
}
