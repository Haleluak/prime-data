package app

import (
	"context"
	"prime-data/ent"
)

func Migrate(client *ent.Client) {
	ctx := context.Background()
	createUser(ctx, client)
	createPolicy(ctx, client )
}

func createUser(ctx context.Context, client *ent.Client) ( error) {
	bulk := make([]*ent.UserCreate, 2)
	bulk[0] = client.User.Create().SetUsername("duc").SetPassword("12345")
	bulk[1] = client.User.Create().SetUsername("rayn").SetPassword("12345")

	_, err := client.User.CreateBulk(bulk...).Save(ctx)
	return err
}

func createPolicy(ctx context.Context, client *ent.Client) ( error) {
	bulk := make([]*ent.CasbinRuleCreate, 3)
	bulk[0] = client.CasbinRule.Create().SetPType("p").SetV0("1").SetV1("/app/hello").SetV2("GET").SetV3("").SetV4("").SetV5("")
	bulk[1] = client.CasbinRule.Create().SetPType("p").SetV0("2").SetV1("/app/hello").SetV2("GET").SetV3("").SetV4("").SetV5("")
	bulk[2] = client.CasbinRule.Create().SetPType("p").SetV0("2").SetV1("/app/request").SetV2("POST").SetV3("").SetV4("").SetV5("")

	_, err := client.CasbinRule.CreateBulk(bulk...).Save(ctx)
	return err
}