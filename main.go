package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/justinas/nosurf"

	"github.com/gin-contrib/sessions"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gin-contrib/gzip"
	"github.com/rakd/gin_sample/app/controllers"
	"github.com/rakd/gin_sample/app/libs/ezgintemplate"
	_ "github.com/rakd/gin_sample/app/models"
	"gopkg.in/gin-contrib/cors.v1"

	"gopkg.in/gin-gonic/gin.v1"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	router := gin.Default()

	//gzip
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	router.StaticFile("/robots.txt", "./assets/robots.txt")

	//router.Use(middleware.AdminGoogleAuth())
	//router.Use(middleware.APIAuth())
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://127.0.0.1:3000",
		},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin", "Accept", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//    return origin == "https://github.com"
		//},
		MaxAge: 12 * time.Hour,
	}))

	// session
	store := sessions.NewCookieStore([]byte("secret1233"))
	router.Use(sessions.Sessions("mysession", store))
	// templates
	render := ezgintemplate.New()
	render.TemplatesDir = "app/views/"
	render.Layout = "layouts/base"
	render.AmpLayout = "layouts/amp"
	render.AdminLayout = "layouts/admin"
	render.Ext = ".tmpl"
	render.Debug = true
	funcMap := template.FuncMap{
		"is_active": func(uri1, uri2 string) template.HTML {
			if uri1 == uri2 {
				return template.HTML("active")
			}
			return ""
		},
	}
	render.TemplateFuncMap = funcMap
	router.HTMLRender = render.Init()

	router.GET("/", controllers.AppIndex)
	router.GET("/search", controllers.SearchIndex)
	router.GET("/flash", controllers.FlashIndex)
	router.GET("/logout", controllers.Logout)
	router.GET("/login", controllers.LoginIndex)
	router.POST("/login", controllers.LoginIndexPost)
	router.GET("/signup", controllers.SignupIndex)
	router.POST("/signup", controllers.SignupIndexPost)
	router.NoRoute(controllers.NoRoute)

	router.POST("/api/me", controllers.APIMe)
	router.POST("/api/login", controllers.APILogin)

	//csrf := nosurf.New(router)
	//csrf.SetFailureHandler(http.HandlerFunc(csrfFailHandler))
	//
	//http.ListenAndServe(":3000", csrf)
	//CSRF := csrf.Protect([]byte("32-byte-long-auth-key"))
	//http.ListenAndServe(":3000", CSRF)
	router.Run(":3000")
}

func csrfFailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", nosurf.Reason(r))
}
