package cmd

import "fmt"

type CmdMenu struct {
	Options        []string
	OptionSelected int
	OptionPrev     int
}

func NewCmdMenu() *CmdMenu {
	menu := &CmdMenu{}

	menu.fillMenu()

	return menu
}

func (m *CmdMenu) fillMenu() {
	m.Options = append(m.Options, "Inicio")
	m.Options = append(m.Options, "Targets")
	m.Options = append(m.Options, "Targets")
	m.Options = append(m.Options, "Directivas")
	m.Options = append(m.Options, "Ejecutar")
	m.Options = append(m.Options, "Logs")
}

func (m *CmdMenu) ShowMenu() {
	for k, v := range m.Options {

		format := fmt.Sprintf("%d - %s", k, v)
		fmt.Println(format)

	}
}
