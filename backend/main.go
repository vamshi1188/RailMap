package main

import (
	"RailMap/api"
	"net/http"
)

func main() {

	// var TrainData pkg.TrainDetails

	// TrainData.AddTrain("krishna", 17405)
	// TrainData.AddTrain("golconda", 28099)
	// TrainData.GetTrain()

	api.LoadAPI()
	http.HandleFunc("/station", api.GetStation)
	http.HandleFunc("/train", api.GetTrainLiveLocation)
	http.HandleFunc("/livemap", api.GetLATandLONG)
	http.HandleFunc("/key", api.HandleGetKey)
	http.Handle("/", http.FileServer(http.Dir("../frontend")))

	http.ListenAndServe(":8080", nil)

	// api.GetLATandLONG()

}
