package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"simple-demo/dao"
	"simple-demo/model"
	"simple-demo/utils"
	"strconv"
	"time"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	id := utils.GetUserIdFromToken(token)

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	title := c.PostForm("title")
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s_%d", id, filename, time.Now().Unix())
	saveFile := filepath.Join("./public/", finalName)
	video := &model.Video{
		UserID:   id,
		Title:    title,
		PlayUrl:  "http://10.0.2.2:8080/static/" + finalName,
		CoverUrl: "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
	}
	err = dao.AddVideo(video)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	id := utils.GetUserIdFromToken(token)
	userid := c.Query("user_id")
	useridInt64, _ := strconv.ParseInt(userid, 10, 64)
	user, err := dao.GetUserByID(useridInt64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	videos, err := dao.GetVideoListByUserID(useridInt64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	var videolist []Video
	for i := range videos {
		var video Video
		video.Id = videos[i].Id
		video.PlayUrl = videos[i].PlayUrl
		video.CoverUrl = videos[i].CoverUrl
		video.Author = User{
			Id:            user.Id,
			Name:          user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      dao.CheckIfFollow(id, user.Id),
		}
		video.FavoriteCount = dao.GetFavoriteCount(videos[i].Id)
		video.IsFavorite = dao.CheckIfFavorite(id, videos[i].Id)
		videolist = append(videolist, video)
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videolist,
	})
}
