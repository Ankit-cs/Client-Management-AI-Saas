package server

import (
	"backend/internal/config"
	"backend/internal/https/handlers"
	"backend/internal/https/middlewares"
	"backend/internal/https/routes"
	"backend/internal/repositories"
	"backend/internal/services"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recovermw "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg config.Config, db *pgxpool.Pool) *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	// now we have to create midddleware->fiber recover for handling any panic we canno use recover which is builtin function 
	app.Use(recovermw.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.FrontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// intiate our user repo
	userRepo := repositories.NewUserRepository(db)
	submissionRepo:=repositories.NewSubmissionRepository(db)
	authService := services.NewAuthService(cfg)
	n8nService:=services.NewN8NService(cfg)
	// initate handler
	authHandler := handlers.NewAuthHandler(cfg, authService, userRepo)
	submissionHandler:=handlers.NewSubmissionHandler(submissionRepo,n8nService)
	adminSubmissionHandler:=handlers.NewAdminSubmissionHandler(submissionRepo)
	authMiddleware := middlewares.NewAuthMiddleware(cfg, authService, userRepo)

	routes.Register(app, routes.RouteDependencies{
		AuthHandler:    authHandler,
		AuthMiddleware: authMiddleware,
		SubmissionHandler: submissionHandler,
		AdminSubmissionHandler: adminSubmissionHandler,
	})

	return app
}
