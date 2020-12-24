package app

import (
	"github.com/dgrijalva/jwt-go"
	"prime-data/config"
	"prime-data/pkg/errors"
	jwtAuth "prime-data/pkg/jwt"
)

func InitAuth() (jwtAuth.IJWTAuth, error) {
	var opts []jwtAuth.Option
	//access token
	opts = append(opts, jwtAuth.WithKeyFunc(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrTokenInvalid
		}
		return []byte(config.SigningKey), nil
	}))
	if config.Expired != 0 {
		opts = append(opts, jwtAuth.WithExpired(config.Expired))
	}
	opts = append(opts, jwtAuth.WithSigningKey([]byte(config.SigningKey)))

	//refresh token
	opts = append(opts, jwtAuth.WithKeyFuncRefresh(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrTokenInvalid
		}
		return []byte(config.SigningRefreshKey), nil
	}))

	if config.ExpiredRefreshToken != 0 {
		opts = append(opts, jwtAuth.WithExpiredRefresh(config.ExpiredRefreshToken))
	}
	opts = append(opts, jwtAuth.WithSigningKeyRefresh([]byte(config.SigningRefreshKey)))
	return jwtAuth.NewJWTAuth(opts...), nil
}