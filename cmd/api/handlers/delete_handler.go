package handlers

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
)

func DeleteHandler(c echo.Context) error {

	// Get the id name from the URL
	id := "videos" + c.Param("id")

	// Log the id
	log.Println(id)

	// Cloud Storage SDK
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	bucket := client.Bucket("mili-streaming-video")
	obj := bucket.Object(id)

	// Delete the file
	if err := obj.Delete(context.Background()); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
