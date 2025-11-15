package main

import (
	"RailMap/backend/api"
	"net/http"
	// "RailMap/pkg"
)

func main() {

	// var TrainData pkg.TrainDetails

	// TrainData.AddTrain("krishna", 17405)
	// TrainData.AddTrain("golconda", 28099)
	// TrainData.GetTrain()

	http.HandleFunc("/station", api.GetStation)
	http.HandleFunc("/train", api.GetTrainLiveLocation)
	http.ListenAndServe(":8080", nil)
}
