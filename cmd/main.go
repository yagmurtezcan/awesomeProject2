package main

import (
	"awesomeProject2/internal/entity"
	"awesomeProject2/internal/server"
	"awesomeProject2/pkg/postgres"
)

func main() {
	pstgrs := postgres.NewDatabase()
	pstgrs.AutoMigrate(
		&entity.User{},
	)

	// service

	//handler

	srv := server.NewServer(pstgrs.Db)
	srv.StartServer()
}
