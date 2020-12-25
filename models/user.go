package models

import (
	"context"
	"fmt"
	"prime-data/ent"
	"prime-data/ent/user"
)

func QueryUser(ctx context.Context, client *ent.Client, username, password string) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Username(username), user.Password(password)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}

	return u, nil
}

func QueryRole(ctx context.Context, client *ent.Client) error{
	u, err := client.CasbinRule.
		Query().All(ctx)
	if err != nil {
		return  fmt.Errorf("failed querying user: %v", err)
	}

	fmt.Print(u)
	return nil
}