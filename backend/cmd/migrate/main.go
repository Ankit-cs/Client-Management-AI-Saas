package main // this is going to use logic of migrator.go file
import (
	"backend/internal/config"
	"backend/internal/database"
	"context"
	"log"
)
func main() {
	// we will call the function to run the migration config our env variables 
	cfg, err := config.Load();
	if err != nil {
		log.Fatalf("Failed to load configuration: ", err)
	}
	ctx:=context.Background()//it is creating empty context
	db,err:=database.NewPool(ctx,cfg.DatabaseURL)
	if err!=nil{
	log.Fatalf("Failed to connect to database: ", err)
	}
	defer db.Close()//it is closing the connection to database after the migration is done when the main function finishes
	//this runs as the end of the function
	
}
