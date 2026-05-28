package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/server"
	"context"
	"log"
)

func main (){
	cfg,err:=config.Load()
	if err !=nil{
		log.Fatal("Load config:",err)
	}``
	ctx:=context.Background()
	db,err:=database.NewPool(ctx,cfg.DatabaseURL)
	if err !=nil{
		log.Fatalf("Connected to Database",err)
	}
	//close the db connection
	defer db.Close()
	//create the server  which is the main server
	app := server.New(cfg, db)
	log.Printf("Server runnig on https://localhost:%s",cfg.Port)
	//listen the port 
	if err := app.Listen(":"+cfg.Port); err != nil {
		log.Fatalf("Listen : %v", err)
	}
}