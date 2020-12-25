package app

import (
	"fmt"
	"context"
	"github.com/google/logger"
	"prime-data/ent"
)

func Migrate(client *ent.Client) {
	ctx := context.Background()
	createUser(ctx, client)
	createPolicy(ctx, client )
}

func createUser(ctx context.Context, client *ent.Client) ( error) {
	_, err := client.User.
		Create().
		SetUsername("duc").
		SetPassword("12345").
		Save(ctx)
	if err != nil {
		logger.Info(err)
		return fmt.Errorf("failed creating user: %v", err)
	}
	return nil
}

func createPolicy(ctx context.Context, client *ent.Client) ( error) {
	_, err := client.CasbinRule.
		Create().
		SetPType("p").
		SetV0("user_a").
		SetV1("/app/hello").
		SetV2("GET").
		Save(ctx)
	if err != nil {
		return  fmt.Errorf("failed creating user: %v", err)
	}

	return nil
}