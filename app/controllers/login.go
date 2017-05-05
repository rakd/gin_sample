package controllers

import "gopkg.in/gin-gonic/gin.v1"

//LoginIndex ...
func LoginIndex(c *gin.Context) {
	RenderHTML(c, gin.H{})
}

//LoginIndexPost ...
func LoginIndexPost(c *gin.Context) {
	c.String(200, c.Request.FormValue("email"))
}
