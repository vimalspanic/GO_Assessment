package lastfavourite

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Track struct {
	Name       string `json:"name"`
	Duration   string `json:"duration"`
	Playcount  string `json:"playcount"`
	Listeners  string `json:"listeners"`
	MBID       string `json:"mbid"`
	URL        string `json:"url"`
	Streamable struct {
		Text      string `json:"#text"`
		Fulltrack string `json:"fulltrack"`
	} `json:"streamable"`
	Artist struct {
		Name string `json:"name"`
		MBID string `json:"mbid"`
		URL  string `json:"url"`
	} `json:"artist"`
	Image []struct {
		Text string `json:"#text"`
		Size string `json:"size"`
	} `json:"image"`
}

type TracksResponse struct {
	Tracks struct {
		Track []Track `json:"track"`
	} `json:"tracks"`
}

type TopTrackResponse struct {
	TopTracks struct {
		Track []Track `json:"track"`
	} `json:"tracks"`
}



func getTopTracks(apiKey string) (TracksResponse, error) {
	url := "https://ws.audioscrobbler.com/2.0/?method=chart.gettoptracks&api_key=" + apiKey + "&format=json"

	resp, err := http.Get(url)
	if err != nil {
		return TracksResponse{}, err
	}
	defer resp.Body.Close()

	var data TracksResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return TracksResponse{}, err
	}

	return data, nil
}

func TopTracksHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := "4b7791f101520fda28bd4f90ab58429e"

	data, err := getTopTracks(apiKey)
	if err != nil {
		http.Error(w, "Failed to fetch top tracks", http.StatusInternalServerError)
		log.Println("Failed to fetch top tracks:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func fetchTopTrack(region string) (*TopTrackResponse, error) {
	apiKey := "YOUR_API_KEY" // Replace with your Last.fm API key
	url := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&country=%s&api_key=%s&format=json", region, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var topTrackResponse TopTrackResponse
	if err := json.NewDecoder(resp.Body).Decode(&topTrackResponse); err != nil {
		return nil, err
	}

	return &topTrackResponse, nil
}
