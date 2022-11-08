package config

import "github.com/spf13/viper"

const (
	DebugModeKey = "rcon-client.app.debug"
)

func DebugMode() bool {
	return viper.GetBool(DebugModeKey)
}
