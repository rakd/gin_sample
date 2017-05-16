package controllers

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/gin-contrib/sessions"

	"github.com/justinas/nosurf"

	"gopkg.in/gin-gonic/gin.v1"
)

// OutputErrorJSON ...
func OutputErrorJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "error",
		"message": msg,
	})

}

// OutputOKJSON ...
func OutputOKJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": msg,
	})
}

// OutputOKDataJSON ...
func OutputOKDataJSON(c *gin.Context, msg string, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": msg,
		"data":    data,
	})
}

// RenderTemplate ...
func RenderTemplate(c *gin.Context, tmpl string, data gin.H, statusCode int) {

	// setFlash
	data["flash_error"] = GetFlashError(c)
	data["flash_warning"] = GetFlashWarning(c)
	data["flash_info"] = GetFlashInfo(c)
	data["flash_success"] = GetFlashSuccess(c)
	data["csrf_token"] = nosurf.Token(c.Request)

	data["is_login"] = IsLogin(c)
	log.Printf("is_login:%v", data["is_login"])
	data["current_uri"] = c.Request.URL.Path

	c.HTML(statusCode, tmpl, data)
}

// RenderHTML ...
func RenderHTML(c *gin.Context, data gin.H) {

	pc, _, _, _ := runtime.Caller(1)
	callerName := runtime.FuncForPC(pc).Name()
	for strings.Contains(callerName, ".") {
		a := strings.Index(callerName, ".")
		callerName = callerName[a+1:]
	}

	tmpl := tmplNameConvert(callerName)
	// check whether the file is existed.
	_, err := os.Stat("app/views/" + tmpl + ".tmpl")
	if err != nil {
		c.String(200, "%s not found", "app/views/"+tmpl+".tmpl")
		return
	}

	// for AMP
	if c.Request.URL.Query().Get("amp") == "1" {
		amptmpl := tmpl + "_amp"
		_, err := os.Stat("app/views/" + amptmpl + ".tmpl")
		if err == nil {
			tmpl = amptmpl
			log.Print(amptmpl)
		}
	}

	RenderTemplate(c, tmpl, data, 200)
}

const flashKeyInfo = "flash_key_info"
const flashKeyError = "flash_key_Error"
const flashKeyWarning = "flash_key_warning"
const flashKeySuccess = "flash_key_Success"

// SetFlashInfo ...
func SetFlashInfo(c *gin.Context, msg string) error {
	return setFlash(c, msg, flashKeyInfo)
}

// GetFlashInfo ...
func GetFlashInfo(c *gin.Context) string {
	return getFlash(c, flashKeyInfo)
}

// SetFlashWarning ...
func SetFlashWarning(c *gin.Context, msg string) error {
	return setFlash(c, msg, flashKeyWarning)
}

// GetFlashWarning ...
func GetFlashWarning(c *gin.Context) string {
	return getFlash(c, flashKeyWarning)
}

// SetFlashError ...
func SetFlashError(c *gin.Context, msg string) error {
	return setFlash(c, msg, flashKeyError)
}

// GetFlashError ...
func GetFlashError(c *gin.Context) string {
	return getFlash(c, flashKeyError)
}

// SetFlashSuccess ...
func SetFlashSuccess(c *gin.Context, msg string) error {
	return setFlash(c, msg, flashKeySuccess)
}

// GetFlashSuccess ...
func GetFlashSuccess(c *gin.Context) string {
	return getFlash(c, flashKeySuccess)
}

func setFlash(c *gin.Context, msg, key string) error {
	session := sessions.Default(c)
	if msg == "" {
		return nil
	}
	session.Set(key, msg)
	session.Save()
	return nil
}

func getFlash(c *gin.Context, key string) string {
	session := sessions.Default(c)
	obj := session.Get(key)
	if msg, ok := obj.(string); ok {
		session.Delete(key)
		session.Save()
		return msg
	}
	return ""
}

// IsLogin ...
func IsLogin(c *gin.Context) bool {

	isLogin, ok := c.Get("is_login")
	if ok && isLogin.(bool) {
		return true
	}

	session := sessions.Default(c)
	if flag := session.Get("is_login"); flag != nil {
		val, ok := flag.(int)
		if ok && val == 1 {
			return true
		}
	}
	return false
}

// SetAuth ...
func SetAuth(c *gin.Context, email string) {
	session := sessions.Default(c)
	session.Set("email", email)
	session.Set("is_login", 1)
	session.Save()
}

// ClearAuth ...
func ClearAuth(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("email")
	session.Delete("is_login")
	session.Save()
}

// Redirect ...
func Redirect(c *gin.Context, url string) {

	c.Redirect(302, url)
}

var tmplCamel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

// tmplNameConvert ...
func tmplNameConvert(s string) string {
	var a []string
	for _, sub := range tmplCamel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	//	log.Print(a)
	if len(a) > 1 {
		//log.Print(a[0])
		//log.Print(a[1:])
		return strings.ToLower(a[0] + "/" + strings.Join(a[1:], "_"))
	}
	return strings.ToLower(a[0])
}
