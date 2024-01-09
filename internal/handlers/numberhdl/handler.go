package numberhdl

import (
	"fmt"
	"github.com/Trafilea/internal/core/ports"
	"github.com/Trafilea/pkg/logs"
	"github.com/go-chi/render"
	"net/http"
)

type Handler struct {
	numberService ports.NumberService
}

func New(numberService ports.NumberService) *Handler {
	return &Handler{numberService: numberService}
}

func (hdl *Handler) Create(w http.ResponseWriter, r *http.Request) {
	dto, err := buildDTO(r).create()
	if err != nil {
		logs.Error(err)
		http.Error(w, err.Error(), err.StatusCode())
		return
	}

	hdl.numberService.Create(dto.Number)

	resp := CreateResponse{Message: fmt.Sprintf("the number %v has been stored successfully", dto.Number)}

	render.JSON(w, r, resp)
}

func (hdl *Handler) Get(w http.ResponseWriter, r *http.Request) {
	dto, err := buildDTO(r).get()
	if err != nil {
		logs.Error(err)
		http.Error(w, err.Error(), err.StatusCode())
		return
	}

	number, err := hdl.numberService.Get(dto.Number)
	if err != nil {
		logs.Error(err)
		http.Error(w, err.Error(), err.StatusCode())
		return
	}

	resp := GetResponse{Value: number}

	render.JSON(w, r, resp)
}

func (hdl *Handler) BulkGetValue(w http.ResponseWriter, r *http.Request) {
	numbersArray, err := hdl.numberService.BulkGet()
	if err != nil {
		logs.Error(err)
		http.Error(w, err.Error(), err.StatusCode())
		return
	}

	resp := BulkGetValueResponse{Value: numbersArray}

	render.JSON(w, r, resp)
}

func (hdl *Handler) BulkGetType(w http.ResponseWriter, r *http.Request) {
	numbersArray, err := hdl.numberService.BulkGetTypes()
	if err != nil {
		logs.Error(err)
		http.Error(w, err.Error(), err.StatusCode())
		return
	}

	resp := BulkGetTypeResponse{Type: numbersArray}

	render.JSON(w, r, resp)
}
