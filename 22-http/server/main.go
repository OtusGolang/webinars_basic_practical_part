package main

import (
	"context"
	"gohttp2/internal/app"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newApp, err := app.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = newApp.Start()
	if err != nil {
		log.Fatal(err)
	}
}

// curl -uri http://localhost:8989 -method get
// curl -uri http://localhost:8989/login -method post -body '{"username":"a", "password":"b"}'
// curl -uri http://localhost:8989/login -method post -body '{"username":"a@b.c", "password":"b"}'
