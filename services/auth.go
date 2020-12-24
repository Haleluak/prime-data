package services

import (
	"context"
	"prime-data/pkg/jwt"
	"prime-data/schema"
)

type AuthService struct {
	jwt      jwt.IJWTAuth
}


type IAuthService interface {
	Login(ctx context.Context) (*schema.UserTokenInfo, error)
}

func NewAuthService(jwt jwt.IJWTAuth) IAuthService{
	return &AuthService{
		jwt:      jwt,
	}
}

func (a *AuthService) Login(ctx context.Context) (*schema.UserTokenInfo, error){
	token, err := a.jwt.GenerateToken("12")
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