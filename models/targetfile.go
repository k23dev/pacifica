package models

type TargetFile struct {
	Targets [][]struct {
		Name   string `json:"name"`
		Fields []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"fields"`
	} `json:"targets"`
}
