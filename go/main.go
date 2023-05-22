package main

import (
	"database/sql"
	"gin/common"
	"gin/routers"

	"github.com/gin-gonic/gin"
)

// 业务的处理
func main() {
	DB := common.InitDB()
	sqlDB, _ := DB.DB()
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			return
		}

	}(sqlDB)
	//1.创建路由
	router := gin.Default()
	//2.创建路由的规则 绑定执行函数
	r := routers.CollectRoute(router)

	//3.设置监听端口,不写默认8080
	err := r.Run(":8000")
	if err != nil {
		return
	} //别忘了:
}
