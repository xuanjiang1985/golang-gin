package main

import (
	//"errors"
	"errors"
	"github.com/robvdl/pongo2gin"
	"golang-gin/controller"
	"golang-gin/csrf"
	"golang-gin/sessions"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"time"
)

//import controllers
var indexC *ctrl.IndexController
var articleC *ctrl.ArticleController
var authC *ctrl.AuthController

func main() {

	r := gin.New()
	//open session middleware
	store := sessions.NewCookieStore([]byte("zhougang6233"))
	r.Use(sessions.Sessions("mysession", store))
	//open csrf middleware
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "wang123",
		ErrorFunc: func(c *gin.Context) {
			c.String(403, "403, token mismatch, please return and refresh the page.")
		},
	}))
	r.GET("/test", forTest)
	r.Use(AuthMiddleware())
	//open log middleware
	r.Use(Logger())
	r.Static("/public", "./public")
	r.HTMLRender = pongo2gin.Default()
	r.GET("/", indexC.Get)
	r.GET("/article", articleC.Get)
	r.GET("/article/add-thank/:id", articleC.AddThank)
	r.POST("/article/add-comment", articleC.AddComment)
	r.GET("/article/get-comments/:id", articleC.GetComments)
	r.POST("/article/store", articleC.Store)
	r.GET("/register", authC.GetRegister)
	r.GET("/logout", authC.GetLogout)
	r.POST("/register", authC.PostRegister)
	r.POST("/login", authC.PostLogin)
	authorized := r.Group("/auth")
	authorized.Use(AuthRequired())
	{
		authorized.GET("/setting/name", authC.GetSettingName)
		authorized.POST("/setting/name", authC.PostSettingName)
		authorized.GET("/setting/sex", authC.GetSettingSex)
		authorized.POST("/setting/sex", authC.PostSettingSex)
		authorized.GET("/setting/header", authC.GetSettingHeader)
		authorized.POST("/setting/header", authC.PostSettingHeader)
		authorized.GET("/setting/password", authC.GetSettingPassword)
	}
	r.Run(":8081")
}

//middlewares
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		//c.Set("example", "12345")

		// before request
		c.Next()

		// after request
		latency := time.Since(t)

		// access the status we are sending
		status := c.Writer.Status()
		path := c.Request.URL
		log.Println(status, latency, c.Request.Method, path)
	}
}

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// if Auth
func AuthMiddleware() gin.HandlerFunc {
	type HasAuth struct {
		Check    bool
		UserName string
		UserId   int
		Sex      int
		Header   string
	}
	return func(c *gin.Context) {
		session := sessions.Default(c)
		//authValue := session.Get("authUserName")
		if s, ok := session.Get("authUserName").(string); ok && len(s) != 0 {
			user := HasAuth{true, s, session.Get("authUserId").(int), session.Get("authUserSex").(int), session.Get("authUserHeader").(string)}
			c.Set("authUser", user)
			//log.Println(authValue)
		} else {
			user := HasAuth{false, "", 0, 0, ""}
			c.Set("authUser", user)
		}
		c.Next()
	}
}

//authorized
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		//authValue := session.Get("authUserName")
		if s, ok := session.Get("authUserName").(string); ok && len(s) != 0 {
			c.Next()
		} else {
			c.AbortWithError(401, errors.New("没有权限访问。"))
		}
	}
}

//test
func forTest(c *gin.Context) {
	session := sessions.Default(c)
	data := session.Flashes("authUserName", "authUserId", "authUserSex", "authUserHeader")
	log.Println(data)
	c.String(200, "clear sessions")
}
