package middleware

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-contrib/sessions"
	"gopkg.in/gin-gonic/gin.v1"
	//"github.com/gin-gonic/gin"
)

const loginURI = "/auth/google/login"

func isValidAdminEmail(email string) bool {
	adminEmails := map[string]bool{}
	adminEmails["rakd0930@gmail.com"] = true
	adminEmails["ohno@a-fis.com"] = true

	return adminEmails[email] == true
}

// GetGoogleAuthEmail ...
func GetGoogleAuthEmail(c *gin.Context) (string, error) {
	session := sessions.Default(c)
	googleEmail := session.Get("google_oauth_email")

	if googleEmail == nil || googleEmail.(string) == "" {
		log.Print("need to login via google oauth")
		return "", fmt.Errorf("need to login via google oauth")
	}
	if !isValidAdminEmail(googleEmail.(string)) {
		return "", fmt.Errorf("invalid email address")
	}
	return googleEmail.(string), nil
}

// AdminGoogleAuth ...
func AdminGoogleAuth() gin.HandlerFunc {

	loginNotRequiredPathPrefixes := map[string]bool{
		"/api":                  true,
		"/error":                true,
		"/auth/google/logout":   true,
		"/auth/google/login":    true,
		"/auth/google/callback": true,
		"assets":                true,
		"favicon.ico":           true,
		"robots.txt":            true,
	}

	return func(c *gin.Context) {
		/*
			if strings.HasPrefix(c.Request.URL.Path, "/auth/google/logout") {
				RedirectGoogleLogin(c)
				c.Abort()
				return
			}
			if strings.HasPrefix(c.Request.URL.Path, "/auth/google/login") {
				RedirectGoogleLogin(c)
				c.Abort()
				return
			}

			if strings.HasPrefix(c.Request.URL.Path, "/auth/google/callback") {
				GoogleCallback(c)
				c.Abort()
				return
			}
		*/
		for path, b := range loginNotRequiredPathPrefixes {
			if b {
				if strings.HasPrefix(c.Request.URL.Path, path) {
					c.Next()
					return
				}
			}
		}
		log.Print("")
		googleEmail, err := GetGoogleAuthEmail(c)
		log.Print(googleEmail)
		//log.Print(err)

		if err != nil {
			log.Print(err)
			/*
				scheme := "http"
				if strings.Contains(strings.ToLower(c.Request.Proto), "https") {
					scheme = "https"
				}*/
			redierctURL := getLoginURL(c)
			c.Redirect(302, redierctURL)
			return
		}
		c.Set("googleEmail", googleEmail)

		//log.Print(googleEmail)

		state := c.Request.URL.Query().Get("state")
		code := c.Request.URL.Query().Get("code")
		log.Print(state)
		log.Print(code)

	}
}

//func getLoginURL(scheme, host string) string {
func getLoginURL(c *gin.Context) string {

	scheme := "https"
	if strings.Contains(strings.ToLower(c.Request.Proto), "https") {
		scheme = "https"
	} else {
		scheme = "http"
	}
	log.Print(scheme)
	return fmt.Sprintf(scheme+"://%s%s", c.Request.Host, loginURI)
}
