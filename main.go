package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"stock-and-order-management/infrastructure/config"
	"stock-and-order-management/infrastructure/database"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	loadConfig, err := config.LoadConfig("./infrastructure/config/env")
	if err != nil {
		panic(err)
	}

	config.AppConfig = loadConfig

	if err := database.InitDatabase(); err != nil {
		panic(err)
	}
}

func main() {
	app := gin.Default()
	app.SetTrustedProxies(nil)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGALRM)

	port := config.AppConfig.Port

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      app,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Start server : %s", port)
		}
	}()

	<-quit

	timeoutShutdown := 30 * time.Second
	timeoutShutdownContent, cancel := context.WithTimeout(context.Background(), timeoutShutdown)
	defer cancel()

	if err := server.Shutdown(timeoutShutdownContent); err != nil {
		log.Printf("%+v\n", err.Error())
	}
}
