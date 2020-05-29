package config

type strContextKey string

// Constants to be managed in main with Context
const (
	CTXUserService = strContextKey("userService")
	CTXLog         = strContextKey("log")
	CTXConfig      = strContextKey("config")
)
