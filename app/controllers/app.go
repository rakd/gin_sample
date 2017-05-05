package controllers

import "gopkg.in/gin-gonic/gin.v1"

//AppIndex ...
func AppIndex(c *gin.Context) {
	//RenderTemplate(c, "app/index_amp", gin.H{}, 200)
	RenderHTML(c, gin.H{})
}
