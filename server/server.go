package server

import (
	"context"
	"github.com/aliakbariaa1996/URL-Shortening/config"
	"github.com/aliakbariaa1996/URL-Shortening/internal/api"
	httpx "github.com/aliakbariaa1996/URL-Shortening/internal/http"
	loggerx "github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"time"
)

func RunServer(cfg *config.Config, logger *loggerx.Logger) error {
	// HTTP Server
	router := httpx.InitRouter()
	server, err := api.NewServer(router, cfg, logger)
	if err != nil {
		return err
	}
	server.Server.Addr = ":" + cfg.Port
	server.Server.Handler = router
	server.Server.ReadTimeout = 10 * time.Second
	server.Server.WriteTimeout = 10 * time.Second
	server.Server.MaxHeaderBytes = 1 << 20
	go func() {
		if err := server.Server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return server.Server.Shutdown(ctx)
}
