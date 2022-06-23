package main

import (
	"awesomeProject2/internal/model"
	"awesomeProject2/internal/server"
	"awesomeProject2/internal/validator"
	"awesomeProject2/pkg/postgres"
)

func main() {
	createDBTable := postgres.NewDatabase()
	createDBTable.AutoMigrate(
		&model.User{},
	)

	validator.Init()

	srv := server.NewServer(createDBTable.Db)
	srv.StartServer()

}
