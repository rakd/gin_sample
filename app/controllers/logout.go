package controllers

import "gopkg.in/gin-gonic/gin.v1"

//Logout ...
func Logout(c *gin.Context) {
	ClearAuth(c)
	msg := "logout"
	SetFlashSuccess(c, msg)
	Redirect(c, "/")
}
