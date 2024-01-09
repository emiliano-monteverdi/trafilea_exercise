package ports

import (
	"github.com/Trafilea/internal/core/domain"
	"github.com/Trafilea/pkg/errors"
)

type NumberRepository interface {
	Save(number domain.Number)
	Find(number int) (string, errors.Error)
	FindAll() []domain.Number
}
