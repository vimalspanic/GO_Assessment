package images

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const lastfmAPIKey = "4b7791f101520fda28bd4f90ab58429e"

type ArtistImageResponse struct {
	ArtistImages struct {
		ArtistImage []struct {
			URL string `json:"#text"`
		} `json:"artistimage"`
	} `json:"images"`
}

func getArtistImage(artistName string) (string, error) {
	url := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=artist.getimages&artist=%s&api_key=%s&format=json", artistName, lastfmAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var artistImageResponse ArtistImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&artistImageResponse); err != nil {
		return "", err
	}

	// Retrieve the first image URL
	if len(artistImageResponse.ArtistImages.ArtistImage) > 0 {
		return artistImageResponse.ArtistImages.ArtistImage[0].URL, nil
	}

	return "", fmt.Errorf("no image found for artist %s", artistName)
}

func ArtistImageHandler(w http.ResponseWriter, r *http.Request) {
	artistName := r.URL.Query().Get("artist")

	imageURL, err := getArtistImage(artistName)
	if err != nil {
		http.Error(w, "Failed to get artist image", http.StatusInternalServerError)
		log.Println("Failed to get artist image:", err)
		return
	}

	response := map[string]string{
		"imageURL": imageURL,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		log.Println("Failed to encode JSON response:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
