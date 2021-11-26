package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	auth := func(c *gin.Context){
		//获取cookie
		value, err := c.Cookie("gin_cookie")
		if err != nil {
			c.JSON(403,gin.H{
				"message":"认证失败,没有cookie",
			})
			//认证失败
			c.Abort()
		}else{
			//将获取到的cookie的值写入上下文
			c.Set("cookie",value)
		}
	}
	r.POST("/enroll",func (c *gin.Context){
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		c.SetCookie("gin_cookie", username, 3600, "/", "", false, true)
		c.SetCookie("gin_password", password, 3600, "/", "", false, true)
		c.JSON(200,gin.H{
				"msg": "enroll successfully",
		})
	})//注册部分
	r.POST("/login",func (c *gin.Context){
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		value1, _ := c.Cookie("gin_cookie")
		value2, _ := c.Cookie("gin_password")
		if username == value1 && password == value2{
			c.SetCookie("gin_cookie", username, 3600, "/", "", false, true)
			c.JSON(200,gin.H{
				"msg": "login successfully",
			})
		}else{
			c.JSON(403,gin.H{
				"message":"认证失败,账号密码错误",
			})
		}
	})//登入部分
	r.POST("/logout",func (c *gin.Context){
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		value1, _ := c.Cookie("gin_cookie")
		value2, _ := c.Cookie("gin_password")
		if username == value1 && password == value2{
			c.SetCookie("gin_cookie", username, -1, "/", "", false, true)
			c.JSON(200,gin.H{
				"msg": "logout successfully",
			})
		}else{
			c.JSON(403,gin.H{
				"message":"退出失败,账号密码错误",
			})
		}
	})//登出部分
	//在中间放入鉴权中间件
	r.GET("/hello",auth, func(c *gin.Context) {
		cookie,_:=c.Get("cookie")
		str:=cookie.(string)
		c.String(200,"hello world"+str)
	})
	r.Run()
}