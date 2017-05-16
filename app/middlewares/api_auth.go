package middleware

import "github.com/rakd/gin_sample/app/models"
import "github.com/rakd/gin_sample/app/config"
import "github.com/gin-gonic/gin"
import "strings"
import "log"
import jwt "github.com/dgrijalva/jwt-go"

// APIAuth ...
func APIAuth() gin.HandlerFunc {
	loginNotRequiredPaths := map[string]bool{
		"/api/signup": true,
		"/api/login":  true,
		"/api/csrf":   true,
	}

	//	loginNotRequiredPathPrefixes := map[string]bool{
	//		"/api/verify":         true,
	//		"/api/forgetpw":       true,
	//		"/api/resetpw":        true,
	//		"/api/request/grant/": true,
	//	}

	return func(c *gin.Context) {
		if loginNotRequiredPaths[c.Request.URL.Path] {
			c.Next()
			return
		}

		if !strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.Next()
			return
		}

		/*
			for path, b := range loginNotRequiredPathPrefixes {
				if b {
					if strings.HasPrefix(c.Request.URL.Path, path) {
					c.Next()
						return
					}
				}
			}
		*/

		tokenString := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)
		if len(tokenString) == 0 {
			//apiv1handlers.OutputErrorJSON(c, "invalid sig")
			//c.Abort()
			log.Print("invalid sig")
			c.Set("is_login", false)
			c.Next()
			return
		}

		//user := models.JWTUser{}
		user := models.User{}
		_, err := jwt.ParseWithClaims(tokenString, &user, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetJWTSalt()), nil
		})

		if err != nil {
			//apiv1handlers.OutputErrorJSON(c, "JWT error")
			//c.Abort()
			log.Print("JWT error")
			c.Set("is_login", false)
			c.Next()
			return
		}

		c.Set("is_login", true)
		c.Set("me", user)
		c.Next()
	}

}
