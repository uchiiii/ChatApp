package main

type Config struct {
	Auth AuthConfig
}

type AuthConfig struct {
	SecurityKey string
	Each map[string]Each
}

type Each struct {
	Id string 
	Secret string
	RedirectURL string
}

