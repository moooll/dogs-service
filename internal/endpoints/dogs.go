// Package endpoints contains http-server endpoints
package endpoints

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/moooll/dogs-service/internal/models"
	"github.com/moooll/dogs-service/internal/storage"
)

// Service contains *storage.Storage to interact with storage funcs
type Service struct {
	St *storage.Storage
	V  *Validator
}

// NewService returns new initialized *Service with *storage.Storage
func NewService(st *storage.Storage, v *Validator) *Service {
	return &Service{
		St: st,
		V:  v,
	}
}

// Create endpoint creates new dog
func (s *Service) Create(c echo.Context) error {
	dog := models.Dog{}
	err := c.Bind(&dog)
	if err != nil {
		log.Error("error creating dog: ", err.Error())
		return echo.NewHTTPError(500, "error creating dog")
	}

	errr := s.V.v.Register(models.Dog{})
	if errr != nil {
		log.Error("error updating dog: ", err.Error())
		return echo.NewHTTPError(500, "error occurred")
	}

	vi, e := s.V.validate(dog)
	if e != nil {
		log.Error("error creating dog: "+fmt.Sprint(vi), e.Error())
		return echo.NewHTTPError(400, "error creating dog, check the input")
	} else if vi != nil && e == nil {
		log.Error("error creating dog: ", fmt.Sprint(vi))
		return echo.NewHTTPError(400, "error creating dog, check the input")
	}

	er := s.St.Create(dog)
	if er != nil {
		log.Error("error creating dog: ", er.Error())
		return echo.NewHTTPError(500, "error creating dog")
	}

	return nil
}

// Read endpoint returns a dog by id or 204 if no sych dog exists
func (s *Service) Read(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error("error reading dog: ", err.Error())
		return echo.NewHTTPError(204, "error occurred, nothing found")
	}

	dog, er := s.St.Read(id)
	if er != nil {
		log.Error("error reading dog: ", er.Error())
		return echo.NewHTTPError(204, "error occurred, nothing found")
	}

	return c.JSON(200, dog)
}

// ReadAll returns all the dog in the storage
func (s *Service) ReadAll(c echo.Context) error {
	dogs, err := s.St.ReadAll()
	if err != nil {
		log.Error("error reading all dogs: ", err.Error())
		return echo.NewHTTPError(204, "error occurred, nothing found")
	}

	return c.JSON(200, dogs)
}

// Update updates the dog by id to what provided in body, and if no such dog exists, creates it
func (s *Service) Update(c echo.Context) error {
	dog := models.Dog{}
	err := c.Bind(&dog)
	errr := s.V.v.Register(models.Dog{})
	if errr != nil {
		log.Error("error updating dog: ", err.Error())
		return echo.NewHTTPError(500, "error occurred")
	}

	vi, e := s.V.validate(dog)
	if e != nil {
		log.Error("error creating dog: "+fmt.Sprint(vi), e.Error())
		return echo.NewHTTPError(400, "error creating dog, check the input")
	} else if vi != nil && e == nil {
		log.Error("error creating dog: ", fmt.Sprint(vi))
		return echo.NewHTTPError(400, "error creating dog, check the input")
	}

	dog, er := s.St.Update(dog)
	if er != nil {
		log.Error("error updating dog: ", err.Error())
		return echo.NewHTTPError(500, "error occurred")
	}

	return c.JSON(200, dog)
}

// Delete deletes a dog by id
func (s *Service) Delete(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error("error deleting dog: ", err.Error())
		return echo.NewHTTPError(500, "error")
	}

	er := s.St.Delete(id)
	if er != nil {
		return echo.NewHTTPError(500, "error")
	}

	return c.JSON(200, "deleted")
}
