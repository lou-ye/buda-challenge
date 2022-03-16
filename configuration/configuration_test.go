package configuration

import (
	"buda-challenge/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	e "buda-challenge/error"
	"testing"
)

const (
	readInputMethodName = "ReadInput"
	readFileMethodName = "ReadFile"
)

func Test_WhenInputsCanBeReadCorrectly_ReturnsValidConfiguration(t *testing.T) {
	mockReader := new(MockReader)

	mockReader.On(readInputMethodName, mock.Anything, mock.Anything).Return(getConfiguration(StationA, StationF, TrainRed), nil)

	config := ConfigurationImpl{
		Reader: mockReader,
	}

	result, err := config.GetConfiguration()

	resultExpected := dto.Configuration{
		InitialStation: StationA,
		FinalStation:   StationF,
		TrainColor:     TrainRed,
	}

	assert.Equal(t, resultExpected, result)
	assert.Nil(t, err)
}

func Test_WhenInputsCanNotBeReadCorrectly_ReturnsError(t *testing.T) {
	mockReader := new(MockReader)

	mockReader.On(readInputMethodName, mock.Anything, mock.Anything).Return(nil, errors.New(e.ErrorReadingInput))

	config := ConfigurationImpl{
		Reader: mockReader,
	}

	_, err := config.GetConfiguration()

	assert.NotNil(t, err)
	assert.Equal(t, e.ErrorReadingInput, err.Error())
}

func Test_WhenFileCanBeReadCorrectly_ReturnsValidTrainNetwork(t *testing.T) {
	mockReader := new(MockReader)

	mockReader.On(readInputMethodName, mock.Anything, mock.Anything).Return(getConfiguration(StationA, StationF, TrainRed), nil)
	mockReader.On(readFileMethodName, mock.Anything).Return(getStations())

	config := ConfigurationImpl{
		Reader: mockReader,
	}

	result, err := config.GetTrainNetwork()

	resultExpected := getStations()

	assert.Equal(t, resultExpected, result)
	assert.Nil(t, err)
}

func Test_WhenFileCanNotBeReadCorrectly_ReturnsError(t *testing.T) {
	mockReader := new(MockReader)

	mockReader.On(readInputMethodName, mock.Anything, mock.Anything).Return(getConfiguration(StationA, StationF, TrainRed), nil)
	mockReader.On(readFileMethodName, mock.Anything).Return(nil, errors.New(e.ErrorReadingFile))

	config := ConfigurationImpl{
		Reader: mockReader,
	}

	_, err := config.GetTrainNetwork()

	assert.NotNil(t, err)
	assert.Equal(t, e.ErrorReadingFile, err.Error())
}

func Test_ReturnValidTrainWithoutColor(t *testing.T) {
	result := ConfigurationImpl{}.GetTrainWithoutColor()

	assert.Equal(t, TrainWithoutColour, result)
}

func Test_ReturnValidFinalStation(t *testing.T) {
	result := ConfigurationImpl{}.GetFinalStation()

	assert.Equal(t, StationF, result)
}

type MockReader struct { mock.Mock }

func (s *MockReader) ReadInput(stations, colors []string) (dto.Configuration, error) {
	args := s.Called(stations, colors)

	if args.Get(0) == nil {
		return dto.Configuration{}, args.Error(1)
	}

	return args.Get(0).(dto.Configuration), nil
}

func (s *MockReader) ReadFile(fileName string) ([]dto.Station, error) {
	args := s.Called(fileName)

	if args.Get(0) == nil {
		return []dto.Station{}, args.Error(1)
	}

	return args.Get(0).([]dto.Station), nil
}

func (s *MockReader) Read(requiredValue string, validValues []string) (string, error) {
	args := s.Called(requiredValue, validValues)

	if args.Get(0) == nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), nil
}

func getConfiguration(initialStation, finalStation, trainColor string) dto.Configuration {
	return dto.Configuration{
		InitialStation: initialStation,
		FinalStation:   finalStation,
		TrainColor:     trainColor,
	}
}

func getStations() []dto.Station {
	stationA := dto.Station{Name: StationA, Forks: nil, TrainColor: TrainWithoutColour}
	stationB := dto.Station{Name: StationB, Forks: nil, TrainColor: TrainWithoutColour}
	stationC := dto.Station{Name: StationC, Forks: [][]dto.Station{
		{
			{
				Name:       StationD,
				Forks:      nil,
				TrainColor: TrainWithoutColour,
			},
			{
				Name:       StationE,
				Forks:      nil,
				TrainColor: TrainWithoutColour,
			},
		},
		{
			{
				Name:       StationG,
				Forks:      nil,
				TrainColor: TrainGreen,
			},
			{
				Name:       StationH,
				Forks:      nil,
				TrainColor: TrainRed,
			},
			{
				Name:       StationI,
				Forks:      nil,
				TrainColor: TrainGreen,
			},
		},
	},
		TrainColor: TrainWithoutColour,
	}
	stationF := dto.Station{Name: StationF, Forks: nil, TrainColor: TrainWithoutColour}

	return []dto.Station{stationA, stationB, stationC, stationF}
}