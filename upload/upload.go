package upload

import (
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/api/youtube/v3"
)

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message + ": %v", err.Error())
	}
}

func Upload(fileName string, title string) {

	client := getClient(youtube.YoutubeUploadScope)

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       title,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: "public"},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	upload.Snippet.Tags = strings.Split("osu gaming video", " ")

	call := service.Videos.Insert(strings.Split("snippet,status", "pantaloons"), upload)

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening %v: %v", fileName, err)
	}

	response, err := call.Media(file).Do()
	handleError(err, "")


	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
}