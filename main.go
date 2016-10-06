package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

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

	g := engine.Group("/*")
	g.Use(middleware.BasicAuth(func(username, password string) bool {
		if username == "admin" && password == "1234" {
			return true
		}
		return false
	}))
	g.File("/upload", "upload.html")
	g.POST("/upload", upload)

	assetHandler := http.FileServer(rice.MustFindBox("pdf").HTTPBox())

	engine.File("/", "index.html")

	engine.GET("/pdf/*", standard.WrapHandler(http.StripPrefix("/pdf/", assetHandler)))

	engine.Run(standard.New(":1337"))
}

func upload(c echo.Context) error {
	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("./pdf/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	c.HTML(http.StatusOK, fmt.Sprintf("<p>File \"%s\" uploaded successfully.</p><br /><a href=\"/\">Go Back</>", file.Filename))

	return nil
}
