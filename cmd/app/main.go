package main

import (
	"log"
	
	"mock/internal/dto"
	"mock/internal/repository/logrequest"
	"mock/internal/repository/mock"
	"mock/internal/routes"
	"mock/pkg/gormdb"
)

func main() {
	db, err := gormdb.DB("mock.db")
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&dto.Mock{}, &dto.LogRequest{})
	if err != nil {
		log.Fatal(err)
	}
	
	repoMock := mock.NewRepository(db)
	repoLogRequest := logrequest.NewRepository(db)
	routes.Init(repoMock, repoLogRequest)
}
