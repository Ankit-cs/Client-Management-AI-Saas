package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/joho/godotenv" // to read env file we have to import this file
)

type Config struct {
	AppEnv             string
	ServicePort        string
	DatabaseURL        string
	AppBaseURL         string
	FrontendURL        string
	JWTSecretKey       string
	JWTExpirationTime  int
	AuthCookieName     string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
	CookiesDomain      string
	CookiesSecure      bool
	CookiesSameSite    string
}

func Load() (Config, error) {
	loadEnvFile()
	config := Config{
		AppEnv:             getEnv("APP_ENV", "development"),
		ServicePort:        getEnvAny("SERVICE_PORT", "PORT", "8080"),
		DatabaseURL:        getEnv("DATABASE_URL", ""),
		AppBaseURL:         strings.TrimSpace(getEnvAny("APP_BASE_URL", "BACKEND_URL", "")),
		FrontendURL:        strings.TrimSpace(getEnv("FRONTEND_URL", "")),
		JWTSecretKey:       getEnvAny("JWT_SECRET_KEY", "JWT_SECRET", ""),
		JWTExpirationTime:  getEnvAsIntAny(3600, "JWT_EXPIRATION_TIME", "JWT_EXPIRES_IN_HOURS"),
		AuthCookieName:     getEnv("AUTH_COOKIE_NAME", "auth_token"),
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		GoogleRedirectURL:  getEnv("GOOGLE_REDIRECT_URL", ""),
		CookiesDomain:      getEnvAny("COOKIES_DOMAIN", "COOKIE_DOMAIN", ""),
		CookiesSecure:      getEnvAsBoolAny(false, "COOKIES_SECURE", "COOKIE_SECURE"),
		CookiesSameSite:    strings.ToLower(getEnvAny("COOKIES_SAMESITE", "COOKIE_SAME_SITE", "Lax")),
	}
	if config.AppEnv == "" {
		return Config{}, fmt.Errorf("APP_ENV is required")
	}
	if config.ServicePort == "" {
		return Config{}, fmt.Errorf("SERVICE_PORT is required")
	}
	if config.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}
	if config.AppBaseURL == "" {
		return Config{}, fmt.Errorf("APP_BASE_URL is required")
	}
	if config.FrontendURL == "" {
		return Config{}, fmt.Errorf("FRONTEND_URL is required")
	}
	if config.JWTSecretKey == "" {
		return Config{}, fmt.Errorf("JWT_SECRET_KEY is required")
	}
	if config.GoogleClientID == "" {
		return Config{}, fmt.Errorf("GOOGLE_CLIENT_ID is required")
	}
	if config.GoogleClientSecret == "" {
		return Config{}, fmt.Errorf("GOOGLE_CLIENT_SECRET is required")
	}
	if config.GoogleRedirectURL == "" {
		return Config{}, fmt.Errorf("GOOGLE_REDIRECT_URL is required")
	}
	return config, nil
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key) // to get env variable
	if value == "" {
		return fallback
	}
	return value
}

func getEnvAny(keys ...string) string {
	if len(keys) == 0 {
		return ""
	}
	fallback := keys[len(keys)-1]
	for _, key := range keys[:len(keys)-1] {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	// return atoi(value)
	parser, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("Error parsing %s: %v\n", key, err)
		return fallback
	}
	return parser
}

func getEnvAsIntAny(fallback int, keys ...string) int {
	for _, key := range keys {
		if value := os.Getenv(key); value != "" {
			parser, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("Error parsing %s: %v\n", key, err)
				return fallback
			}
			return parser
		}
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	// return Boolean (value)
	parser, err := strconv.ParseBool(value)
	if err != nil {
		fmt.Printf("Error parsing %s: %v\n", key, err)
		return fallback
	}
	return parser
}

func getEnvAsBoolAny(fallback bool, keys ...string) bool {
	for _, key := range keys {
		if value := os.Getenv(key); value != "" {
			parser, err := strconv.ParseBool(value)
			if err != nil {
				fmt.Printf("Error parsing %s: %v\n", key, err)
				return fallback
			}
			return parser
		}
	}
	return fallback
}

func loadEnvFile() {
	for _, path := range []string{"../../.env", "../.env", ".env"} {
		if _, err := os.Stat(path); err == nil {
			_ = godotenv.Overload(path)
			return
		}
	}
}
