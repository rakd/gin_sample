package controllers

import (
	"log"

	"github.com/rakd/gin_sample/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

//APISignup ...
func APISignup(c *gin.Context) {
	log.Print("APISignup")
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		OutputErrorJSON(c, err.Error())
		return
	}

	user, err := user.Create()
	if err != nil {
		OutputErrorJSON(c, err.Error())
		return
	}

	msg := "signup ok, please login"
	OutputOKJSON(c, msg)
}
