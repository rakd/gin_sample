package controllers

import (
	"log"

	"github.com/rakd/gin_sample/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

//APILogin ...
func APILogin(c *gin.Context) {

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		OutputErrorJSON(c, err.Error())
		return
	}

	user, err := user.Login()
	if err != nil {
		OutputErrorJSON(c, err.Error())
		return
	}
	token, err := user.CreateJWToken()
	if err != nil {
		OutputErrorJSON(c, err.Error())
		return
	}
	OutputOKDataJSON(c, "login ok", gin.H{
		"token": token,
	})
	return
	/*
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
	*/
}
