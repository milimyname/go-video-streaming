package handlers

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
)

func GetOneHandler(c echo.Context) error {

	// Get the file name from the URL
	filename := c.Param("filename")

	// Cloud Storage SDK
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	bucket := client.Bucket("mili-streaming-platform")
	obj := bucket.Object(filename)

	// Get the file
	rc, err := obj.NewReader(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()

	// Get the file size
	attrs, err := obj.Attrs(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Return the file
	return c.Stream(http.StatusOK, attrs.ContentType, rc)
}
