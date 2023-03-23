package main

import (
	"github.com/webframework/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>hello world</h1>")
	})

	r.POST("/hello", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.GET("/home/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s", c.Param("name"))
	})
	r.Run(":8081")
}
