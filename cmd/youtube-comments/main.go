package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"youtube-comments/pkg/youtubeComments"
)

func main() {
	videoId := flag.String("videoId", "", "video id (the string after '/watch?v=')")
	outputPath := flag.String("outputPath", "comments.json", "output file path")
	help := flag.Bool("h", false, "print help")
	flag.Parse()

	if *help {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *videoId == "" {
		fmt.Println("Missing video ID (Example: for https://www.youtube.com/watch?v=dQw4w9WgXcQ -videoId dQw4w9WgXcQ)")
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	} else {

		client := youtubeComments.NewYoutubeCommentsClient()
		comments, err := client.GetComments(*videoId)
		if err != nil {
			fmt.Println("Failed getting comments for ", *videoId)
		}

		fmt.Println("writing to file...")

		jsonString, _ := json.Marshal(comments)
		ioutil.WriteFile(*outputPath, jsonString, os.ModePerm)
		fmt.Println("Success!")
	}
}
