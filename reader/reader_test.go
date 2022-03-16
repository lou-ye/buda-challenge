package reader

import (
	"buda-challenge/dto"
	"buda-challenge/validator"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	stationA             = "A"
	stationB             = "B"
	stationC             = "C"
	stationD             = "D"
	stationE             = "E"
	stationF             = "F"
	stationG             = "G"
	stationH             = "H"
	stationI             = "I"
	trainRed             = "RED"
	trainGreen           = "GREEN"
	trainWithoutColour   = "WITHOUT COLOR"
	trainNetworkFileValidPath = "../configuration/train_network.json"
	trainNetworkFileInvalidPath = "../../configuration/train_network.json"
)

func Test_GivenAValidTrainNetworkFilePath_ReturnStations(t *testing.T) {
	result, err := ReaderImpl{}.ReadFile(trainNetworkFileValidPath)

	assert.Equal(t, getStations(), result)
	assert.Nil(t, err)
}

func Test_GivenAInvalidTrainNetworkFilePath_ReturnError(t *testing.T) {
	_, err := ReaderImpl{}.ReadFile(trainNetworkFileInvalidPath)

	assert.NotNil(t, err)
}

func Test_WhenInputsCanNotBeReadCorrectly_ReturnError(t *testing.T) {
	reader := ReaderImpl{Validator: validator.ValidatorImpl{}, Mocked: func() (string, error) {
		return "", errors.New("mocked to test")
	}}

	_, err := reader.ReadInput([]string{stationA, stationB, stationC}, []string{trainRed, trainGreen, trainWithoutColour})

	assert.NotNil(t, err)
}

func getStations() []dto.Station {
	stationA := dto.Station{Name: stationA, Forks: nil, TrainColor: trainWithoutColour}
	stationB := dto.Station{Name: stationB, Forks: nil, TrainColor: trainWithoutColour}
	stationC := dto.Station{Name: stationC, Forks: [][]dto.Station{
		{
			{
				Name:       stationD,
				Forks:      nil,
				TrainColor: trainWithoutColour,
			},
			{
				Name:       stationE,
				Forks:      nil,
				TrainColor: trainWithoutColour,
			},
		},
		{
			{
				Name:       stationG,
				Forks:      nil,
				TrainColor: trainGreen,
			},
			{
				Name:       stationH,
				Forks:      nil,
				TrainColor: trainRed,
			},
			{
				Name:       stationI,
				Forks:      nil,
				TrainColor: trainGreen,
			},
		},
	},
		TrainColor: trainWithoutColour,
	}
	stationF := dto.Station{Name: stationF, Forks: nil, TrainColor: trainWithoutColour}

	return []dto.Station{stationA, stationB, stationC, stationF}
}
