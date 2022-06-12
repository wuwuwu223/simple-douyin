package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"simple-demo/global"
	"simple-demo/model"
	"simple-demo/service"
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
			StatusMsg:  "上传文件失败",
		})
		return
	}
	title := c.PostForm("title")
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s_%d", id, filename, time.Now().Unix())
	saveFile := filepath.Join("./public/", finalName)
	var video *model.Video
	if !global.Config.UseCos {
		video = &model.Video{
			UserID:   id,
			Title:    title,
			PlayUrl:  global.Config.BaseUrl + finalName,
			CoverUrl: "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		}
		if err = c.SaveUploadedFile(data, saveFile); err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "保存文件失败",
			})
			return
		}
	} else {
		err = utils.UploadVideoToCos(data, finalName)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "上传文件失败",
			})
			return
		}
		video = &model.Video{
			UserID:   id,
			Title:    title,
			PlayUrl:  global.Config.Cos.Address + "/" + finalName,
			CoverUrl: global.Config.BaseUrl + finalName + ".jpg",
		}
	}
	err = service.AddVideo(video)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "发布失败",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + "上传成功",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	id := utils.GetUserIdFromToken(token)
	userid := c.Query("user_id")
	useridInt64, _ := strconv.ParseInt(userid, 10, 64)
	user, err := service.GetUserByID(useridInt64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	videos, err := service.GetVideoListByUserID(useridInt64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "获取视频列表失败",
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
			IsFollow:      service.CheckIfFollow(id, user.Id),
		}
		video.Title = videos[i].Title
		video.CommentCount = videos[i].CommentCount
		video.FavoriteCount = videos[i].FavoriteCount
		video.IsFavorite = service.CheckIfFavorite(id, videos[i].Id)
		videolist = append(videolist, video)
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videolist,
	})
}
