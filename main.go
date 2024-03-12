package main

import (
	images "artist_music/image"
	"artist_music/lastfavourite"
	"artist_music/musicmatch"
	"fmt"
	"log"
	"net/http"
)

func main() {
	apiKey := "4b7791f101520fda28bd4f90ab58429e"
	track := "the boy is mine"
	artist := "Some Artist"
	apiEndpoint := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=track.addTags&api_key=%s&artist=%s&track=%s&tags=rock&format=json", apiKey, artist, track)

	response, err := http.Get(apiEndpoint)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	http.HandleFunc("/top-tracks", lastfavourite.TopTracksHandler)
	http.HandleFunc("/lyrics", musicmatch.LyricsHandler)
	http.HandleFunc("/artist-info", musicmatch.ArtistInfoHandler)
	http.HandleFunc("/image", images.ArtistImageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
