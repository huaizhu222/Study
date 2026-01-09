package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"Rpc.Study.go/simple_gin/gee"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.Run(":9999")
}
