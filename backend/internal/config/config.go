package config

import (
	"fmt"
	"os"
	"strconv"
    "strings"
	"github.com/joho/godotenv" // to read env file we have to import this file
)
type Config struct{
	AppEnv string
	ServicePort string
	DatabaseURL string
	AppBaseURL string
	FrontendURL string
	JWTSecretKey string
	JWTExpirationTime int
	AuthCookieName string
	GoogleClientID string
	GoogleClientSecret string
	GoogleRedirectURL string
	CookiesDomain string
	CookiesSecure bool
	CookiesSameSite string
}

func Load()(Config,error){
	_= godotenv.Load()// to load env file
	config := Config{
		AppEnv: getEnv("APP_ENV","development"),
		ServicePort: getEnv("SERVICE_PORT","8080"),
		DatabaseURL: getEnv("DATABASE_URL",""),
		AppBaseURL: strings.TrimSpace(getEnv("APP_BASE_URL","")),
		FrontendURL: strings.TrimSpace(getEnv("FRONTEND_URL","")),
		JWTSecretKey: getEnv("JWT_SECRET_KEY",""),
		JWTExpirationTime: getEnvAsInt("JWT_EXPIRATION_TIME",3600),
		AuthCookieName: getEnv("AUTH_COOKIE_NAME","auth_token"),
		GoogleClientID: getEnv("GOOGLE_CLIENT_ID",""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET",""),
		GoogleRedirectURL: getEnv("GOOGLE_REDIRECT_URL",""),
		CookiesDomain: getEnv("COOKIES_DOMAIN",""),
		CookiesSecure: getEnvAsBool("COOKIES_SECURE",false),
		CookiesSameSite: strings.ToLower(getEnv("COOKIES_SAMESITE","Lax")),
	}
	if config.AppEnv==""{
		return Config{},fmt.Errorf("APP_ENV is required")
	}
	if config.ServicePort==""{
		return Config{},fmt.Errorf("SERVICE_PORT is required")
	}
	if config.DatabaseURL==""{
		return Config{},fmt.Errorf("DATABASE_URL is required")
	}
	if config.AppBaseURL==""{
		return Config{},fmt.Errorf("APP_BASE_URL is required")
	}
	if config.FrontendURL==""{
		return Config{},fmt.Errorf("FRONTEND_URL is required")
	}
	if config.JWTSecretKey==""{
		return Config{},fmt.Errorf("JWT_SECRET_KEY is required")
	}
	if config.GoogleClientID==""{
		return Config{},fmt.Errorf("GOOGLE_CLIENT_ID is required")
	}
	if config.GoogleClientSecret==""{
		return Config{},fmt.Errorf("GOOGLE_CLIENT_SECRET is required")
	}
	if config.GoogleRedirectURL==""{
		return Config{},fmt.Errorf("GOOGLE_REDIRECT_URL is required")
	}
	return config,nil
}

func getEnv(key string, fallback string) string {
	value:=os.Getenv(key)// to get env variable
	if value == "" {
		return fallback
	}
	return value
}
func getEnvAsInt(key string, fallback int) int {
	value := os.Getenv(key)
		if value == "" {
			return fallback
		}
		// return atoi(value)
		parser,err:=strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Error parsing %s: %v\n", key, err)
			return fallback
		}
		return parser
}

func getEnvAsBool(key string, fallback bool) bool {
	value := os.Getenv(key)	
	if value == "" {
		return fallback
	}
	// return Boolean (value)
	parser,err:=strconv.ParseBool(value)
	if err != nil {
		fmt.Printf("Error parsing %s: %v\n", key, err)
		return fallback
	}
	return parser
}