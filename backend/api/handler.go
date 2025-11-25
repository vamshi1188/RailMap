package api

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "os"
)

func GetStation(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("q")

	if query == "" {
		http.Error(w, "parameter empty", http.StatusBadRequest)
		return
	}

	url := "https://railradar.in/api/v1/search/stations?q=" + query + "&provider=railradar"

	resp, err := http.Get(url)

	if err != nil {
		http.Error(w, "invalid url", http.StatusBadGateway)
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)

}
func GetTrainLiveLocation(w http.ResponseWriter, r *http.Request) {

	trainnum := r.URL.Query().Get("trainnum")

	if trainnum == "" {
		http.Error(w, "Invalid train number!", http.StatusBadRequest)
		return
	}

	url := "https://railradar.in/api/v1/trains/" + trainnum

	resp, err := http.Get(url)

	if err != nil {
		http.Error(w, "invalid url", http.StatusBadGateway)
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}

// train live cordinates

func GetLATandLONG(w http.ResponseWriter, r *http.Request) {
	url := "https://railradar.in/api/v1/trains/live-map"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var result LiveMapResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(result)
}
