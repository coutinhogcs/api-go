package main

import (
	"api-go/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
}
