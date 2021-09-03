package endpoints

import (
	"github.com/moooll/dogs-service/internal/models"
	log "github.com/sirupsen/logrus"
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

// Register registers the type for validation
func (v *Validator) Register() error {
	err := v.v.Register(models.Dog{})
	if err != nil {
		log.Error("error updating dog: ", err.Error())
		return err
	}

	er := v.v.Register(models.User{})
	if er != nil {
		log.Error("error updating dog: ", er.Error())
		return er
	}

	return nil
}

func (v *Validator) validate(in interface{}) (vi []string, err error) {
	vi, err = v.v.Violations(in)
	if err != nil {
		return vi, err
	}

	return vi, nil
}
