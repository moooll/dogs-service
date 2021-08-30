package endpoints

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"

	"github.com/moooll/dogs-service/intenal/models"
	"github.com/moooll/dogs-service/intenal/storage"
)

type Service struct {
	St *storage.Storage
}

func (s *Service) Create(c echo.Context) error {
	dog := models.Dog{}
	err := c.Bind(&dog)
	if err != nil {
		log.Errorln("error creating dog: ", err.Error())
		return c.JSON(500, "error creating dog")
	}

	er := s.St.Create(dog)
	if er != nil {
		log.Errorln("error creating dog: ", er.Error())
		return c.JSON(500, "error creating dog")
	}

	return nil
}

func (s *Service) Read(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Errorln("error reading dog: ", err.Error())
		return c.JSON(204, "error occured, nothing found")
	}

	dog, er := s.St.Read(id)
	if er != nil {
		log.Errorln("error reading dog: ", er.Error())
		return c.JSON(204, "error occured, nothing found")
	}

	return c.JSON(200, dog)
}

func (s *Service) ReadAll(c echo.Context) error {
	dogs, err := s.St.ReadAll()
	if err != nil {
		log.Errorln("error reading all dogs: ", err.Error())
		return c.JSON(204, "error occured, nothing found")
	}

	return c.JSON(200, dogs)
}

func (s *Service) Update(c echo.Context) error {
	dog := models.Dog{}
	err := c.Bind(&dog)
	dog, er := s.St.Update(dog)
	if er != nil {
		log.Errorln("error updating dog: ", err.Error())
		return c.JSON(500, "error occured")
	}

	return c.JSON(200, dog)

}

func (s *Service) Delete(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Errorln("error deleting dog: ", err.Error())
		return c.JSON(500, "error")
	}

	er := s.St.Delete(id)
	if er != nil {
		return c.JSON(500, "error")
	}

	return c.JSON(200, "deleted")
}

func RandDog(c echo.Context) error {
	id := uuid.New()
	name := randstr.String(8, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	breed := randstr.String(8, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	color := randstr.String(8, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	age := rand.Float32() * 15
	price := rand.Float32() * 15
	return c.JSON(200, models.Dog{
		ID:    id,
		Name:  name,
		Breed: breed,
		Color: color,
		Age:   age,
		Price: price,
	})
}
