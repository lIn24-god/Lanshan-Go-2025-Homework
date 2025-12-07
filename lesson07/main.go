package main

import (
	"lesson06/api"
)

func main() {
	r := api.InitRouterGin()

	r.Run(":8080")
}
