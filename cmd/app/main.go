package main

import (
	"log"
	
	"github.com/dhnikolas/mock/internal/dto"
	"github.com/dhnikolas/mock/internal/repository/logrequest"
	"github.com/dhnikolas/mock/internal/repository/mock"
	"github.com/dhnikolas/mock/internal/routes"
	"github.com/dhnikolas/mock/pkg/gormdb"
)

func main() {
	db, err := gormdb.DB("/var/mock/mock.db")
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
