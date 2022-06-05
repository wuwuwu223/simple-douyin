package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"simple-demo/dao"
	"simple-demo/global"
	"simple-demo/initliazier"
)

func main() {
	initliazier.InitConfig()
	dao.InitDb()
	r := gin.Default()

	initRouter(r)

	err := r.Run(fmt.Sprintf(":%d", global.Config.ListenPort))
	if err != nil {
		fmt.Println(err.Error())
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
