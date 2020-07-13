package main

import (
	"go-api/src/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"go-api/src/middlewares"
	"os"
	"os/signal"
	"go-api/pkg"
	"go-api/src/repository"
	"strconv"
	"syscall"
	_ "github.com/lib/pq"
)

var pgRepository *repository.Repository

func main() {
	initDb()

	r := gin.New()

	api := r.Group("/api")

	api.Use(gin.Logger())
	api.Use(middlewares.Auth)

	v1 := api.Group("/v1")

	r.GET("/", handlers.HandleApp)
	r.GET("/health", handlers.HandleHealth)

	v1.GET("/orders", handlers.HandleListOrders(pgRepository))
	v1.GET("/orders/:orderName", handlers.HandleGetOrder(pgRepository))

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		s := <-sigc
		log.Print("Signal received: ", s.String())
		log.Print("Gracefully shutting down server")
		_ = pgRepository.CloseConn()
		log.Print("Server shut down")
		code, _ := strconv.Atoi(s.String())
		os.Exit(code)
	}()

	if err := r.Run(":90"); err != nil {
		_ = pgRepository.CloseConn()
		log.Fatal("error in starting server: ", err)
	}

}

func initDb() {
	pgRepository = &repository.Repository{
		DB:         nil,
		DBUri:      pkg.GetEnvOrDefault("POSTGRES_URI", "localhost:5432"),
		DBName:     pkg.GetEnvOrDefault("DB_NAME", "postgres"),
		DBUsername: pkg.GetEnvOrDefault("DB_USER", "postgres"),
		DBPassword: pkg.GetEnvOrDefault("DB_PASSWORD", "hunter2"),
		Table:      pkg.GetEnvOrDefault("DB_TABLE", "pods"),
	}

	err := pgRepository.Init()
	if err != nil {
		log.Fatal("error creating database connection: ", err)
	}
}
