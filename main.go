package main

import "RailMap/pkg"

func main() {

	var TrainData pkg.TrainDetails

	TrainData.AddTrain("krishna", 17405)
	TrainData.AddTrain("golconda", 28099)
	TrainData.GetTrain()
}
