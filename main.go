package main

import (
	"context"
	"go-template/app/template"
	"go-template/configs"
	"go-template/pkg/cache"
	"go-template/pkg/database"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	runtime.GOMAXPROCS(1)
	err := configs.Init("")
	if err != nil {
		panic(err)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	conf := configs.GetConfig()

	db, err := database.ConnectDB(conf)
	if err != nil {
		panic(err)
	}
	defer database.CloseDB()

	redis, err := cache.NewCacheClient(conf.Redis, conf.Secrets)
	if err != nil {
		panic(err)
	}
	defer redis.Close()

	var (
		repo    = template.NewTransactionRepository(db)
		service = template.NewTemplateService(repo)
		handler = template.NewHandler(service)
		router  = gin.New()
	)

	srv := &http.Server{
		Addr:        ":" + conf.App.Port,
		Handler:     router,
		ReadTimeout: conf.App.Timeout,
	}

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Ready!!")
	})

	router.POST("/template", handler.Handler)

	log.Printf("Server starting on port %s", conf.App.Port)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	gracefulShutdown(srv)
}

func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
