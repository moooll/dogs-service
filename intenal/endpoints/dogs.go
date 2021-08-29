package endpoints

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/moooll/dogs-service/intenal/models"
	"github.com/moooll/dogs-service/intenal/storage"
)

type Service struct {
	st storage.Storage
}

func (s *Service) Create(c echo.Context) error {
	dog := models.Dog{}
	err := c.Bind(&dog)
	if err != nil {
		log.Errorln("error creating dog: ", err.Error())
		return c.JSON(500, "error creating dog")
	}

	er := s.st.Create(dog)
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

	dog, er := s.st.Read(id)
	if er != nil {
		log.Errorln("error reading dog: ", er.Error())
		return c.JSON(204, "error occured, nothing found")
	}

	return c.JSON(200, dog)
}

func (s *Service) ReadAll(c echo.Context) error {
	dogs, err := s.st.ReadAll()
	if err != nil {
		log.Errorln("error reading all dogs: ", err.Error())
		return c.JSON(204, "error occured, nothing found")
	}

	return c.JSON(200, dogs)
}

// func (s *Service) Update(c echo.Context) error {
// 	dog, err := c.Bind(&dog)
// 	dog, err := s.st.Update()
// }

// func (s *Service) Delete(c echo.Context) error {

// }
