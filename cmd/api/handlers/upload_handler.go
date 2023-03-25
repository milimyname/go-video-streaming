package handlers

import (
	"context"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
)

func UploadHandler(c echo.Context) error {

	form, err := c.MultipartForm()
	if err != nil {
		log.Printf("Failed to parse multipart form: %v", err)
		return c.String(http.StatusBadRequest, "Failed to parse multipart form")
	}

	file, err := form.File["video"][0].Open()
	if err != nil {
		log.Printf("Failed to open file: %v", err)
		return c.String(http.StatusBadRequest, "Failed to open file")
	}
	defer file.Close()

	// Create a new GCS client and bucket handle
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("Failed to create GCS client: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to create GCS client")
	}
	defer client.Close()

	bucketName := "mili-streaming-video"
	bucket := client.Bucket(bucketName)

	// Create a new object in the bucket and upload the file data
	objectName := "videos/" + form.File["video"][0].Filename
	object := bucket.Object(objectName)

	writer := object.NewWriter(ctx)
	defer writer.Close()

	if _, err := io.Copy(writer, file); err != nil {
		log.Printf("Failed to upload file to GCS: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to upload file to GCS")
	}

	// Return a success response
	return c.String(http.StatusOK, "File uploaded successfully")

}
