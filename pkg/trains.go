package pkg

import "fmt"

type Train struct {
	trainName string
	trainNum  int
}
type TrainDetails struct {
	Trains []Train
}

func (trn *TrainDetails) AddTrain(name string, num int) {

	trn.Trains = append(trn.Trains, Train{trainName: name, trainNum: num})
}

func (trn TrainDetails) GetTrain() {
	fmt.Println(trn.Trains)
}
