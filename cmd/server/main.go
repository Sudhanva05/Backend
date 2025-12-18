package main

import (
	"log"

	db "github.com/Sudhanva05/Backend/db/sqlc"
	"github.com/Sudhanva05/Backend/internal/handler"
	"github.com/Sudhanva05/Backend/internal/logger"
	"github.com/Sudhanva05/Backend/internal/routes"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {

	// Initialize logger
	logger.InitLogger()
	defer logger.Log.Sync()

	logger.Log.Info("Starting server")

	app := fiber.New()

	conn, err := db.NewPostgresConnection()
	if err != nil {
		logger.Log.Fatal("Database connection failed", zap.Error(err))
	}
	defer conn.Close()

	queries := db.New(conn)
	userHandler := handler.NewUserHandler(queries)

	routes.RegisterUserRoutes(app, userHandler)

	logger.Log.Info("Server running on port 3000")

	log.Fatal(app.Listen(":3000"))
}
