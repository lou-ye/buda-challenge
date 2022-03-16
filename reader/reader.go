package reader

import (
	"buda-challenge/dto"
	"buda-challenge/validator"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Reader interface {
	ReadInput(stations, colors []string) (dto.Configuration, error)
	ReadFile(fileName string) ([]dto.Station, error)
	Read(requiredValue string, validValues []string) (string, error)
}

type ReaderImpl struct {
	Validator validator.Validator
	Mocked    func() (string, error)
}

func(r ReaderImpl) ReadInput(stations, colors []string) (dto.Configuration, error) {
	initialStation, err := r.Read("initial station", stations)
	if err != nil {
		return dto.Configuration{}, err
	}

	finalStation, err := r.Read("final station", stations)
	if err != nil {
		return dto.Configuration{}, err
	}

	trainColor, err := r.Read("train color", colors)
	if err != nil {
		return dto.Configuration{}, err
	}

	return dto.Configuration{
		InitialStation: initialStation,
		FinalStation:   finalStation,
		TrainColor:     trainColor,
	}, nil
}

func(r ReaderImpl) Read(requiredValue string, validValues []string) (string, error) {
	if r.Mocked != nil {
		return "", errors.New("mocked to test")
	}

	reader := bufio.NewReader(os.Stdin)
	var enteredValue string

	for {
		fmt.Print("Enter " + requiredValue + " [ Valid values: " + strings.Join(validValues, " - ") + " ] : ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		enteredValue = strings.ToUpper(strings.TrimSuffix(input, "\n"))
		if r.Validator.Validate(enteredValue, validValues) {
			break
		}
		fmt.Println("Invalid value! Try again!")
	}

	return enteredValue, nil
}

func(r ReaderImpl) ReadFile(fileName string) ([]dto.Station, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []dto.Station{}, err
	}

	var stations []dto.Station
	err = json.Unmarshal(content, &stations)
	if err != nil {
		return []dto.Station{}, err
	}

	return stations, nil
}
