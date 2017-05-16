package controllers

import (
	"log"

	"github.com/rakd/gin_sample/app/libs/igsearch"
	"gopkg.in/gin-gonic/gin.v1"
)

// SearchIndex ...
func SearchIndex(c *gin.Context) {
	q := c.Request.URL.Query().Get("q")
	var err error
	log.Print(q)
	var res igsearch.IGSearchResult
	if q != "" {
		log.Print(q)
		res, err = igsearch.UserSearch(q)
		log.Print(q)
		if err != nil {
			SetFlashError(c, err.Error())
			Redirect(c, "/search")
			return
		}
		log.Print(res)
	}
	if len(res.Users) > 0 {
		log.Printf(res.Users[0].User.Username)
	}
	RenderHTML(c, gin.H{
		"result": res,
	})
}
