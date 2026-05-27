package models

import (
	"time"
)	
//user struct can be compare with mongoose schema
type User struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	AvatarURL string `json:"avatar_url"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
} 
//we have to create repository layer
//repositry interact with database and tehn we create services layer
//service interact with handler
//then we create httpHandlers create handler files and route files 4 files total
//user handler
