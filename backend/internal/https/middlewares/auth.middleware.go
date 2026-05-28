package middlewares

import (
	"backend/internal/config"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/services"
     "strings"
	"github.com/gofiber/fiber/v2"
)

//we have to protect routes so we are creating a auth middleware
//we have to get the user information from the JWT token and store it in the context
//this key is used to store in fiber locals
const currentUserLocalkey="current_user"

func ExtractCurrentUserLocalKey() string{
	return currentUserLocalkey
}


func CurrentUserFromContext(c *fiber.Ctx) (*models.User,bool){
	//locals make it possible to access the data from any where in the handler and also by the middlewares therefore available to all follwing routes that matches the request 
currentUser,ok:=c.Locals(currentUserLocalkey).(*models.User)
return currentUser,ok	

}

type AuthMiddleware struct{
	config config.Config
	authService *services.AuthService
	userRepo *repositories.UserRepository
}
func NewAuthMiddleware(cfg config.Config, authService *services.AuthService, userRepo *repositories.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{config: cfg, authService: authService, userRepo: userRepo}
}
//auth and get user roles and admin protection routes also 
//this middleware protects the routes
func(m *AuthMiddleware)Protect(c *fiber.Ctx) error {
	return nil 
}

func (m *AuthMiddleware) RequireAuth() fiber.Handler{

return func(c *fiber.Ctx) error{
   tokenString:=c.Cookies(m.config.AuthCookieName,"")
    if tokenString == ""{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"Auth cookie missing",
		})

	}
	claims,err:=m.authService.ParseToken(tokenString)
    if err!=nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"Invalid auth token",
		})
	}
	//fin by id if 
	user,err:=m.userRepo.FindById(context.Background(),claims.UserID)
	if err!=nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"Invalid auth token",
		})
	}
	//now we have to  store the user in the context
	c.Locals(currentUserLocalkey,user)

return c.Next()
	
	}
}

func (m *AuthMiddleware)RequireAdmin() fiber.Handler{
	return func (c *fiber.Ctx) error {
		currentUser,ok:=CurrentUserFromContext(c)
		if !ok || currentUser.Role!= "admin"{
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message":"Unauthorized access",
			})
		}
		if !strings.EqualFold(currentUser.Role,"admin"){
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message":"Admin access required",
			})
		}
		return c.Next()
		
	}
}