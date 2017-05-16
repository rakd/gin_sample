package controllers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/satori/go.uuid"
	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
	"gopkg.in/gin-gonic/gin.v1"
)

var googleOAuthConf oauth2.Config

const googleOAuthCookieName = "google-websta-cokies"

const callbackURI = "/auth/google/callback"

var googleOAuthOK bool

func init() {

	clientID := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	googleOAuthOK = false

	if clientID != "" && clientSecret != "" {
		googleOAuthOK = true
		googleOAuthConf = oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       []string{"openid", "email", "profile"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://accounts.google.com/o/oauth2/v2/auth",
				TokenURL: "https://www.googleapis.com/oauth2/v4/token",
			},
		}
	}

}

// GoogleLogin ....
func GoogleLogin(c *gin.Context) {

	state := uuid.NewV4().String()
	log.Print("state=" + state)
	session := sessions.Default(c)
	session.Set(googleOAuthCookieName, state)
	session.Set("google_oauth_email", nil)
	uri := c.Request.URL.Path
	if strings.HasPrefix(uri, "/") && !strings.HasPrefix(uri, "/api") && !strings.HasPrefix(uri, "/auth") {
		session.Set("return_uri", uri)
	}

	session.Save()

	redierctURL := getCallbackURL(c)
	log.Print(redierctURL)
	googleOAuthConf.RedirectURL = redierctURL
	url := googleOAuthConf.AuthCodeURL(state)
	c.Redirect(302, url)
	c.Abort()
	return
}

// GoogleCallback ...
func GoogleCallback(c *gin.Context) {
	log.Print("GoogleCallback")
	if googleOAuthOK == false {
		c.String(200, "no setup clientID or clientSecret for google auth")
		c.Abort()
		return
	}

	session := sessions.Default(c)
	session.Get(googleOAuthCookieName)
	stateOrig := session.Get(googleOAuthCookieName)

	log.Print(stateOrig)
	if stateOrig == nil || stateOrig.(string) == "" {
		c.String(200, "no session")
		c.Abort()
		return
	}
	state := stateOrig.(string)

	state2 := c.Request.URL.Query().Get("state")
	if state != state2 {
		c.String(200, "state is not match. state="+state+", state2="+state2)
		c.Abort()
		return
	}

	redierctURL := getCallbackURL(c)
	log.Print(redierctURL)
	googleOAuthConf.RedirectURL = redierctURL
	log.Print("")
	// get auth code
	code := c.Request.URL.Query().Get("code")
	tok, err := googleOAuthConf.Exchange(c, code)
	if err != nil {
		c.String(200, "Exchange error. cannot get token from code : "+err.Error())
		c.Abort()
		return
	}

	// check whether the token is valid
	if tok.Valid() == false {
		c.String(200, "token is invalid.")
		c.Abort()
		return
	}

	// get oauth2 clinet service
	// // if you don't need to get user information, we can skip it.
	service, err := v2.New(googleOAuthConf.Client(c, tok))
	if err != nil {
		c.String(200, err.Error())
		c.Abort()
		return
	}

	// get token info
	// it has email & user id, etc. if you don't need to get it, we can skip it.
	tokenInfo, err := service.Tokeninfo().AccessToken(tok.AccessToken).Context(c).Do()
	if err != nil {
		c.String(200, err.Error())
		c.Abort()
		return
	}

	//helpers.StoreGoogleAuthEmail(c, tokenInfo.Email)
	session.Set("google_oauth_email", tokenInfo.Email)
	returnURI := session.Get("return_uri")
	session.Set("return_uri", "")
	session.Save()

	if returnURI != nil {
		if uri, ok := returnURI.(string); ok {
			if uri != "" && strings.HasPrefix(uri, "/") {
				c.Redirect(302, uri)
				c.Abort()
				return
			}
		}
	}

	c.Redirect(302, "/admin")
	c.Abort()
	return

	//c.String(200, "okok")
}

//func getCallbackURL(scheme, host string) string {
func getCallbackURL(c *gin.Context) string {

	scheme := "https"
	if strings.Contains(strings.ToLower(c.Request.Proto), "https") {
		scheme = "https"
	} else {
		scheme = "http"
	}

	log.Print(scheme)
	return fmt.Sprintf(scheme+"://%s%s", c.Request.Host, callbackURI)
}

// GoogleLogout ...
func GoogleLogout(c *gin.Context) {
	if googleOAuthOK == false {
		c.String(200, "no setup clientID or clientSecret for google auth")
		c.Abort()
		return
	}

	session := sessions.Default(c)
	session.Clear()
	session.Set("google_oauth_email", nil)
	session.Clear()
	session.Save()
	c.Redirect(302, "/")
	c.Abort()
	return
}
