package controllers

import "gopkg.in/gin-gonic/gin.v1"

//APIMe ...
func APIMe(c *gin.Context) {
	if !IsLogin(c) {
		msg := "no login"
		OutputErrorJSON(c, msg)
		return
	}

	msg := "ok"
	OutputOKJSON(c, msg)
	return
}
