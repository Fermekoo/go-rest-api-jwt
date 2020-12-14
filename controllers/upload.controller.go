package controllers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
)

func Upload(c echo.Context) error {
	file, err := c.FormFile("gambar")

	if err != nil {
		return err
	}

	src, err := file.Open()

	if err != nil {
		return err
	}

	defer src.Close()

	// dst, err := os.Create(file.Filename)
	dst, err := os.Create(filepath.Join("assets", filepath.Base(file.Filename)))
	if err != nil {
		return err
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
