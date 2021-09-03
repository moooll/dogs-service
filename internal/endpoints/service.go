package endpoints

import (
	"github.com/moooll/dogs-service/internal/storage"
	"github.com/moooll/dogs-service/internal/tokens"
)

// Service contains *storage.Storage to interact with storage funcs
type Service struct {
	St *storage.Storage
	V  *Validator
	Sk *tokens.SigningKey
}

// NewService returns new initialized *Service with *storage.Storage
func NewService(st *storage.Storage, v *Validator, sk *tokens.SigningKey) *Service {
	return &Service{
		St: st,
		V:  v,
		Sk: sk,
	}
}
