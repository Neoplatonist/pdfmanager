package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	engine := echo.New()

	engine.SetDebug(true)

	engine.Use(middleware.Logger())
	engine.Use(middleware.Recover())

	assetHandler := http.FileServer(rice.MustFindBox("pdf").HTTPBox())

	engine.File("/", "index.html")

	engine.GET("/pdf/*", standard.WrapHandler(http.StripPrefix("/pdf/", assetHandler)))

	engine.POST("/upload", upload)

	engine.Run(standard.New(":1337"))
}
