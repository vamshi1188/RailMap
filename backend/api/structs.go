package api

import ()

type LiveTrainData struct {
	TrainNumber string  `json:"train_number"`
	Latitude    float64 `json:"current_lat"`
	Longitude   float64 `json:"current_lng"`
}

type LiveMapResponse struct {
	Data []LiveTrainData `json:"data"`
}
