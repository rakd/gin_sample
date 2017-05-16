

# how to make website with gin


I love Gin. cuz it's too simple, I think a lot of ppls wonders what some libraries should be used.
I'm not specialist but tried to make some which are working as productions. I'd like to make this repo to explain how to make website with my experiences for beginners.

## branches

This repo has some branches. would like to keep some branches simple to explain.

- master => active repo with all stuff.
- [x] 01_hello => almost pure gin with assets/glide.
- [x] 02_gzip
- [x] 03_templates => using ezgintemplate, it's supporting switching layouts.
- [x] 04_flash => sample of flash messages with templates.
- [x] 05_csrf => supporting csrf, with flash/templates.
- [ ] 06_oauth => google oauth sample for admin pages.
- [x] 07_login => login/logout sample, using gorm (db library), with csrf/templates/flash.
- [x] 08_cors => cors/JWT sample for APIs.
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

But I think it's not enough to manage versions. I strongly recommend you to use some managers, like glide.

you should install glide firstly. You can install by  
```
curl https://glide.sh/get | sh
```
OR
```
brew install glide
```

Please check the official repo of glide. ref: https://github.com/Masterminds/glide


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

### install mysql

### install docker



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
- http://localhost:3000/about
- http://localhost:3000/hoge
- http://localhost:3000/search

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

### flash message test

access http://localhost:3000/flash, then you will be redirect to / with the message.

-----


## support csrf ( 05_csrf branch )

### main.go

add import like this.
```
import "github.com/justinas/nosurf"
```

and add this.
```
func csrfFailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", nosurf.Reason(r))
}
```

then, Replace `router.Run(":3000")` to
```
csrf := nosurf.New(router)
csrf.SetFailureHandler(http.HandlerFunc(csrfFailHandler))
http.ListenAndServe(":3000", csrf)
```

### app/controllers/common.go

in common.go, you should add import,
```
import 	"github.com/justinas/nosurf"
```

and add below in `RenderTemplate`,
```
data["csrf_token"] = nosurf.Token(c.Request)
```

### all POST forms in templates

```
{{if .csrf_token}}<input type="hidden" name="csrf_token" value="{{.csrf_token}}" />{{end}}
```

### prepare login controllers/views

I've prepared login controllers and views, you can try to acccess/POST on http://localhost:3000/login


----


## support oauth for admin page ( 06_oauth branch )

### set oauth client ID and Secret

You need to get OAUTH clientID & secret on console.cloud.google.com, and set env in .zshrc ( or .bashrc or .profile ).
After access above URL, you can create projects. `API Manager => Credentials => Create Credentials => Oauth Client ID => Web Application` helps you to create client ID and secret.

On the setting, you should add callback URls as `Authorized redirect URIs`. Please add `http://localhost:3000/auth/google/callback` for test.


```.zshrc
export GOOGLE_OAUTH_CLIENT_ID="YOUR_CLIENT_ID"
export GOOGLE_OAUTH_CLIENT_SECRET="YOUR_SECRET_ID"
```

### prepare google oauth middleware

When users access admin pages, we should check cookies whether it has google email address.
If it doesn't not have it, users should be redirected to google login page.
I've prepared the middleware in `app/middlewares/admin_google_oauth.go`.


To use the middleware, you need to amend main.go a little.

```main.go
import "github.com/rakd/gin_sample/app/middlewares"
```


```main.go
	router.Use(middleware.AdminGoogleAuth())
```

It accepts only email addresses what you write in admin_google_oauth.go, so you need to add your email address.


### /auth/google/login, /auth/google/callback

After user's login on google login page, users will be redirected to our callback page with OAUTH code.
We should check the code on callback page.

For the google login page and callback, I've prepared `app/controllers/auth_google.go`.



### accesss admin page

you can try to access http://localhost:3000/admin

----


## login/logout ( 07_login branch )


### launch mysql server on your MacOSX.

If you installed mysql-server by homebrew, you can launch mysql server like this.
```
mysql.server start
```

### setup envs for DB on your local

DB_HOST
DB_

### prepare User model and database.

I've prepared `app/models/user.go` for User models.

There are some libraries to access database(mysql). Actually, you might want to use simple mysql library without ORM. (https://github.com/go-sql-driver/mysql). I have no objection, I know some love it cuz simple and fast.
But I'm using gorm as ORM. There is no big reason to use it, it's just coz my comfortable, especially for migrations.

I've preapred `app/models/common.go` to connect mysql.s

refs:
- https://github.com/jinzhu/gorm
- http://jinzhu.me/gorm/


### main.go

To use mysql specifically, you need to write below in main.go
```
import _ "github.com/jinzhu/gorm/dialects/mysql"
```

You may realize it's including `_` before the path. This mean the code loads init logic of the library, even though not using the library in main.go. the init logic tells you are using mysql dialect in gorm.


### login & logout controllers

### amend template

To show login status, I've amended templates.

### try to login.

you can try to login on http://localhost:3000/login


-----

## cors/JWT ( 08_cors branch )

Some must want to use JWT/cors for APIs.


### setup envs on your local

- JWT_SALT

### main.go

### API Auth middleware.

You need to check JWT when users access APIs. there are some exceptions, for example, it's login.
Before login, users should not need JWT. `app/middleware/api_auth.go` which I've prepared doesn't check `/api/login`.

You might want to add more exceptions, please try to add some if you want.

### controllers for POST:/api/login and POST:/api/me

`app/controllers/api_login.go` checks authentication as POST (JSON format), it returns JWT when the credential is correct.

`app/controllers/api_me.go` checks JWT to show user information.

### try to login and access via APIs.


#### localhost:3000/api/csrf


#### http://localhost:3000/api/login

```
curl -X POST -v -d "{\"email\": \"rakd0930@gmail.com\", \"password\":\"rakdrakd\"}"  localhost:3000/api/login
```

```
{"data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZW1haWwiOiJyYWtkMDkzMEBnbWFpbC5jb20iLCJwYXNzd29yZCI6IiIsInRva2VuIjoiMjAzZmIzZTY5MzFkNGNkOWE3NjMxM2U0ZjAzNWExYzYiLCJ2ZXJpZnkiOmZhbHNlfQ.x0KMAdiumaL8T3V8b6s_ZM8EEaxHtLo0H53VKBJ50ig"},"message":"login ok","status":"ok"}
```

#### http://localhost:3000/api/me


```
curl -X POST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZW1haWwiOiJyYWtkMDkzMEBnbWFpbC5jb20iLCJwYXNzd29yZCI6IiIsInRva2VuIjoiMjAzZmIzZTY5MzFkNGNkOWE3NjMxM2U0ZjAzNWExYzYiLCJ2ZXJpZnkiOmZhbHNlfQ.x0KMAdiumaL8T3V8b6s_ZM8EEaxHtLo0H53VKBJ50ig" localhost:3000/api/v1/me
```
----


## JSON  ( 09_json branch )



----

## docker ( 10_docker branch )


### Dockerfile

We are using alpine linux, it's small and simple. You will be able to understand the size is very comfortable when deploying.

- ref: https://alpinelinux.org/

With the Dockerfile, you can build image as follows.

```
docker build
```

### prepare Makefile

Go supports cross-compile, this mean you can make linux binary on MacOS.

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags netgo -a -v -o main *.go
```

It's pain in the neck to type every time you want to make it, so I'm using Makefile. We can use it for docker build as well.

With Makefile, you can build linux binary as follow.
```
make go
```

For docker image, like this.
```
make docker
```


### check docker image after the docker build


you can check the image,
```
docker images
```

---

## using cache, for JSON with docker-compose ( 11_cache branch )




----

## deploy sample with CircleCI/ElasticBeanstalk (12_deploy branch )

### IAM for deployment


### create docker repos on AWS. for your docker images.


### setup env vars on CircleCI ( and on your local )


You need to set the environment vars on CircleCI.

If you want to deploy from your local, you should set the envs on your local as well.

### setup VPC

### setup RDS

### setup ElasticBeanstalk

#### multi containers? single container?

#### envs on ElasticBeanstalk

you need to set some environment vars
- GOOGLE_OAUTH_CLIENT_ID
- GOOGLE_OAUTH_CLIENT_SECRET
- DB_HOST
- DB_USER
- DB_PASS


### circle.yml

circle.yml

### Dockerrun.json.template


### scripts/deploy.sh
