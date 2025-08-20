package main

import (
	"github.com/yadieloscar/budget-api/internal/api"
)

func main() {

	r := api.SetupRouter()

	r.Run() // listen and serve on 0.0.0.0:8080
}
