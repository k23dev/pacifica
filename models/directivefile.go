package models

type DirectiveFile struct {
	Directives [][]struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		TargetID     int    `json:"target_id"`
		TargetField  int    `json:"target_field"`
		CommandID    int    `json:"command_id"`
		CommandArgID int    `json:"command_arg_id"`
	} `json:"directives"`
}
