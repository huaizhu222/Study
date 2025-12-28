package main

import (
	"net/http"

	"Rpc.Study.go/simple_gin/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, map[string]string{
			"你好": "hello world",
		})
	})

	r.Run(":9999")
}
