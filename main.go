package main

import (
	"go-video-steaming/cmd/api/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Config CORS
	e.Use(middleware.CORS())

	// Get token from cookie
	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte("secret"),
	// }))

	e.POST("/upload", handlers.UploadHandler)
	e.DELETE("/delete:id", handlers.DeleteHandler)
	e.GET("/", handlers.GetAllHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
