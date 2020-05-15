package main

import (
	"gingo/gingo"
	"log"
	"net/http"
	"time"
)


func onlyForV2() gingo.HandlerFunc {
	return func(c *gingo.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gingo.New()
	r.GET("/index", func(c *gingo.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gingo.Context) {
			c.HTML(http.StatusOK, "<h1>Hello gingo</h1>")
		})

		v1.GET("/hello", func(c *gingo.Context) {
			// expect /hello?name=gingoktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gingo.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}