package config

var (
	SigningKey   string   = "access"
	Expired     int = 900
	SigningRefreshKey string = "refresh"
	ExpiredRefreshToken     int    = 0
)
