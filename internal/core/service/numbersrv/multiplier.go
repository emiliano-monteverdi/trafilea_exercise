package numbersrv

import (
	"strconv"
)

const (
	multiplierThree = 3
	multiplierFive  = 5
)

var (
	multiplierMapTypes = map[bool]map[bool]string{
		true: {
			true:  "Type 3",
			false: "Type 1",
		},
		false: {
			true: "Type 2",
		},
	}
)

func getTypeIfIsAnMultiplier(number int) string {
	var (
		result string
		found  bool
	)

	if result, found = multiplierMapTypes[number%multiplierThree == 0][number%multiplierFive == 0]; !found {
		result = strconv.Itoa(number)
	}

	return result
}
