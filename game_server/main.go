package main

import (
	"game_server/conf"
	"game_server/game"
	"game_server/gate"
	"game_server/login"
	"github.com/leaf"
	lconf "github.com/leaf/conf"
)

func main(){
	lconf.LogLevel = "debug"
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}

