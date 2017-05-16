package middleware

//"github.com/rakd/gin_sample/app/models"
//"github.com/gin-gonic/gin"

/*
// APIAuth ...
func APIAuth() gin.HandlerFunc {
	loginNotRequiredPaths := map[string]bool{
		"/api/logout": true,
		"/api/signup": true,
		"/api/login":  true,
	}

	//	loginNotRequiredPathPrefixes := map[string]bool{
	//		"/api/v1/verify":         true,
	//		"/api/v1/forgetpw":       true,
	//		"/api/v1/resetpw":        true,
	//		"/api/v1/request/grant/": true,
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


		//	for path, b := range loginNotRequiredPathPrefixes {
		//		if b {
		//			if strings.HasPrefix(c.Request.URL.Path, path) {
		//				c.Next()
		//				return
		//			}
		//		}
		//	}


		tokenString := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)
		if len(tokenString) == 0 {
			//apiv1handlers.OutputErrorJSON(c, "invalid sig")
			//c.Abort()
			log.Print("invalid sig")
			c.Set("is_login", false)
			c.Next()
			return
		}

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

		if user.InstagramID == "" || user.InstagramUsername == "" {
			//c.JSON(http.StatusUnauthorized, gin.H{"err": "invalid sig or no such user"})
			controllers.OutputErrorJSON(c, "invalid sig or no such user")
			c.Abort()
			return
		}
		c.Set("is_login", true)
		c.Set("access_token", user.AccessToken)
		c.Set("me", user)
		c.Next()
	}

}
*/
