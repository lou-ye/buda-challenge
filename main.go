package main

import (
	"buda-challenge/configuration"
	"buda-challenge/handler"
	"buda-challenge/processor"
	"buda-challenge/reader"
	"buda-challenge/validator"
	"fmt"
)

func main() {
	result, err := handler.Handler{
		Configuration: configuration.ConfigurationImpl{
			Reader: reader.ReaderImpl{
				Validator: validator.ValidatorImpl{},
			},
		},
		Processor: processor.ProcessorImpl{
			Validator: validator.ValidatorImpl{},
		},
	}.HandleRequest()

	fmt.Println("Shortest route: ", result, err)
}