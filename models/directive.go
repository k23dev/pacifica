package models

import (
	"fmt"

	"github.com/k23dev/pacifica/pkg/go4error.ModelError"
	"gorm.io/gorm"
)

type Directive struct {
	gorm.Model
	Name string
}
type DirectiveDTO struct {
	Name string `json:"name" param:"name" query:"name" form:"name"`
}

type DirectiveCounter struct {
	Total int
}

func NewDirective() *Directive {
	return &Directive{}
}

func (d *Directive) Count(db *gorm.DB) (int, error) {
	counter := &DirectiveCounter{}
	db.Model(&Directive{}).Select("count(ID) as total").Find(&counter)
	return counter.Total, nil
}

func (d *Directive) FindOne(db *gorm.DB, id int) (*Directive, error) {
	var directive Directive
	db.First(&directive, id)
	if directive.ID == 0 {
		return nil, &go4error.ModelError{
			ModelName: "Directive",
			Code:      0,
			Message:   go4error.ModelError.MsgIDNotFound(id),
		}
	}
	return &directive, nil
}

func (d *Directive) FindAll(db *gorm.DB) ([]Directive, error) {
	var directives []Directive
	db.Order("created_at ASC").Find(&directives)
	if len(directives) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Directive",
			Code:      0,
			Message:   go4error.ModelError.MsgZeroRecordsFound(),
		}
	}
	return directives, nil
}

func (d *Directive) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Directive, error) {
	directives := []Directive{}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&directives)
	if len(directives) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Directive",
			Code:      0,
			Message:   go4error.ModelError.MsgZeroRecordsFound(),
		}
	}
	return &directives, nil
}

func (d *Directive) Create(db *gorm.DB, dto DirectiveDTO) (*Directive, error) {
	d.SatinizeDTOCreate(&dto)
	directive := Directive{
		Name: dto.Name,
	}
	db.Create(&directive)
	return &directive, nil
}

func (d *Directive) Update(db *gorm.DB, id int, dto DirectiveDTO) (*Directive, error) {
	d.SatinizeDTOUpdate(&dto)
	db.Model(&Directive{}).Where("ID =?", id).Update("name", dto.Name)
	return d, nil
}

func (d *Directive) Delete(db *gorm.DB, id int) (*Directive, error) {
	directive, err := d.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&directive)
	return directive, nil
}

func (d *Directive) GetIDAsString() string {
	return fmt.Sprintf("%d", d.ID)
}

func (d *Directive) SatinizeDTOCreate(dto *DirectiveDTO) error {
	// TODO
	return nil
}

func (d *Directive) SatinizeDTOUpdate(dto *DirectiveDTO) error {
	// TODO
	return nil
}
