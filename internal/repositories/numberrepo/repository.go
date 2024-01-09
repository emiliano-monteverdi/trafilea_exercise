package numberrepo

import "github.com/Trafilea/internal/core/ports"

type repository struct {
	storageClient map[int]string
}

func New(storageClient map[int]string) ports.NumberRepository {
	return &repository{storageClient: storageClient}
}
