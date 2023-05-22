package controller

import (
	"gin/common"
	"gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
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
	if len(passwd) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "密码不能小于6位",
		})
		return
	}
	if checkTel(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "手机号已存在",
		})
		return
	}
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "500",
			"msg":  "加密失败",
		})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Passwd:    string(hashedPasswd),
	}
	DB.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

// func Login(c *gin.Context) {
// 	telephone := c.PostForm("telephone")
// 	passwd := c.PostForm("passwd")
// 	if len(telephone) != 11 {
// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"msg": "手机号必须是11位",
// 		})
// 		return
// 	}
// 	if len(passwd) < 6 {
// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"msg": "密码不能小于6位",
// 		})
// 		return
// 	}
// 	DB := common.GetDB()
// 	var user model.User
// 	DB.Where("telephone = ?", telephone).First(&user)
// 	if user.ID == 0 {
// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"msg": "用户不存在",
// 		})
// 		return
// 	}
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(passwd)); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"msg": "密码错误",
// 		})
// 		return
// 	}
// 	token := "11"
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": gin.H{
// 			"token": token,
// 		},
// 		"msg": "登录成功",
// 	})
// }

func checkTel(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
