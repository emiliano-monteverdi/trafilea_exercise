package mocks

import (
	"errors"
	"github.com/Trafilea/internal/core/domain"
	"strconv"
)

var (
	localStorage = map[int]string{
		1:  "1",
		2:  "2",
		3:  "Type 1",
		4:  "4",
		5:  "Type 2",
		6:  "Type 1",
		7:  "7",
		8:  "8",
		9:  "Type 1",
		10: "Type 2",
		11: "11",
		12: "Type 1",
		13: "13",
		14: "14",
		15: "Type 3",
	}
)

func Number() domain.Number {
	return domain.Number{
		Value: 1,
		Type:  "1",
	}
}

func NumbersArray() []domain.Number {
	numbersArray := make([]domain.Number, 0)
	for key, value := range localStorage {
		if key == 0 {
			continue
		}

		number := domain.Number{}
		number.Value = key
		number.Type = value

		numbersArray = append(numbersArray, number)
	}

	return numbersArray
}

func NumbersArrayAsString(option string) []string {
	numbersArray := make([]string, 0)
	for key, value := range localStorage {
		if key == 0 {
			continue
		}

		switch option {
		case "int":
			numbersArray = append(numbersArray, strconv.Itoa(key))

		case "type":
			numbersArray = append(numbersArray, value)

		default:
			panic(errors.New("unsupported option"))
		}
	}

	return numbersArray
}
