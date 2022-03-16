package processor

import (
	"buda-challenge/dto"
	"buda-challenge/validator"
	"math"
	"sort"
)

type Processor interface {
	GetStations(stations []dto.Station, trainColor string) ([]string, [][]string)
	GetRoutes(stationsWithoutForks []string, forks [][]string) [][]string
	SortRoutes(routes [][]string, lastStation string)
	GetRoute(routes []string, initialStation, lastStation string) []string
	GetShortestRoute(routes [][]string, initialStation, lastStation string) []string
}

type ProcessorImpl struct {
	Validator validator.Validator
}

func(p ProcessorImpl) GetStations(stations []dto.Station, trainColor string) ([]string, [][]string) {
	var stationsWithoutForks []string
	var forks [][]string

	for _, station := range stations {
		if validateTrainColor(trainColor, station) {
			for _, v := range station.Forks {
				forks = append(forks, GetForkNames(v, trainColor))
			}
			stationsWithoutForks = append(stationsWithoutForks, station.Name)
		}
	}

	return stationsWithoutForks, forks
}

func GetForkNames(forks []dto.Station, trainColor string) []string {
	var forksNames []string

	for _, fork := range forks {
		if validateTrainColor(trainColor, fork) {
			forksNames = append(forksNames, fork.Name)
		}
	}

	return forksNames
}

func validateTrainColor(trainColor string, fork dto.Station) bool {
	return trainColor == "WITHOUT COLOR" || fork.TrainColor == "WITHOUT COLOR" || fork.TrainColor == trainColor
}

func(p ProcessorImpl) GetRoutes(stationsWithoutForks []string, forks [][]string) [][]string {
	var routes [][]string
	var route []string

	for _, v := range forks {
		route = append(stationsWithoutForks, v...)
		routes = append(routes, route)
	}

	return routes
}

func(p ProcessorImpl) SortRoutes(routes [][]string, lastStation string) {
	for _, route := range routes {
		sortRoute(route, lastStation)
	}
}

func sortRoute(route []string, lastStation string) []string {
	r := sort.SearchStrings(route, lastStation)
	return append(append(route[:r], route[r+1:]...), lastStation)
}

func(p ProcessorImpl) GetRoute(route []string, initialStation string, lastStation string) []string {
	initialStationPosition := GetPosition(route, initialStation)
	finalStationPosition := GetPosition(route, lastStation)

	if initialStationPosition <= finalStationPosition {
		return getRoute(route, initialStationPosition, finalStationPosition)
	} else {
		return reverse(getRoute(route, finalStationPosition, initialStationPosition))
	}
}

func getRoute(route []string, initial int, final int) []string {
	return append(route[initial : final+1])
}

func reverse(input []string) []string {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}

	return input
}

func(p ProcessorImpl) GetShortestRoute(routes [][]string, initialStation, lastStation string) []string {
	var shortestDistance int
	var shortestRoute []string

	for _, route := range routes {
		if p.Validator.Validate(lastStation, route) {
			initialStation := GetPosition(route, initialStation)
			finalStation := GetPosition(route, lastStation)

			currentDistance := CalculateDistance(finalStation, initialStation)
			shortestDistance = GetShorterDistance(shortestDistance, currentDistance)

			if IsShortestRoute(shortestDistance, currentDistance) {
				shortestRoute = route
			}
		}
	}

	return shortestRoute
}

func GetShorterDistance(shortestDistance int, distance int) int {
	if shortestDistance == 0 {
		shortestDistance = distance
	}
	return shortestDistance
}

func IsShortestRoute(lastDistance int, distance int) bool {
	if distance <= lastDistance {
		return true
	}
	return false
}

func CalculateDistance(finalStation int, initialStation int) int {
	distance := finalStation - initialStation
	return int(math.Abs(-float64(distance)))
}

func GetPosition(stations []string, station string) int {
	var position int
	for i, v := range stations {
		if v == station {
			position = i
		}
	}
	return position
}
