package services

import (
	"backend/internal/config"
	"crypto/rand"
	"encoding/base64"
	// "encoding/hex"
	"fmt"
	"net/http"
	"time"
	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthService struct {
	config config.Config
	oauthConfig *oauth2.Config
	httpClient *http.Client
}
//google auth login 

type GoogleUserInfo struct{
	ID string `json:"id"`
	Email string `json:"email"`
	VerifiedEmail bool `json:"verified_email"`//check for verifications
	Name string `json:"name"`
	Picture string `json:"picture"`
}
//auth claims for custiom info in jwt
type AuthClaims struct{

	UserID string `json:"user_id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims

}
const oauthStateCookieName="google_oauth_state"

func NewAuthService (cfg config.Config)*AuthService{
	return &AuthService{
		config:cfg,
		oauthConfig: &oauth2.Config{
			ClientID:     cfg.GoogleClientID,
			ClientSecret: cfg.GoogleClientSecret,
			RedirectURL:  cfg.GoogleRedirectURL,
			Scopes:       []string{"openid","email", "profile"},//baisc indentity details
			Endpoint:     google.Endpoint,
		},
		httpClient:&http.Client{ Timeout: 15* time.Second},
	}
}

func (s *AuthService) GenerateStateToken() (string, error) {
   buffer :=make([]byte, 64)
   if _,err:=rand.Read(buffer);err!=nil{
     return "",fmt.Errorf("failed to generate state token :%w",err)	   
    }
   return base64.RawURLEncoding.EncodeToString(buffer), nil//bytes into url safe string 

}
func (s *AuthService) SetOauthStateCookie(c fiber.Ctx,value string){
	c.Cookie(&fiber.Cookie{
		Name: oauthStateCookieName,
		Value: value,
		path:"/",
		Expires: time.Now().Add(10* time.Minute),
		HTTPOnly: true,
		Secure: s.config.CookiesSecure,
		SameSite: s.config.CookiesSameSite,
		Domain:s.config.CookiesDomain,
		MaxAge:20*60,//20 mintue
	})
	return 
}

func (s *AuthService)BuildGoogleAuthUrl(state string) string{
	//AuthCodeURL created google oauth consent url states value is included it combakc to teh google callback page 
	return s.oauthConfig.AuthCodeURL(state)
}