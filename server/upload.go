package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

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
	dst, err := os.Create("../client/pdf/" + file.Filename)
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
