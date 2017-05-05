package controllers

import "gopkg.in/gin-gonic/gin.v1"

//AdminIndex ...
func AdminIndex(c *gin.Context) {
	RenderHTML(c, gin.H{})
}
