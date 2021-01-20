package main

import (
	"context"
	"log"
	"os"
	"rt-chat/internal/db/mgo"
	"rt-chat/service"
	v1 "rt-chat/service/v1"
)

func main() {
	ctx := context.Background()
	client, err := mgo.NewClient(ctx, mgo.WithURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	if err := service.Serve(v1.NewMux(&v1.Tools{Client: client}), os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
