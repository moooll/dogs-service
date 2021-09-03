package main

import (
	"context"

	"github.com/caarlos0/env"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/moooll/dogs-service/internal/config"
	"github.com/moooll/dogs-service/internal/endpoints"
	"github.com/moooll/dogs-service/internal/storage"
	"github.com/moooll/dogs-service/internal/tokens"
	log "github.com/sirupsen/logrus"
)

// @title Dogs API
// @version 1.0
// @description This is a server for creating, reading, writing updating dogs.

// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.NewConfig()
	err := env.Parse(cfg)
	if err != nil {
		log.Error("error parsing config: ", err.Error())
	}

	conn, er := pgxpool.Connect(context.Background(), cfg.DatabaseURI)
	if er != nil {
		log.Error("error connecting to the db: ", er.Error())
	}

	defer conn.Close()

	st := storage.NewStorage(context.Background(), conn)
	v := endpoints.NewValidator()
	err = v.Register()
	if err != nil {
		log.Error("error registering the dog: ", err.Error())
	}

	service := endpoints.NewService(st, v, nil)
	sk := tokens.NewSigningKey([]byte(cfg.JWTSecret))
	authService := endpoints.NewService(st, v, sk)
	e := echo.New()
	r := e.Group("")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(cfg.JWTSecret),
	}))
	r.POST("/dogs", service.Create)
	r.GET("/dogs/:id", service.Read)
	r.GET("/dogs", service.ReadAll)
	r.PUT("/dogs", service.Update)
	r.DELETE("/dogs/:id", service.Delete)
	e.POST("/auth/login", authService.Login)
	e.POST("/auth/refresh", authService.Refresh)
	e.POST("/auth/logout", authService.Logout)
	if err := e.Start(cfg.ServerPort); err != nil {
		log.Error("error starting server: ", err.Error())
	}
}
