package controllers

import (
	"log"

	"github.com/rakd/gin_sample/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

//SignupIndex ...
func SignupIndex(c *gin.Context) {

	RenderHTML(c, gin.H{})
}

//SignupIndexPost ...
func SignupIndexPost(c *gin.Context) {
	log.Print("SignupIndexPost")
	email := c.Request.PostFormValue("email")
	password := c.Request.PostFormValue("password")
	log.Print(email + password)

	user := models.NewUser()
	user.Email = email
	user.Password = password

	user, err := user.Create()
	if err != nil {
		msg := "signup error"
		SetFlashError(c, msg)
		Redirect(c, "/signup")
		return
	}

	msg := "signup ok, please login"
	SetFlashSuccess(c, msg)
	Redirect(c, "/login")
}
