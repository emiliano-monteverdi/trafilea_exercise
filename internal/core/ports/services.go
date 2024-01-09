package ports

import (
	"github.com/Trafilea/pkg/errors"
)

type NumberService interface {
	Create(number int)
	Get(number int) (string, errors.Error)
	BulkGet() ([]string, errors.Error)
	BulkGetTypes() ([]string, errors.Error)
}
