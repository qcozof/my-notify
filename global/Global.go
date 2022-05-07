package global

import (
	"github.com/bwmarrin/discordgo"
	"github.com/qcozof/my-notify/config"
	"github.com/spf13/viper"
)

var (
	SERVER_CONFIG config.ServerConfig
	VIPER         *viper.Viper
	IN_ARG        string
	DISCORD       *discordgo.Session
)
