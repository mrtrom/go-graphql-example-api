package config

type strContextKey string

// Constants to be managed in main with Context
const (
	CTXUserService = strContextKey("userService")
	CTXChatService = strContextKey("chatService")
	CTXLog         = strContextKey("log")
	CTXConfig      = strContextKey("config")
)
