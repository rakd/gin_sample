package controllers

import (
	"os"
	"strings"

	"gopkg.in/gin-gonic/gin.v1"
)

//NoRoute ....
func NoRoute(c *gin.Context) {

	// even though there is no controllers, we can show page as long as there is a view.
	tmpl := c.Request.URL.Path
	if string(tmpl[0]) == "/" {
		tmpl = tmpl[1:]
	}
	if string(tmpl[len(tmpl)-1]) == "/" {
		tmpl = tmpl[:len(tmpl)-1]
	}
	if !strings.Contains(tmpl, "/") {
		tmpl += "/index"
	}
	_, err := os.Stat("app/views/" + tmpl + ".tmpl")
	if err == nil {
		// there is a view, so we can show the page.
		RenderTemplate(c, tmpl, gin.H{}, 200)
		return
	}

	// there is no view, so we are show error page.
	Error404(c)
	//SetFlashError(c, "page not found")
	//Redirect(c, "/")
}

//Error404 ....
func Error404(c *gin.Context) {
	RenderTemplate(c, "errors/404", gin.H{}, 404)
}
