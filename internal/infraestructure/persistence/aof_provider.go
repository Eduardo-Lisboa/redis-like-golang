package persistence

import (
	"redis-like-golang/internal/domain/repository"
)

type AOFProviderOption struct {
	Enabled  bool
	filepath string
}

func NewAOFProvider(opt AOFProviderOption) (repository.PersistenceRepository, error) {
	if !opt.Enabled {
		return nil, nil
	}

	return NewAOF(opt.filepath)

}
