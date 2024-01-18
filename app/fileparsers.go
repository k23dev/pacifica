package app

import (
	"github.com/k23dev/go4it/interact"
	"github.com/k23dev/pacifica/models"
)

func ReadTargetsFile(filepath string) *models.TargetFile {
	targets := &models.TargetFile{}

	if filepath == "" {
		filepath = "./_data/targets.json"
	}

	interact.ReadAndParseJson(filepath, targets)

	return targets
}

func ReadCommandsFile(filepath string) *models.CommandFile {
	commands := &models.CommandFile{}

	if filepath == "" {
		filepath = "./_data/commands.json"
	}

	interact.ReadAndParseJson(filepath, commands)

	return commands
}

func ReadDirectivesFile(filepath string) *models.DirectiveFile {
	directives := &models.DirectiveFile{}

	if filepath == "" {
		filepath = "./_data/directives.json"
	}

	interact.ReadAndParseJson(filepath, directives)

	return directives
}
