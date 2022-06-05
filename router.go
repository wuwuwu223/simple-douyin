package main

import (
	"github.com/gin-gonic/gin"
	"simple-demo/controller"
	"simple-demo/utils"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", utils.JwtMiddleware(), controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", utils.JwtMiddleware(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", utils.JwtMiddleware(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", utils.JwtMiddleware(), controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", utils.JwtMiddleware(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
