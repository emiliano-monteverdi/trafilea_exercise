package numberrepo

import (
	"github.com/Trafilea/internal/core/domain"
	"github.com/Trafilea/pkg/errors"
	"net/http"
)

func (repo *repository) Save(number domain.Number) {
	repo.storageClient[number.Value] = number.Type
}

func (repo *repository) Find(number int) (string, errors.Error) {
	value, ok := repo.storageClient[number]
	if !ok {
		return "", errors.New(http.StatusNotFound, "the number couldn't be found in the local storage")
	}

	return value, nil
}

func (repo *repository) FindAll() []domain.Number {
	numbersArray := make([]domain.Number, 0)
	for key, value := range repo.storageClient {
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
