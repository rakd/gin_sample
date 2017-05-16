package controllers

import "github.com/rakd/gin_sample/app/models"
import "gopkg.in/gin-gonic/gin.v1"

//LoginIndex ...
func LoginIndex(c *gin.Context) {
	RenderHTML(c, gin.H{})
}

//LoginIndexPost ...
func LoginIndexPost(c *gin.Context) {
	user := models.User{}
	user.Email = c.Request.PostFormValue("email")
	user.Password = c.Request.PostFormValue("password")

	user, err := user.Login()
	if err != nil {
		msg := err.Error()
		SetFlashError(c, msg)
		Redirect(c, "/login")
		return
	}

	msg := "login ok"
	SetAuth(c, user.Email)
	SetFlashSuccess(c, msg)
	Redirect(c, "/")
}
