package mocks

import (
	"fmt"
	"github.com/Trafilea/internal/handlers/numberhdl"
)

func ResponseCreate() numberhdl.CreateResponse {
	number := Number()
	return numberhdl.CreateResponse{Message: fmt.Sprintf("the number %v has been stored successfully", number.Value)}
}

func ResponseGet() numberhdl.GetResponse {
	number := Number()
	return numberhdl.GetResponse{
		Value: number.Type,
	}
}

func ResponseBulkGetValue() numberhdl.BulkGetValueResponse {
	numbers := NumbersArrayAsString("int")
	return numberhdl.BulkGetValueResponse{Value: numbers}
}

func ResponseBulkGetType() numberhdl.BulkGetTypeResponse {
	numbers := NumbersArrayAsString("type")
	return numberhdl.BulkGetTypeResponse{Type: numbers}
}
