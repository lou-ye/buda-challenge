package configuration

import (
	"buda-challenge/dto"
	"buda-challenge/reader"
)

const (
	StationA             = "A"
	StationB             = "B"
	StationC             = "C"
	StationD             = "D"
	StationE             = "E"
	StationF             = "F"
	StationG             = "G"
	StationH             = "H"
	StationI             = "I"
	TrainRed             = "RED"
	TrainGreen           = "GREEN"
	TrainWithoutColour   = "WITHOUT COLOR"
	trainNetworkFilePath = "configuration/train_network.json"
)

var stations = []string{StationA, StationB, StationC, StationD, StationE, StationF, StationG, StationH, StationI}
var colors = []string{TrainRed, TrainGreen, TrainWithoutColour}

type Configuration interface {
	GetConfiguration() (dto.Configuration, error)
	GetTrainNetwork() ([]dto.Station, error)
	GetFinalStation() string
	GetTrainWithoutColor() string
}

type ConfigurationImpl struct {
	Reader reader.Reader
}

func(c ConfigurationImpl) GetConfiguration() (dto.Configuration, error) {
	config, err := c.Reader.ReadInput(stations, colors)
	if err != nil {
		return dto.Configuration{}, err
	}

	return config, nil
}

func(c ConfigurationImpl) GetTrainNetwork() ([]dto.Station, error) {
	stations, err := c.Reader.ReadFile(trainNetworkFilePath)
	if err != nil {
		return nil, err
	}

	return stations, nil
}

func(c ConfigurationImpl) GetFinalStation() string {
	return StationF
}

func(c ConfigurationImpl) GetTrainWithoutColor() string {
	return TrainWithoutColour
}