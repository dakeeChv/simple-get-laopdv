package main

import (
	"lao-pdv/src/router"
)

func main() {
	r := router.New()
	r.Logger.Fatal(r.Start(":3300"))
}
