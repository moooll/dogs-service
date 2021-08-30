package main

import (
	"github.com/labstack/echo/v4"
	"github.com/moooll/dogs-service/intenal/endpoints"
	"github.com/moooll/dogs-service/intenal/storage"
	log "github.com/sirupsen/logrus"
)

func main() {
	// init config
	// init storage
	st := storage.NewStorage()
	service := endpoints.Service{
		St: st,
	}
	// start server
	e := echo.New()
	e.POST("/dogs", service.Create)
	e.GET("/dogs/:id", service.Read)
	e.GET("/dogs/rand", endpoints.RandDog)
	e.GET("/dogs", service.ReadAll)
	e.PUT("/dogs", service.Update)
	e.DELETE("/dogs/:id", service.Delete)
	if err := e.Start(":8080"); err != nil {
		log.Errorln("error starting server: ", err.Error())
	}
}
