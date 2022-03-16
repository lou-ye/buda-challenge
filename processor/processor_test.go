package processor

import (
	"buda-challenge/dto"
	"buda-challenge/validator"
	"buda-challenge/configuration"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	lastStation = "F"
	initialStation = "A"
	endStation = "F"
)

func Test_GivenATrainNetworkAndATrainColorGreen_ReturnValidStationsAndForks(t *testing.T) {
	processor := ProcessorImpl{Validator: validator.ValidatorImpl{}}

	stationsWithoutForks, forks := processor.GetStations(getTrainNetwork(), configuration.TrainGreen)

	stationsWithoutForksExpected := []string{configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationF}
	forksExpectedForRedTrain := [][]string{{configuration.StationD, configuration.StationE}, {configuration.StationG, configuration.StationI}}

	assert.Equal(t, stationsWithoutForksExpected, stationsWithoutForks)
	assert.Equal(t, forksExpectedForRedTrain, forks)
}

func Test_GivenATrainNetworkAndATrainColorRed_ReturnValidStationsAndForks(t *testing.T) {
	processor := ProcessorImpl{Validator: validator.ValidatorImpl{}}

	stationsWithoutForks, forks := processor.GetStations(getTrainNetwork(), configuration.TrainRed)

	stationsWithoutForksExpected := []string{configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationF}
	forksExpectedForRedTrain := [][]string{{configuration.StationD, configuration.StationE}, {configuration.StationH}}

	assert.Equal(t, stationsWithoutForksExpected, stationsWithoutForks)
	assert.Equal(t, forksExpectedForRedTrain, forks)
}

func Test_GivenStationsAndForks_ReturnPossibleRoutes(t *testing.T) {
	processor := ProcessorImpl{Validator: validator.ValidatorImpl{}}

	stationsWithoutForks := []string{configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationF}
	forks := [][]string{{configuration.StationD, configuration.StationE}, {configuration.StationG, configuration.StationH, configuration.StationI}}

	routes := processor.GetRoutes(stationsWithoutForks, forks)

	routesExpected := [][]string{{configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationF, configuration.StationD, configuration.StationE}, {configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationF, configuration.StationG, configuration.StationH, configuration.StationI}}

	assert.Equal(t, routesExpected, routes)
}

func Test_GivenRoutes_ReturnsRoutesOrderedAccordingToTheirPosition(t *testing.T) {
	processor := ProcessorImpl{Validator: validator.ValidatorImpl{}}

	routes := [][]string{{configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationF, configuration.StationD, configuration.StationE}, {configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationF, configuration.StationG, configuration.StationH, configuration.StationI}}

	processor.SortRoutes(routes, lastStation)

	routesExpected := [][]string{{configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationD, configuration.StationE, configuration.StationF}, {configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationG, configuration.StationH, configuration.StationI, configuration.StationF}}

	assert.Equal(t, routesExpected, routes)
}


func Test_GivenPossibleRoutes_ReturnTheShortestRoute(t *testing.T) {
	processor := ProcessorImpl{Validator: validator.ValidatorImpl{}}

	routes := [][]string{{configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationD, configuration.StationE, configuration.StationF}, {configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationG, configuration.StationH, configuration.StationI, configuration.StationF}}

	route := processor.GetShortestRoute(routes, initialStation, endStation)

	routeExpected := []string{configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationD, configuration.StationE, configuration.StationF}

	assert.Equal(t, routeExpected, route)
}

func Test_GivenAStationsList_ReturnsOnlyThoseThatCanBeUsedAccordingToTheChosenColor(t *testing.T) {
	result := GetForkNames(getTrainNetwork(), configuration.TrainWithoutColour)
	assert.Equal(t, []string{configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationF}, result)
}

func Test_GivenTwoPositions_ReturnsDistanceBetweenThem(t *testing.T) {
	distance := CalculateDistance(6, 0)

	assert.Equal(t, 6, distance)
}

func Test_GivenAStationsList_ReturnsStationPositionIndicated(t *testing.T) {
	position := GetPosition([]string{configuration.StationA, configuration.StationB, configuration.StationC, configuration.StationF}, configuration.StationC)

	assert.Equal(t, 2, position)
}

func Test_GiveTwoDistances_ReturnIfShorter(t *testing.T) {
	isShortestRoute := IsShortestRoute(5, 3)

	assert.True(t, isShortestRoute)
}

func Test_GiveTwoDistances_ReturnTheShortest(t *testing.T) {
	shortestRoute := GetShorterDistance(3, 2)

	assert.Equal(t, 3, shortestRoute)
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
