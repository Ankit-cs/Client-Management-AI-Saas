package routes

import (
	"backend/internal/https/handlers"
	"backend/internal/https/middleware"
	"github.com/gofiber/fiber/v3"	
)

type RouteDependencies struct{
	AuthHandler *handlers.AuthHandler
	AuthMiddleware *middleware.AuthMiddleware
}

func Register(app *fiber.App, deps RouteDependencies){
	auth:=app.Group("/auth")
	auth.Gt("/google",deps.AuthHandler.StartGoogleAuth)
	auth.Get("/google/callback",deps.AuthHandler.GoogleAuthCallback)
	auth.Post("/logout",deps.AuthHandler.Logout())
   auth.Get("/me",deps.AuthMiddleware.RequireAuth(),deps.AuthHandler.GetUserInfo)
}