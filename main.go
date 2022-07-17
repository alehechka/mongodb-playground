package main

import (
	"github.com/alehechka/mongodb-playground/constants"
	"github.com/alehechka/mongodb-playground/database"
	"github.com/alehechka/mongodb-playground/opentel"
	"github.com/alehechka/mongodb-playground/rest"
)

func main() {
	constants.InitializeConstants()
	shutdownTracer, err := opentel.InitTracer()
	check(err)
	defer shutdownTracer()

	disconnectDatabase, err := database.InitializeMongoDB()
	check(err)
	defer disconnectDatabase()

	check(rest.SetupRouter().Run())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
