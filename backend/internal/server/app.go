package server

import (
	"backend/internal/config"
	"backend/internal/https/handlers"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/logger"
	recovermw "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/oauth2/authhandler"
)
func new(Cfg config.Config,db *pgxpool.Pool)*fiber.App{
	app:=fiber.New()

	app.Use(logger.New())
//now we have to create midddleware->fiber recover for handling any panic we canno use recover which is builtin function 
	app.Use(recovermw.New())
	app.Use(cors.New(cors.Config){
		AllowOrigins:"*" || []string{cfg.FRONTEND_URL},
		AllowMethods:[]string{"GET","POST","PUT","DELETE","OPTIONS","PATCH"},
		AllowHeaders:[]string{"Origin","Content-Type","Accept","Authorization"},
		AllowCredentials:true,
		
	})
	//intiate our user repo
	userRepo:=repositories.NewUserRepository(db)
    authService:=services.NewAuthService(cfg)
	//initate handler
	authHandler:=handlers.AuthHandler(cfg,authService,userRepo)
	authMiddleware:=middleware.NewAuthMiddleware(cfg,authService,userRepo)
	routes.Register(app,routes.RouteDependecies{
		AuthHandler:authHandler,
		AuthMiddleware:authMiddleware,
	})

	return app
	

}