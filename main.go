package main

import (
	"github.com/k23dev/go4it"
	"github.com/k23dev/pacifica/app"
	"github.com/k23dev/pacifica/cmd"
)

func main() {

	configfile := "./config/appconfig"
	pacifica := go4it.NewApp(configfile)

	pacifica.Connect2Db("local")
	pacifica.DB.SetPrimaryDB(0)

	menu := cmd.NewCmdMenu()

	app.ReadTargetFile()

	//cmd.Banner(pacifica.Config.App_name, app.Config.App_version)
	//menu.ShowMenu()
}
