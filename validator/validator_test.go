package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	stationA = "A"
	stationB = "B"
	stationC = "C"
	trainRed = "RED"
	trainGreen = "GREEN"
	trainWithoutColour = "WITHOUT COLOR"
	invalidTrainColor = "BLUE"
)

func Test_WhenValueIsWithinValues_ReturnTrue(t *testing.T) {
	value := stationB
	values := []string{stationA, stationB, stationC}

	result := ValidatorImpl{}.Validate(value, values)

	assert.True(t, result)
}

func Test_WhenValueIsNotWithinValues_ReturnFalse(t *testing.T) {
	value := invalidTrainColor
	values := []string{trainRed, trainGreen, trainWithoutColour}

	result := ValidatorImpl{}.Validate(value, values)

	assert.False(t, result)
}
