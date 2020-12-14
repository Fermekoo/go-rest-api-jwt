package main

import (
	"github.com/Fermekoo/blog-dandi/db"
	"github.com/Fermekoo/blog-dandi/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1200"))
}
