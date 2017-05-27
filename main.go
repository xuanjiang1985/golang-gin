package main

import (
	//"errors"
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

func main() {

	r := gin.New()
	//open session middleware
	store := sessions.NewCookieStore([]byte("zhougang6233"))
	r.Use(sessions.Sessions("mysession", store))
	//open csrf middleware
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "wang123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "400, timeout status, please return and refresh the page.")
		},
	}))
	//middleware
	r.Use(Logger(), Recover())
	r.Static("/public", "./public")
	r.HTMLRender = pongo2gin.Default()
	r.GET("/", indexC.Get)
	r.GET("/test", forTest)
	r.GET("/article", articleC.Get)
	r.GET("/article/add-thanks/:id", articleC.AddThanks)
	r.POST("/article/store", articleC.Store)
	r.Run(":8080")
}

//middlewares
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request
		c.Next()

		// after request
		latency := time.Since(t)

		// access the status we are sending
		status := c.Writer.Status()
		path := c.Request.URL.Path
		log.Println(status, latency, c.Request.Method, path)
	}
}

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

//test
func forTest(c *gin.Context) {

	c.String(200, "done")
}
