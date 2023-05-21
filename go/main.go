package main

import (
	"gin/common"

	"github.com/gin-gonic/gin"
)

// 业务的处理
func main() {
	DB := common.GetDB()
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
	//1.创建路由
	router := gin.Default()
	//2.创建路由的规则 绑定执行函数
	r := CollectRoute(router)
	//前面的/xxxx是请求的资源如：http://127.0.0.1/xxxx ，后面的这个xxxx调用函数xxxx
	//3.设置监听端口,不写默认8080
	r.Run(":8000") //别忘了:
}
