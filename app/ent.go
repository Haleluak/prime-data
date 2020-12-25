package app

import (
	"log"
	"context"
	"prime-data/ent"
	_ "github.com/mattn/go-sqlite3"
)

func InitOrmDB()  *ent.Client{
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}