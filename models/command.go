package models

import (
	"fmt"

	"github.com/k23dev/go4it/go4error"
	"gorm.io/gorm"
)

type Command struct {
	gorm.Model
	Name    string
	Path    string
	IsSudo  bool
	Comment string
}

type CommandDTO struct {
	Name    string `json:"name" param:"name" query:"name" form:"name"`
	Path    string `json:"path" param:"path" query:"path" form:"path"`
	IsSudo  bool   `json:"is_sudo" param:"is_sudo" query:"is_sudo" form:"is_sudo"`
	Comment string `json:"comment" param:"comment" query:"comment" form:"comment"`
}

type CommandCounter struct {
	Total int
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Count(db *gorm.DB) (int, error) {
	counter := &CommandCounter{}
	db.Model(&Command{}).Select("count(ID) as total").Find(&counter)
	fmt.Printf("%v", counter.Total)
	return counter.Total, nil
}

func (c *Command) FindOne(db *gorm.DB, id int) (*Command, error) {
	var Command Command
	db.First(&Command, id)
	if Command.ID == 0 {
		return nil, &go4error.ModelError{
			ModelName: "Command",
			Code:      0,
			Message:   go4error.MsgIDNotFound(id),
		}
	}
	return &Command, nil
}

func (c *Command) FindAll(db *gorm.DB) ([]Command, error) {
	var commands []Command
	db.Order("created_at ASC").Find(&commands)
	if len(commands) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Command",
			Code:      0,
			Message:   go4error.MsgZeroRecordsFound(),
		}
	}
	return commands, nil
}

func (c *Command) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Command, error) {
	commands := []Command{}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&commands)
	if len(commands) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Command",
			Code:      0,
			Message:   go4error.MsgZeroRecordsFound(),
		}
	}
	return &commands, nil
}

func (c *Command) Create(db *gorm.DB, cDTO CommandDTO) (*Command, error) {
	command := Command{
		Name:    cDTO.Name,
		Path:    cDTO.Path,
		IsSudo:  cDTO.IsSudo,
		Comment: cDTO.Comment,
	}
	db.Create(&command)
	return &command, nil
}

func (c *Command) Update(db *gorm.DB, id int, name string, path string, isSudo bool, comment string) (*Command, error) {
	db.Model(&Command{}).Where("ID =?", id).Update("name", name).Update("path", path).Update("is_sudo", isSudo).Update("comment", comment)
	return c, nil
}

func (c *Command) Delete(db *gorm.DB, id int) (*Command, error) {
	command, err := c.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&command)
	return command, nil
}

func (c *Command) GetIDAsString() string {
	return fmt.Sprintf("%d", c.ID)
}

func (c *Command) GetIsSudoAsString() string {
	return fmt.Sprintf("%v", c.IsSudo)
}
