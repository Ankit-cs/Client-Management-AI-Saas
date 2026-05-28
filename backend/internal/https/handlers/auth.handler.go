//all http handler related to auth related to authentication

package handlers

import (
	"backend/internal/config"
	"backend/internal/repositories"
	"backend/internal/services"
	"context"

	"github.com/gofiber/fiber/v3"
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
func (h *AuthHandler) startGoogleAuth(c fiber.Ctx) error{
 state,err:=h.authService.GenerateStateToken()
 if err!=nil{
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to generate state token"})
 }
 //we have to keep that in http only moduel wehn google callback er compare with this cookie for verification
 h.authService.SetOauthStateCookie(c,state)
 return  c.Redirect().To(h.authService.BuildGoogleAuthUrl(state))
}
//most important is googleCallBack in this when user visit will be check all the authentication and redirects the logic 
func (h *AuthHandler) googleCallBack(c fiber.Ctx) error{
	stateFromQuery:=c.Query("state")
	stateFromCookie:=h.authService.ReadOauthStateCookie(c)

	if stateFromCookie== "" || stateFromQuery== "" || stateFromCookie!=stateFromQuery{
		h.authService.ClearOauthStateCookie(c)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid oauth State"})
	}
	h.authService.ClearOauthStateCookie(c)
	//auth code from googles callback url 
	//temp code and will be exchanges for  a google token access

	code:=c.Query("code")
	if code ==""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Missing auth code"})
	}

	//exchange google auth code for users google infromation 

	googleUser,err:=h.authService.ExchangeGoogleAuthCode(context.Background(),code)
	if err!=nil{
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": "Failed to fetch google user information",
			"error":err.Error(),
		})
	}

	//now we have to check the email exist in database or not and will perform the logic based on it 
    user,err:=h.userRepo.FindByEmail(context.Background(),repositories.UpsertUserInput{
		Email: googleUser.Email,
		Name:googleUser.Name,
		AvatarURL: googleUser.Picture,
	})

	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Failed to upsert user in our DB",
		})
	}
	//generate JWT for the user
	token,err=h.authService.SignJWT(user)
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to sign JWT token",
			"error":err.Error(),
		})
	}
	h.authService.SetAuthCookie(c,token)
	return c.Redirect().To(h.config.FrontendURL+"/")
}


// get user infromation 
func (h *AuthHandler) GetUserInfo(c fiber.Ctx) error{
	currentUser,ok:=c.Locals(middlewares.ExtractCurrentUserLocalKey()).(*models.User)
	if !ok || currentUser ==nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized access",
		})
	}
	return c.JSON(fiber.Map{
		"user":currentUser,
	})

}

func (h *AuthHandler) Logout(c fiber.Ctx) error{
	h.authService.ClearAuthCookie(c)
	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
	//add cookie  
}