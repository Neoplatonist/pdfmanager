package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

// strings.HasSuffix <- tests ending

func main() {
	engine := echo.New()
	engine.Pre(middleware.AddTrailingSlash())

	engine.SetDebug(true)

	engine.Use(middleware.Logger())
	engine.Use(middleware.Recover())

	g := engine.Group("/upload/*")
	g.Use(middleware.BasicAuth(func(username, password string) bool {
		if username == "admin" && password == "1234" {
			return true
		}
		return false
	}))
	g.File("", "../client/upload.html")
	g.POST("", upload)

	assetHandler := http.FileServer(rice.MustFindBox("../client/pdf").HTTPBox())

	engine.File("/", "../client/index.html")

	engine.GET("/pdf/*", standard.WrapHandler(http.StripPrefix("/pdf/", assetHandler)))

	engine.Run(standard.New(":1337"))
}
