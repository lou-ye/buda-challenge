package handler

import (
	"buda-challenge/configuration"
	"buda-challenge/dto"
	"buda-challenge/processor"
	"buda-challenge/validator"
	e "buda-challenge/error"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

const (
	readInputMethodName = "ReadInput"
	readFileMethodName = "ReadFile"
)

func Test_WhenInitialStationIsFFinalStationIsBAndTrainColorIsGreen_ReturnStationFIGCB(t *testing.T) {
	mockReader := new(MockReader)

	mockReader.On(readInputMethodName, mock.Anything, mock.Anything).Return(getConfiguration(configuration.StationF, configuration.StationB, configuration.TrainGreen), nil)
	mockReader.On(readFileMethodName, mock.Anything).Return(getTrainNetwork(), nil)

	handler := Handler{
		Configuration: configuration.ConfigurationImpl{
			Reader: mockReader,
		},
		Processor: processor.ProcessorImpl{
			Validator: validator.ValidatorImpl{},
		},
	}

	result, err := handler.HandleRequest()

	resultExpected := []string{configuration.StationF, configuration.StationI, configuration.StationG, configuration.StationC, configuration.StationB}

	assert.Equal(t, resultExpected, result)
	assert.Nil(t, err)
}

func Test_WhenInitialStationIsFFinalStationIsDAndTrainWithOutColor_ReturnStationFED(t *testing.T) {
	mockReader := new(MockReader)

	mockReader.On(readInputMethodName, mock.Anything, mock.Anything).Return(getConfiguration(configuration.StationF, configuration.StationD, configuration.TrainWithoutColour), nil)
	mockReader.On(readFileMethodName, mock.Anything).Return(getTrainNetwork())

	handler := Handler{
		Configuration: configuration.ConfigurationImpl{
			Reader: mockReader,
		},
		Processor: processor.ProcessorImpl{
			Validator: validator.ValidatorImpl{},
		},
	}

	result, err := handler.HandleRequest()

	resultExpected := []string{configuration.StationF, configuration.StationE, configuration.StationD}

	assert.Equal(t, resultExpected, result)
	assert.Nil(t, err)
}

func Test_WhenInitialStationIsAFinalStationIsFAndTrainWithOutColor_ReturnStationABCDEF(t *testing.T) {
	mockReader := new(MockReader)

	mockReader.On(readInputMethodName, mock.Anything, mock.Anything).Return(getConfiguration(configuration.StationA, configuration.StationF, configuration.TrainWithoutColour), nil)
	mockReader.On(readFileMethodName, mock.Anything).Return(getTrainNetwork())

	handler := Handler{
		Configuration: configuration.ConfigurationImpl{
			Reader: mockReader,
		},
		Processor: processor.ProcessorImpl{
			Validator: validator.ValidatorImpl{},
		},
	}

	result, err := handler.HandleRequest()

	resultExpected := []string{configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationD, configuration.StationE, configuration.StationF}

	assert.Equal(t, resultExpected, result)
	assert.Nil(t, err)
}

func Test_WhenInitialStationIsAFinalStationIsFAndTrainColorIsRed_ReturnStationABCHF(t *testing.T) {
	mockReader := new(MockReader)

	mockReader.On(readInputMethodName, mock.Anything, mock.Anything).Return(getConfiguration(configuration.StationA, configuration.StationF, configuration.TrainRed), nil)
	mockReader.On(readFileMethodName, mock.Anything).Return(getTrainNetwork())

	handler := Handler{
		Configuration: configuration.ConfigurationImpl{
			Reader: mockReader,
		},
		Processor: processor.ProcessorImpl{
			Validator: validator.ValidatorImpl{},
		},
	}

	result, err := handler.HandleRequest()

	resultExpected := []string{configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationH, configuration.StationF}

	assert.Equal(t, resultExpected, result)
	assert.Nil(t, err)
}

func Test_WhenInitialStationIsAFinalStationIsFAndTrainColorIsGreen_ReturnStationABCGIF(t *testing.T) {
	mockReader := new(MockReader)

	mockReader.On(readInputMethodName, mock.Anything, mock.Anything).Return(getConfiguration(configuration.StationA, configuration.StationF, configuration.TrainGreen), nil)
	mockReader.On(readFileMethodName, mock.Anything).Return(getTrainNetwork())

	handler := Handler{
		Configuration: configuration.ConfigurationImpl{
			Reader: mockReader,
		},
		Processor: processor.ProcessorImpl{
			Validator: validator.ValidatorImpl{},
		},
	}

	result, err := handler.HandleRequest()

	resultExpected := []string{configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationG, configuration.StationI, configuration.StationF}

	assert.Equal(t, resultExpected, result)
	assert.Nil(t, err)
}

func Test_WhenInitialStationIsBFinalStationIsDAndTrainColorIsRed_ReturnStationABCDEF(t *testing.T) {
	mockReader := new(MockReader)

	mockReader.On(readInputMethodName, mock.Anything, mock.Anything).Return(getConfiguration(configuration.StationB, configuration.StationD, configuration.TrainRed), nil)
	mockReader.On(readFileMethodName, mock.Anything).Return(getTrainNetwork())

	handler := Handler{
		Configuration: configuration.ConfigurationImpl{
			Reader: mockReader,
		},
		Processor: processor.ProcessorImpl{
			Validator: validator.ValidatorImpl{},
		},
	}

	result, err := handler.HandleRequest()

	resultExpected := []string{configuration.StationB, configuration.StationC, configuration.StationD}

	assert.Equal(t, resultExpected, result)
	assert.Nil(t, err)
}

func Test_WhenInputCanNotBeRead_ReturnsError(t *testing.T) {
	mockReader := new(MockReader)

	mockReader.On(readInputMethodName, mock.Anything, mock.Anything).Return(nil, errors.New(e.ErrorReadingInput))

	handler := Handler{
		Configuration: configuration.ConfigurationImpl{
			Reader: mockReader,
		},
		Processor: processor.ProcessorImpl{
			Validator: validator.ValidatorImpl{},
		},
	}

	result, err := handler.HandleRequest()

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, e.ErrorReadingInput, err.Error())
}

func Test_WhenFileCanNotBeRead_ReturnsError(t *testing.T) {
	mockReader := new(MockReader)

	mockReader.On(readInputMethodName, mock.Anything, mock.Anything).Return(getConfiguration(configuration.StationF, configuration.StationB, configuration.TrainGreen), nil)
	mockReader.On(readFileMethodName, mock.Anything, mock.Anything).Return(nil, errors.New(e.ErrorReadingFile))

	handler := Handler{
		Configuration: configuration.ConfigurationImpl{
			Reader: mockReader,
		},
		Processor: processor.ProcessorImpl{
			Validator: validator.ValidatorImpl{},
		},
	}

	result, err := handler.HandleRequest()

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, e.ErrorReadingFile, err.Error())
}

func Test_WhenInitialStationIsAFinalStationIsIAndTrainColorIsRed_ReturnError(t *testing.T) {
	mockReader := new(MockReader)

	mockReader.On(readInputMethodName, mock.Anything, mock.Anything).Return(getConfiguration(configuration.StationA, configuration.StationI, configuration.TrainRed), nil)
	mockReader.On(readFileMethodName, mock.Anything).Return(getTrainNetwork(), nil)

	handler := Handler{
		Configuration: configuration.ConfigurationImpl{
			Reader: mockReader,
		},
		Processor: processor.ProcessorImpl{
			Validator: validator.ValidatorImpl{},
		},
	}

	_, err := handler.HandleRequest()

	assert.NotNil(t, err)
	assert.Equal(t, e.ErrorInvalidCombination, err.Error())
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

func getTrainNetwork() []dto.Station {
	stationA := dto.Station{Name: configuration.StationA, Forks: nil, TrainColor: configuration.TrainWithoutColour}
	stationB := dto.Station{Name: configuration.StationB, Forks: nil, TrainColor: configuration.TrainWithoutColour}
	stationC := dto.Station{Name: configuration.StationC, Forks: [][]dto.Station{
		{
			{
				Name:          configuration.StationD,
				Forks: nil,
				TrainColor: configuration.TrainWithoutColour,
			},
			{
				Name:          configuration.StationE,
				Forks: nil,
				TrainColor: configuration.TrainWithoutColour,
			},
		},
		{
			{
				Name:          configuration.StationG,
				Forks: nil,
				TrainColor: configuration.TrainGreen,
			},
			{
				Name:          configuration.StationH,
				Forks: nil,
				TrainColor: configuration.TrainRed,
			},
			{
				Name:          configuration.StationI,
				Forks: nil,
				TrainColor: configuration.TrainGreen,
			},
		},
	},
		TrainColor: configuration.TrainWithoutColour,
	}
	stationF := dto.Station{Name: configuration.StationF, Forks: nil, TrainColor: configuration.TrainWithoutColour}

	return []dto.Station{stationA, stationB, stationC, stationF}
}






