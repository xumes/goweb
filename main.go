package main

import (
	"fmt"
	"goweb/app"
	"log"
	"os"
)

func main() {
	fmt.Println("starter point")

	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
