package main

import (
	"github.com/alehechka/mongodb-playground/constants"
	"github.com/alehechka/mongodb-playground/database"
	"github.com/alehechka/mongodb-playground/rest"
)

func main() {
	constants.InitializeConstants()

	disconnect, err := database.InitializeMongoDB()
	check(err)
	defer disconnect()

	check(rest.SetupRouter().Run())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
