package main

import (
	"fmt"

	"github.com/k23dev/go4it"
	"github.com/k23dev/pacifica/app"
)

func main() {

	configfile := "./config/appconfig"
	pacifica := go4it.NewApp(configfile)

	pacifica.Connect2Db("local")
	pacifica.DB.SetPrimaryDB(0)

	//menu := cmd.NewCmdMenu()

	// targets := app.ReadTargetsFile("")
	commands := app.ReadCommandsFile("")
	// directives := app.ReadDirectivesFile("")

	// fmt.Printf("%+v \n", targets)
	// fmt.Printf("%+v \n", commands)
	fmt.Printf("%+v \n", commands.Command[0][1].Name)
	fmt.Printf("%+v \n", commands.Command[0][1].Path)
	// fmt.Printf("%+v \n", directives)

	//cmd.Banner(pacifica.Config.App_name, app.Config.App_version)
	//menu.ShowMenu()
}
