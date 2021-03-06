package controllers

import "github.com/gin-gonic/gin"

type User struct {
	Basic
}

// 用户登录页面
func (a *User) Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(200, "login/login.html", gin.H{
			"message": "",
		})
	} else if c.Request.Method == "POST" {
		username := c.Request.PostFormValue("username")
		password := c.Request.PostFormValue("password")
		c.JSON(200, gin.H{
			//"hello": session.Get("mysession"),
			"username": username,
			"password": password,
		})
	}
}

// 验证用户是否登录
func (a *User) IsLogin(c *gin.Context) {
}

// 添加用户
func (a *User) AddUser(c *gin.Context) {
}

// 修改用户
func (a *User) UpdateUser(c *gin.Context) {
}

// 删除用户
func (a *User) DelUser(c *gin.Context) {
}
