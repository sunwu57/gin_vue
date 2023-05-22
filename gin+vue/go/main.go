package main

import (
	"gin/common"
	"gin/routerss"

	"github.com/gin-gonic/gin"
)

// 业务的处理
func main() {
	DB := common.InitDB()
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
	//1.创建路由
	router := gin.Default()
	//2.创建路由的规则 绑定执行函数
	r := routerss.CollectRoute(router)
	r.Run(":8000") //别忘了:
}
