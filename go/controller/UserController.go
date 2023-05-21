package controller

import (
	"gin/common"
	"gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	//ctx.String(http.StatusOK, "%s", "hello")   //可以用占位符，也可以直接字符串
	DB := common.GetDB()
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	passwd := c.PostForm("passwd")
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "手机号必须是11位",
		})
		return
	}
	//判断手机号是否存在
	if checkTel(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "手机号已存在",
		})
		return
	}
	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Passwd:    passwd,
	}
	DB.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}
func checkTel(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
