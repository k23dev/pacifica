package models

type CommandFile struct {
	Commands [][]struct {
		Name    string `json:"name"`
		Path    string `json:"path"`
		IsSudo  bool   `json:"is_sudo"`
		Comment string `json:"comment"`
		Args    []struct {
			Method    string `json:"method"`
			InputType string `json:"input_type"`
			IsFlag    bool   `json:"is_flag"`
			Order     int    `json:"order"`
			Comment   string `json:"comment"`
		} `json:"args"`
	} `json:"commands"`
}
