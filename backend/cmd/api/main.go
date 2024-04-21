package main

import (
	"backend/internal/user/controller"
	"backend/internal/user/repository"
	"backend/internal/user/service"
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func main() {
	connStr := "host=localhost port=5432 user=test dbname=test_users password=test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	userRepo := repository.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router := gin.Default()

	router.POST("/register", userController.RegisterUser)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "This is Steven Portfolio."})
	})
	router.Run()
}
