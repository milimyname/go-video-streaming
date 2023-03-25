package handlers

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/iterator"
)

func GetAllHandler(c echo.Context) error {

	// Cloud Storage SDK
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	bucket := client.Bucket("mili-streaming-video")

	// Get the list of files
	it := bucket.Objects(context.Background(), nil)
	var files []string
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		files = append(files, attrs.Name)
	}

	return c.JSON(http.StatusOK, map[string][]string{
		"names": files,
	})
}
