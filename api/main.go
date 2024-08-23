package main

import (
	"log"
	"os"
	"wachirawut_agnos_backend/config"
	"wachirawut_agnos_backend/internal/controllers"
	"wachirawut_agnos_backend/internal/migrations"
	"wachirawut_agnos_backend/internal/repositories"
	"wachirawut_agnos_backend/internal/usecases"
	"wachirawut_agnos_backend/pkg/database/postgresql"
	"wachirawut_agnos_backend/pkg/middleware"
	"wachirawut_agnos_backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	conf := config.NewConfig()

	postgreSQLClient := postgresql.ConnectDB(conf.PostgresqlClient)
	migrations.Migrate(postgreSQLClient)

	var (
		repositories = repositories.NewRepository(postgreSQLClient)
		usecases     = usecases.NewUsecase(repositories)
		controllers  = controllers.NewController(usecases)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// routes
	routes.GroupRoutes(server, controllers)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	if err := server.Run(":" + port); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
