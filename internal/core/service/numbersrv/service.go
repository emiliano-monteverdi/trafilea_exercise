package numbersrv

import (
	"github.com/Trafilea/internal/core/domain"
	"github.com/Trafilea/internal/core/ports"
	"github.com/Trafilea/pkg/errors"
	"github.com/Trafilea/pkg/logs"
	"net/http"
	"strconv"
)

type service struct {
	numberRepository ports.NumberRepository
}

func New(cardRepository ports.NumberRepository) ports.NumberService {
	return &service{numberRepository: cardRepository}
}

func (srv *service) Create(number int) {
	_number := domain.Number{}
	_number.Value = number
	_number.Type = getTypeIfIsAnMultiplier(number)

	srv.numberRepository.Save(_number)
}

func (srv *service) Get(number int) (string, errors.Error) {
	numberValue, err := srv.numberRepository.Find(number)
	if err != nil {
		logs.Info(err.Error())
		return "", errors.New(err.StatusCode(), "getting the number value from the NumberRepository has failed", errors.NewCauseError(err))
	}

	return numberValue, nil
}

func (srv *service) BulkGet() ([]string, errors.Error) {
	numbersArray := srv.numberRepository.FindAll()

	numbersAsIntArray := make([]string, 0)
	for _, number := range numbersArray {
		numbersAsIntArray = append(numbersAsIntArray, strconv.Itoa(number.Value))
	}

	if len(numbersAsIntArray) == 0 {
		return nil, errors.New(http.StatusNotFound, "there's nothing stored in the NumberRepository")
	}

	return numbersAsIntArray, nil
}

// I added this method because it raised doubts about one of the requirements in the statement. Therefore, I decided
// to include it anyway in case this was indeed what was being asked
func (srv *service) BulkGetTypes() ([]string, errors.Error) {
	numbersArray := srv.numberRepository.FindAll()

	numbersTypeArray := make([]string, 0)
	for _, number := range numbersArray {
		numbersTypeArray = append(numbersTypeArray, number.Type)
	}

	if len(numbersTypeArray) == 0 {
		return nil, errors.New(http.StatusNotFound, "there's nothing stored in the NumberRepository")
	}

	return numbersTypeArray, nil
}
