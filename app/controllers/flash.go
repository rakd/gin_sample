package controllers

import "gopkg.in/gin-gonic/gin.v1"

//FlashIndex ...
func FlashIndex(c *gin.Context) {
	// set flash message.
	SetFlashError(c, "flash message sample")
	// redirect
	Redirect(c, "/")
}
