package utils

import "github.com/h4rimu/kaspro-sdk/config"

func GetRunMode() string {
	serverMode := config.MustGetString("server.mode")
	return serverMode
}
