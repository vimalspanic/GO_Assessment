package musicmatch

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const musixmatchAPIKey = "4b7791f101520fda28bd4f90ab58429e"

type LyricsResponse struct {
	Message struct {
		Body struct {
			Lyrics struct {
				LyricsBody string `json:"lyrics_body"`
			} `json:"lyrics"`
		} `json:"body"`
	} `json:"message"`
}

func GetetLyrics(trackName, artistName string) (string, error) {
	url := fmt.Sprintf("https://api.musixmatch.com/ws/1.1/matcher.lyrics.get?format=json&callback=callback&q_track=%s&q_artist=%s&apikey=%s", "&format=json", "&format=json", musixmatchAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var lyricsResponse LyricsResponse
	if err := json.NewDecoder(resp.Body).Decode(&lyricsResponse); err != nil {
		return "", err
	}

	return lyricsResponse.Message.Body.Lyrics.LyricsBody, nil
}

func LyricsHandler(w http.ResponseWriter, r *http.Request) {
	trackName := r.URL.Query().Get("track")
	artistName := r.URL.Query().Get("artist")

	lyrics, err := GetetLyrics(trackName, artistName)
	if err != nil {
		http.Error(w, "Failed to get lyrics", http.StatusInternalServerError)
		log.Println("Failed to get lyrics:", err)
		return
	}

	response := map[string]string{
		"lyrics": lyrics,
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

// Artist Info
type ArtistResponse struct {
	Message struct {
		Body struct {
			Artist struct {
				ArtistName string `json:"artist_name"`
				URL        string `json:"artist_url"`
			} `json:"artist"`
		} `json:"body"`
	} `json:"message"`
}

func GetArtistInfo(artistName string) (*ArtistResponse, error) {
	api_key := "4b7791f101520fda28bd4f90ab58429e"
	artistName = "Mitski"
	url := fmt.Sprintf("https://api.musixmatch.com/ws/1.1/artist.search?format=json&callback=callback&q_artist=%s&apikey=%s", artistName, api_key)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artistResponse ArtistResponse
	if err := json.NewDecoder(resp.Body).Decode(&artistResponse); err != nil {
		return nil, err
	}

	return &artistResponse, nil
}

func ArtistInfoHandler(w http.ResponseWriter, r *http.Request) {
	artistName := r.URL.Query().Get("artist")

	artistResponse, err := GetArtistInfo(artistName)
	if err != nil {
		http.Error(w, "Failed to get artist information", http.StatusInternalServerError)
		log.Println("Failed to get artist information:", err)
		return
	}

	response := map[string]interface{}{
		"artist": map[string]string{
			"name": artistResponse.Message.Body.Artist.ArtistName,
			"url":  artistResponse.Message.Body.Artist.URL,
		},
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
