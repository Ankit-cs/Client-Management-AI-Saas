//all http handler related to auth related to authentication

package handlers

import (
	"backend/internal/config"
	"backend/internal/repositories"
	"backend/internal/services"
	// "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber"
)

type AuthHandler struct {
	config config.Config
	//auth service
	authService *services.AuthService
	userRepo *repositories.UserRepository
}

func NewAuthHandler(config config.Config, authService *services.AuthService, userRepo *repositories.UserRepository) *AuthHandler {
	return &AuthHandler{config: config, authService: authService, userRepo: userRepo}
}
///start google authentications
func (h *AuthHandler) startGoogleAuth(c *fiber.Ctx) error{
 state,err:=h.authService.GenerateStateToken()
 if err!=nil{
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to generate state token"})
 }
 //we have to keep that in http only moduel wehn google callback er compare with this cookie for verification
 h.authService.SetOauthStateCookie(c,state)
 return  c.Redirect().To(h.authService.BuildGoogleAuthUrl(state))
}