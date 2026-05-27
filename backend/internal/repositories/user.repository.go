// in this we can use all the query with database related to users
package repositories

import (
	"backend/internal/models"
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	//database connection
	db *pgxpool.Pool

}

type UpsertUserInput struct{
	Email string 
	Name string 
	AvatarURL string
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

//methods
func (r *UserRepository) UpsertByEmail(ctx context.Context, input UpsertUserInput) (*models.User, error){
	var user models.User 
	query := `INSERT INTO users (email, name, avatar_url,role) VALUES ($1, $2, $3,'admin') ON CONFLICT (email) DO UPDATE SET name = EXCLUDED.name, avatar_url = EXCLUDED.avatar_url RETURNING id,email,COALESCE(name,''),COALESCE(avatar_url,''),role,created_at`
	//scan will  copies returned values into variables into user structs field 
	err := r.db.QueryRow(ctx, query,strings.TrimSpace(strings.ToLower(input.Email)), strings.TrimSpace(input.Name), strings.TrimSpace(input.AvatarURL)).Scan(&user.Id, &user.Email, &user.Name, &user.AvatarURL, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert user by email %s: %w",input.Email,err)
	}
	return &user, nil
}

//find by id

func(r *UserRepository)FindById(ctx context.Context,userID string)(*models.User,error){
	var user models.User 
	query := `SELECT id,email,COALESCE(name,''),COALESCE(avatar_url,''),role,created_at FROM users WHERE id = $1`
	err := r.db.QueryRow(ctx, query, strings.TrimSpace(userID)).Scan(&user.Id, &user.Email, &user.Name, &user.AvatarURL,
		&user.Role, &user.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to find user by id %s: %w",userID,err)
	}
	return &user, nil
}