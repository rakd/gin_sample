package controllers

import "gopkg.in/gin-gonic/gin.v1"

//AppIndex ...
func AppIndex(c *gin.Context) {
	c.String(200, "hello gin")
}
