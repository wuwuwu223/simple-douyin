package controller

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "http://10.0.2.2:8080/static/2_屏幕录制2022-05-31%2023.59.42.mov",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "TestUser",
	Avatar:        "https://avatars.githubusercontent.com/u/97824201?v=4",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
