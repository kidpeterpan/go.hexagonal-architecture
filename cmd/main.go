package main

import (
	"fmt"
	"log"

	"go.hexagonal-architecture/internal/adapters/app/api"
	"go.hexagonal-architecture/internal/adapters/core/arithmetic"
	"go.hexagonal-architecture/internal/adapters/framework/right/db"
	"go.hexagonal-architecture/internal/ports"
)

func main() {
	var core ports.Arithmetic
	core = arithmetic.NewAdapter()
	var database ports.DbPort
	database = db.NewAdapter()
	var application ports.ApiPort
	application = api.NewAdapter(core, database)

	result, err := application.GetAddition(2, 3)
	if err != nil {
		log.Fatalf("GetAddition error: %v", err)
	}
	fmt.Println(result)

	err = database.CloseDbConnection()
	if err != nil {
		panic(err)
	}
}
