package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// YTVideo : contains useful info about a given video
type YTVideo struct {
	ID    string
	Title string
}

func main() {

	var playListID string
	flag.StringVar(&playListID, "id", "", "Playlist ID")

	flag.Parse()

	b, err := ioutil.ReadFile(".APIKEY")
	if err != nil {
		fmt.Print(err)
	}
	apiKey := strings.TrimSuffix(string(b), "\n")

	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		fmt.Print(err)
	}

	call := youtubeService.PlaylistItems.List("snippet")
	call = call.PlaylistId(playListID)
	call = call.MaxResults(50)
	playList, err := call.Do()
	if err != nil {
		fmt.Print(err)
	}

	var videos []YTVideo

	for ok := true; ok; ok = !(playList.NextPageToken == "") {
		for _, item := range playList.Items {
			var video YTVideo
			video.ID = item.Id
			video.Title = item.Snippet.Title
			videos = append(videos, video)
		}

		playList, err = call.PageToken(playList.NextPageToken).Do()
		if err != nil {
			fmt.Print(err)
		}
	}

	jsonVideos, err := json.Marshal(videos)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(jsonVideos))
}
