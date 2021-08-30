package endpoints

import (
	"github.com/moooll/dogs-service/internal/models"
	"github.com/twharmon/govalid"
)

// Validator contains fields for validation of http requests
type Validator struct {
	v *govalid.Validator
}

// NewValidator returns new initialized validator
func NewValidator() *Validator {
	v := govalid.New()
	return &Validator{
		v: v,
	}
}

func (v *Validator) validate(dog models.Dog) (vi []string, err error) {
	vi, err = v.v.Violations(dog)
	if err != nil {
		return vi, err
	}

	return vi, nil
}
