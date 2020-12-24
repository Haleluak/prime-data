package jwt

import "github.com/dgrijalva/jwt-go"

type options struct {
	signingMethod     jwt.SigningMethod
	signingKey        interface{}
	keyFunc           jwt.Keyfunc
	expired           int
	tokenType         string
	keyFuncRefresh    jwt.Keyfunc
	expiredRefresh    int
	signingRefreshKey interface{}
}

type Option func(*options)

func WithExpired(expired int) Option {
	return func(o *options) {
		o.expired = expired
	}
}

func WithKeyFunc(keyFunc jwt.Keyfunc) Option {
	return func(o *options) {
		o.keyFunc = keyFunc
	}
}

func WithSigningKey(key interface{}) Option {
	return func(o *options) {
		o.signingKey = key
	}
}

func WithExpiredRefresh(expired int) Option {
	return func(o *options) {
		o.expiredRefresh = expired
	}
}

func WithKeyFuncRefresh(keyFunc jwt.Keyfunc) Option {
	return func(o *options) {
		o.keyFuncRefresh = keyFunc
	}
}

func WithSigningKeyRefresh(key interface{}) Option {
	return func(o *options) {
		o.signingRefreshKey = key
	}
}