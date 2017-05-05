

# how to make website with gin


I love Gin. cuz it's too simple, I think a lot of ppls wonders what some libraries should be used.
I'm not specialist but tried to make some which are working as productions. I'd like to make this repo to explain how to make website with my experiences for beginners.

## branches

This repo has some branches. would like to keep some branches simple to explain.

- master => active repo with all stuff.
- [x] 01_hello => almost pure gin with assets/glide.
- [x] 02_gzip
- [x] 03_templates => using ezgintemplate, it's supporting switching layouts.
- [ ] 04_flash => sample of flash messages with templates.
- [ ] 05_csrf => supporting csrf, with flash/templates.
- [ ] 06_oauth/admin => google oauth sample for admin pages.
- [ ] 07_cors => cors/JWT sample for APIs.
- [ ] 08_login => login/logout sample, using gorm (db library), with csrf/templates/flash.
- [ ] 09_json => parse json data and showing.
- [ ] 10_docker => docker sample with alpine.
- [ ] 11_cache => using memcached for json. it's including docker-compose sample and json/Unmarshall.
- [ ] 12_deploy => deply sample to ElasticBeanstalk with CircleCI.

-----

## how to start Go

### install Go

You can install Go by Go official website or homebrew.
You might want to manage some Go versions by gvm. It's up to you.


### setup GOPATH

If you'd like to put repository on

> ~/src/github.com/XXX

you should setup GOPATH in .zshrc ( or .profile or .bashrc ) as below.

```
export GOPATH=$HOME
PATH=$PATH:$GOPATH/bin/
```

The path is very important in terms of importing libraries.

### setup glide

Some must wonder how to manage versioning. Some Go libraries are very fast, so Go support versioning.
Go loads
- vendor
firstly, then
- GOPATH
etcetc..

This mean, if you want to use specific versions of libraries, you can put the libraries under vender dir.

But I think it's not enough to manage versions. I strongly recommend to use some managers, like glide.

you should install glide firstly.

All branches have glide.yaml, so you might need to do
```
glide up
```

Then, you can download specific versions under vendor dir.

However, the repo includes vendor dir. so you might not need to `glide up`.

If you want to add more libraries, you can just do
```
glide get LIBRARY
```
instead of
```
go get LIBRARY
```




### setup ATOM with go-plus

I'm using ATOM editor with go-plus.


### install fresh

Go is compile language but you can use it as LL as follow.

```
go run *.go
```

You still need to re-run after rewrite your code.

To avoid re-run, you can use fresh. It enable you to live-reloading.

To install fresh, you can do
```
go get github.com/pilu/fresh
```



-----

## Hello Gin ( 01_hello branch )

It's just `Hello Gin`. please have a look at the branch (hello).

### logging

`log.Print` & `log.Printf` are useful. and adding below, you can see line where you show the log.

```main.go
func init() {
	log.SetFlags(log.Lshortfile)
}
```


### start gin with static assets


```main.go

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	router.StaticFile("/robots.txt", "./assets/robots.txt")
	router.Use(gin.Recovery())
	router.GET("/", controllers.AppIndex)
	router.Run(":3000")
}

```

```app/controllers/app.go
package controllers

import "gopkg.in/gin-gonic/gin.v1"

//AppIndex ...
func AppIndex(c *gin.Context) {
	c.String(200, "hello gin")
}
```

you can just use `fresh`, then you can access http://localhost:3000/



----

## support gzip ( 02_gzip branch )

### import gzip & setting

you should import `"github.com/gin-contrib/gzip"` in main.go, then,

```main.go
router.Use(gzip.Gzip(gzip.DefaultCompression))
```

You can check size of http://localhost:3000/assets/css/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css

-----

## swithcing layouts ( 03_templates branch )

### import ezgintemplate

import `"html/template"` and `"github.com/rakd/gin_sample/app/libs/ezgintemplate"`.

`"github.com/rakd/gin_sample/app/libs/ezgintemplate"` is originally https://github.com/michelloworld/ez-gin-template, customized by kaz. I wanted to switch layout for AMP/Admin/regular pages.


### prepare app/views

under app/views, you should prepare some directories/templates as follwoing.
- layouts/
 - admin.tmpl
 - base.tmpl
 - amp.tmpl
- partials
 - nav.tmpl
- errors
 - 404.tmpl
- app ( one of controllers )
 - index.tmpl

### prepare common logic for controllers and customize App controller.

prepare
- app/controllers/common.go
and edit
- app/controllers/app.go

### prepare admin controller

now, there is no authentication for admin. it's just sample to switch layout.

- app/controllers/admin.go

you can see
-  http://localhost:3000/admin/ with admin layout.

### prepare amp template for AppIndex

- app/views/app/index_amp.tmpl

 then you can see AMP page on http://localhost:3000/?amp=1

My ezgintemplate regard XXX_amp.tmp as tempalte for AMP.


### prepare errors.go and NoRoute in main.go

```main.go
router.NoRoute(controllers.NoRoute)
```

### prepare hoge/index.tmpl without controller

- app/views/about/index.tmpl
- app/views/hoge/index.tmpl
- app/views/search/index.tmpl

there are no controllers for above template. but you can see,
- http://localhost/about
- http://localhost/hoge
- http://localhost/search

If no template, you can see 404 error.

-----

## flash messages ( 04_flash_message branch )

### main.go
```main.go
"github.com/gin-contrib/sessions"
```

```main.go
// session
store := sessions.NewCookieStore([]byte("secret1233"))
router.Use(sessions.Sessions("mysession", store))
```

### app/contorlers/common.go

```app/contorlers/common.go
"github.com/gin-contrib/sessions"
```
```app/contorlers/common.go
// setFlash
data["flash_error"] = GetFlashError(c)
data["flash_warning"] = GetFlashWarning(c)
data["flash_info"] = GetFlashInfo(c)
data["flash_success"] = GetFlashSuccess(c)
```

```app/contorlers/common.go
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
```
### flash message templates

prepare
- app/views/partials/flash.tmpl

### in layouts

```
{{ template "partials/flash" . }}
```

-----


## support csrf ( csrf branch )


----


## support oauth for admin page ( oauth branch )



-----

## cors/JWT ( cors branch )

Some must want to use JWT/cors for APIs.


----


## login/logout ( login branch )

----

## JSON  ( json branch )



----

## docker ( docker branch )

---

## using memcached for JSON with docker-compose ( cache branch )


----

## deply sample with CircleCI/ElasticBeanstalk (deploy branch )
