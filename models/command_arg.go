package models

import (
	"fmt"
	"strconv"

	"github.com/k23dev/go4it/go4error"
	"gorm.io/gorm"
)

type Command_arg struct {
	gorm.Model
	CommandID  uint
	Command    Command
	CategoryID uint
	Category   Category
	Method     string
	InputType  string
	IsFlag     bool
	Comment    string
	Order      uint
}

type Command_argDTO struct {
	CommandID  uint `json:"Command_id" param:"Command_id" query:"Command_id" form:"Command_id"`
	Command    Command
	CategoryID uint `json:"category_id" param:"category_id" query:"category_id" form:"category_id"`
	Category   Category
	Method     string `json:"method" param:"method" query:"method" form:"method"`
	InputType  string `json:"input_type" param:"input_type" query:"input_type" form:"input_type"`
	IsFlag     string `json:"is_flag" param:"is_flag" query:"is_flag" form:"is_flag"`
	IsFlagBool bool
	Comment    string `json:"comment" param:"comment" query:"comment" form:"comment"`
	Order      uint   `json:"order" param:"order" query:"order" form:"order"`
}

type Command_argCounter struct {
	Total int
}

func NewCommand_arg() *Command_arg {
	return &Command_arg{}
}

func (s *Command_arg) Count(db *gorm.DB) (int, error) {
	counter := &Command_argCounter{}
	db.Model(&Command_arg{}).Select("count(ID) as total").Find(&counter)
	return counter.Total, nil
}

func (s *Command_arg) FindOne(db *gorm.DB, id int) (*Command_arg, error) {
	var command_arg Command_arg
	db.First(&command_arg, id)
	if command_arg.ID == 0 {
		return nil, &go4error.ModelError{
			ModelName: "Command_arg",
			Code:      0,
			Message:   go4error.MsgIDNotFound(id),
		}
	}
	return &command_arg, nil
}

func (s *Command_arg) FindAll(db *gorm.DB) ([]Command_arg, error) {
	var command_args []Command_arg
	db.Order("created_at ASC").Find(&command_args)
	if len(command_args) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Command_arg",
			Code:      0,
			Message:   go4error.MsgZeroRecordsFound(),
		}
	}
	return command_args, nil
}
func (s *Command_arg) FindAllFiltered(db *gorm.DB, CommandID uint) ([]Command_arg, error) {
	var command_args []Command_arg
	db.Order("created_at ASC").Where("Command_id=?", CommandID).Find(&command_args)
	if len(command_args) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Command_arg",
			Code:      0,
			Message:   go4error.MsgZeroRecordsFound(),
		}
	}
	return command_args, nil
}

func (s *Command_arg) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Command_arg, error) {
	command_args := []Command_arg{}

	db.Order("created_at ASC").Preload("Command").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&command_args)
	if len(command_args) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Command_arg",
			Code:      0,
			Message:   go4error.MsgZeroRecordsFound(),
		}
	}
	return &command_args, nil
}

func (s *Command_arg) Create(db *gorm.DB, dto Command_argDTO) (*Command_arg, error) {
	s.SatinizeDTOCreate(&dto)
	command_arg := Command_arg{
		CommandID:  dto.CommandID,
		CategoryID: dto.CategoryID,
		Method:     dto.Method,
		InputType:  dto.InputType,
		IsFlag:     dto.IsFlagBool,
		Comment:    dto.Comment,
		Order:      dto.Order,
	}

	db.Create(&command_arg)
	return &command_arg, nil
}

func (s *Command_arg) Update(db *gorm.DB, id int, dto Command_argDTO) (*Command_arg, error) {
	s.SatinizeDTOUpdate(&dto)
	db.Model(&Command_arg{}).Where("ID =?", id).Update("method", dto.Method).Update("category_id", dto.CategoryID).Update("Command_id", dto.CommandID).Update("input_type", dto.InputType).Update("is_flag", dto.IsFlagBool).Update("comment", dto.Comment).Update("order", dto.Order)
	return s, nil
}

func (s *Command_arg) Delete(db *gorm.DB, id int) (*Command_arg, error) {
	command_arg, err := s.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&command_arg)
	return command_arg, nil
}

func (s *Command_arg) GetIDAsString() string {
	return fmt.Sprintf("%d", s.ID)
}

func (s *Command_arg) SatinizeDTOCreate(dto *Command_argDTO) error {
	flagb, _ := strconv.ParseBool(dto.IsFlag)
	dto.IsFlagBool = flagb
	return nil
}

func (s *Command_arg) SatinizeDTOUpdate(dto *Command_argDTO) error {
	flagb, _ := strconv.ParseBool(dto.IsFlag)
	dto.IsFlagBool = flagb
	return nil
}

func (s *Command_arg) FindAllByCommandID(db *gorm.DB, id uint) (*[]Command_arg, error) {

	list := &[]Command_arg{}
	db.Model(&Command_arg{}).Preload("Command").Where("Command_id=?", id).Find(&list)
	if len(*list) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Command_arg",
			Code:      0,
			Message:   go4error.MsgZeroRecordsFound(),
		}
	}
	return list, nil

}
