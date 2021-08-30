package main

import (
	"github.com/caarlos0/env"
	"github.com/labstack/echo/v4"
	"github.com/moooll/dogs-service/internal/config"
	"github.com/moooll/dogs-service/internal/endpoints"
	"github.com/moooll/dogs-service/internal/storage"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.NewConfig()
	err := env.Parse(cfg)
	if err != nil {
		log.Error("error parsing config: ", err.Error())
	}

	st := storage.NewStorage()
	v := endpoints.NewValidator()
	service := endpoints.NewService(st, v)
	e := echo.New()
	e.POST("/dogs", service.Create)
	e.GET("/dogs/:id", service.Read)
	e.GET("/dogs", service.ReadAll)
	e.PUT("/dogs", service.Update)
	e.DELETE("/dogs/:id", service.Delete)
	if err := e.Start(cfg.ServerPort); err != nil {
		log.Error("error starting server: ", err.Error())
	}
}
