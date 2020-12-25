package services

import (
	"context"
	"prime-data/ent"
	"prime-data/pkg/jwt"
	"prime-data/schema"
	"strconv"
)

type AuthService struct {
	jwt      jwt.IJWTAuth
}


type IAuthService interface {
	Login(ctx context.Context, item *ent.User) (*schema.UserTokenInfo, error)
}

func NewAuthService(jwt jwt.IJWTAuth) IAuthService{
	return &AuthService{
		jwt:      jwt,
	}
}

func (a *AuthService) Login(ctx context.Context, item *ent.User ) (*schema.UserTokenInfo, error){
	token, err := a.jwt.GenerateToken(strconv.Itoa(item.ID))
	if err != nil {
		return  nil, err
	}

	tokenInfo := schema.UserTokenInfo{
		AccessToken:  token.GetAccessToken(),
		RefreshToken: token.GetRefreshToken(),
		TokenType:    token.GetTokenType(),
	}

	return &tokenInfo, nil
}