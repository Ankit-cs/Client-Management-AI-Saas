package services

import (
	"backend/internal/config"
	"crypto/rand"
	"encoding/base64"
	"strings"

	// "encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
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

//function for GoogleCallback 

func (s *AuthService) ReadOauthStateCookie(c fiber.Ctx) string{
	return c.Cookies(oauthStateCookieName,"")
}
//func to delete the cookie 

func(s *AuthService)clearOauthStateCookie(c fiber.Ctx){
	c.Cookie(&fiber.Cookie{
		Name:oauthStateCookieName,
		Value:"",
		Path:"/",
		HTTPOnly:true,
		Secure: s.config.CookiesSecure,
		SameSite: s.config.CookiesSameSite,
		Domain:s.config.CookiesDomain,
		//expire the cookie in the past for better browser compatability
		Expires:time.Unix(0,0),
	})
}
//exchabge Google auth code to get out user information from users google 
func (s *AuthService) ExchangeGoogleAuthCode(ctx context.Context,code string)(*GoogleUserInfo,err){

	token,err:=s.oauthConfig.Exchange(ctx,code)//convert authorisation code into token
	if err!=nil{
		return nil,fmt.Errorf("failed to exchange auth code: %w",err)
	}
	req,err:=http.NewRequestWithContext(ctx,http.MethodGet,"https://www.google.com/oauth2/v2/userinfo",nil)
	if err!=nil{
		return nil,fmt.Errorf("failed to create request for userinfo: %w",err)
	}
	
	req.Header.Set("Authorization","Beared"+token.AccessToken)
	respone,err:=s.httpClient.Do(req)
	if err!=nil{
		return nil,fmt.Errorf("failed to perform userinfo request: %w",err)
	}
	// close response body to avoid resource leekages after the reading is done
	defer respone.Body.Close()

	if response.StatusCode!=http.StatusOK{
		return nil,fmt.Errorf("google returned non-ok status")
	}
	
	var userInfo GoogleUserInfo
	if err:=json.NewDecoder(respone.Body).Decode(&userInfo);err!=nil{
		return nil,fmt.Errorf("failed to decode userinfo response: %w",err)
	}
	if strings.TrimSpace(userInfo.Email)==""{
		return nil,fmt.Errorf("google user email is empty")
	}
	//finally return the user info
	return &userInfo,nil
}
func (s *AuthService) SignJWT(user *models.User)(string,error){
	expiresAt=time.Now().Add(time.Duration(s.config.JWTExpiresInHours)*time.Hour)
	claims:=AuthClaims{
		UserID: user.Id,
		Email: user.Email,
		Name: user.Name,
		RegisteredClaims:jwt.RegisteredClaims{
			Subject: user.ID,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:jwt.NewNumericDate(time.Now()),
			
		}
	}
	toke:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	signed,err:=token.SignedString([]byte(s.config.JWTSecret))
	if err!=nil{
		return "",fmt.Errorf("failed to sign token: %w",err)
	}
	return signed,nil
}

func(s *AuthService)SetAuthCookie(c fiber.Ctx,token string){
	maxAge:=s.config.JWTExpiresInHours * 60 *60//second
	c.Cookie(&fiber.Cookie{
		Name:s.config.AuthCookieName,
		Value:token,
		Path:"/",
		HTTPOnly:true,
		Secure: s.config.CookiesSecure,
		SameSite: s.config.CookiesSameSite,
		Domain:s.config.CookiesDomain,
		MaxAge:maxAge,
	})
}

func (s *AuthService) ParseToken(tokenString string) (*AuthClaims, error) {
	token,err:=jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (any, error) {
		if _,ok:=token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil,fmt.Errorf("unexpected signing method: %v",token.Header["alg"])
		}
		return []byte(s.config.JWTSecret),nil
	})
	if err!=nil{
		return nil,fmt.Errorf("failed to parse token: %w",err)
	}
	//custom generic claims type into our custom auth claims 	type
	//and also check if the token is valid or not 
	claims,ok:=token.Claims.(*AuthClaims)
	if !ok || !token.Valid{
		return nil,fmt.Errorf("invalid jwt token")
	}
	return claims,nil
	
}

func(s *AuthService)ClearAuthCookie(c fiber.Ctx){
	c.Cookie(&fiber.Cookie{
		Name:s.config.AuthCookieName,
		Value:token,
		Path:"/",
		HTTPOnly:true,
		Secure: s.config.CookiesSecure,
		SameSite: s.config.CookiesSameSite,
		Domain:s.config.CookiesDomain,
		MaxAge:-1,
		Expires:time.Unix(0,0),
	})
}