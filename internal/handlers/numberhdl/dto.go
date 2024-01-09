package numberhdl

import (
	"encoding/json"
	"github.com/Trafilea/pkg/errors"
	"github.com/Trafilea/pkg/logs"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type CreateRequest struct {
	Number int `json:"number"`
}

type CreateResponse struct {
	Message string `json:"message"`
}

type getRequest struct {
	Number int
}

type GetResponse struct {
	Value string `json:"value"`
}

type BulkGetValueResponse struct {
	Value []string `json:"value"`
}

type BulkGetTypeResponse struct {
	Type []string `json:"type"`
}

type dto struct {
	request *http.Request
}

func buildDTO(r *http.Request) dto {
	return dto{request: r}
}

func (d dto) create() (CreateRequest, errors.Error) {
	dto := CreateRequest{}

	if err := json.NewDecoder(d.request.Body).Decode(&dto); err != nil {
		logs.Info("BuildDTO - unmarshal error")
		return CreateRequest{}, errors.New(http.StatusBadRequest, "unmarshal error")
	}

	return dto, nil
}

func (d dto) get() (getRequest, errors.Error) {
	dto := getRequest{}

	number := chi.URLParam(d.request, "id")

	numberAsInt, err := strconv.Atoi(number)
	if err != nil {
		logs.Info("BuildDTO - invalid number error")
		return getRequest{}, errors.New(http.StatusBadRequest, "invalid number error")
	}

	dto.Number = numberAsInt

	return dto, nil
}
