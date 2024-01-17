package app

import (
	"fmt"

	"github.com/k23dev/go4it/interact"
	"github.com/k23dev/pacifica/models"
)

func ReadTargetFile(file string) {
	targets := &models.TargetFile{}
	interact.ReadAndParseToml("./_data/targets", targets)
	fmt.Printf("%+v", targets)
}
