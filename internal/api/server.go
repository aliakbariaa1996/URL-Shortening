package api

import (
	"context"
	"fmt"
	"github.com/aliakbariaa1996/URL-Shortening/config"
	"github.com/aliakbariaa1996/URL-Shortening/internal/services/shorteningURL"
	"github.com/aliakbariaa1996/URL-Shortening/internal/services/shorteningURL/store"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	loggerx "github.com/sirupsen/logrus"
)

type Server struct {
	*echo.Echo

	logger     *loggerx.Logger
	ss         *ServiceStorage
	cfg        *config.Config
	handler    Handler
	middleware Middleware
}

type ServiceStorage struct {
	shorteningURL shorteningURL.UseService
	db            *redis.Client
}

type Handler struct {
	db     *redis.Client
	logger *loggerx.Logger
}

type Middleware struct {
	db     *redis.Client
	logger *loggerx.Logger
}

func NewServer(router *echo.Echo, cfg *config.Config, logger *loggerx.Logger) (*Server, error) {
	var err error
	s := &Server{
		Echo:   router,
		cfg:    cfg,
		logger: logger,
	}
	db := initDB(cfg)
	s.ss = NewServiceStorage(cfg, logger, db)
	s.handler = Handler{db: db, logger: logger}
	s.middleware = Middleware{db: db, logger: logger}

	// routes init
	s.initRoutes()
	return s, err
}

func NewServiceStorage(cfg *config.Config, logger *loggerx.Logger, db *redis.Client) *ServiceStorage {
	shStore := store.New(db)
	return &ServiceStorage{
		shorteningURL: shorteningURL.NewShorteningCase(cfg, logger, shStore),
		db:            db,
	}
}

func initDB(cfg *config.Config) *redis.Client {
	var ctx = context.Background()
	storeDSN := fmt.Sprintf("%s:%s", cfg.DBHost, cfg.DBPort)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     storeDSN,
		Password: "",
		DB:       0,
	})
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}
	loggerx.WithError(err).Printf("\nRedis started successfully: pong message = {%s}", pong)
	return redisClient
}
