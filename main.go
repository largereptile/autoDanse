package main

import (
	"barold.dev/render"
	"barold.dev/upload"
	"fmt"
	"log"
	"os"
)

const osuReplayDir = "/home/harry/.local/share/osu-wine/osu!/Replays/"

func main() {

	fmt.Println("Checking osu! replay directory")
	files, err := os.ReadDir(osuReplayDir)
	if err != nil {
		log.Fatal(err)
	}

	if len(files) > 0 {
		fmt.Println("Found files - processing")
	}

	for _, file := range files {
		fmt.Println("Attempting to render " + file.Name())
		filename := render.Render(osuReplayDir + file.Name())
		fmt.Println("Rendered file: " + filename)
		upload.Upload("./danser/videos/" + filename, file.Name())
	}

	//for range time.Tick(time.Minute) {
	//
	//}
}
