package models

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"prime-data/ent"
)

var Client *ent.Client

func InitOrmDB() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	Client = client
}