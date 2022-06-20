package main

import (
	"awesomeProject2/internal/entity"
	"awesomeProject2/internal/server"
	"awesomeProject2/pkg/postgres"
)

func main() {
	createDBTable := postgres.NewDatabase()
	createDBTable.AutoMigrate(
		&entity.User{},
	)

	srv := server.NewServer(createDBTable.Db)
	srv.StartServer()
}
