package app

import (
	"fmt"

	"github.com/k23dev/go4it/interact"
	"github.com/k23dev/pacifica/models"
)

func ReadTargetsFile(filepath string) {
	targets := &models.TargetFile{}

	if filepath == "" {
		filepath = "./_data/targets.json"
	}

	interact.ReadAndParseJson(filepath, targets)

	fmt.Printf("%+v\n", targets)
}

func ReadCommandsFile(filepath string) {
	commands := &models.CommandFile{}

	if filepath == "" {
		filepath = "./_data/commands.json"
	}

	interact.ReadAndParseJson(filepath, commands)

	fmt.Printf("%+v\n", commands)
}

func ReadDirectivesFile(filepath string) {
	commands := &models.CommandFile{}

	if filepath == "" {
		filepath = "./_data/commands.json"
	}

	interact.ReadAndParseJson(filepath, commands)

	fmt.Printf("%+v\n", commands)
}
