package handler

import (
	"buda-challenge/configuration"
	"buda-challenge/processor"
	e "buda-challenge/error"
	"errors"
)

type Handler struct {
	Configuration configuration.Configuration
	Processor processor.Processor
}

func (handler Handler) HandleRequest() ([]string, error) {
	config, err := handler.Configuration.GetConfiguration()
	if err != nil {
		return nil, errors.New(e.ErrorReadingInput)
	}

	stations, err := handler.Configuration.GetTrainNetwork()
	if err != nil {
		return nil, errors.New(e.ErrorReadingFile)
	}

	stationsWithoutForks, forks := handler.Processor.GetStations(stations, config.TrainColor)

	routes := handler.Processor.GetRoutes(stationsWithoutForks, forks)

	handler.Processor.SortRoutes(routes, handler.Configuration.GetFinalStation())

	shortestRoute := handler.Processor.GetShortestRoute(routes, config.InitialStation, config.FinalStation)
	if shortestRoute == nil {
		return nil, errors.New(e.ErrorInvalidCombination)
	}

	route := handler.Processor.GetRoute(shortestRoute, config.InitialStation, config.FinalStation)

	return route, nil
}
