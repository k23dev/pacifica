package models

import (
	"fmt"

	"github.com/k23dev/pacifica/pkg/go4error.ModelError"
	"gorm.io/gorm"
)

type Target struct {
	gorm.Model
	Name    string `json:"name" param:"name" query:"name" form:"name"`
	Comment string `json:"comment" param:"comment" query:"comment" form:"comment"`
}

type TargetDTO struct {
	Name    string `json:"name" param:"name" query:"name" form:"name"`
	Comment string `json:"comment" param:"comment" query:"comment" form:"comment"`
}

type TargetCounter struct {
	Total int
}

type TargetFile struct {
	Target map[string][]Target_fieldFile `toml:"target" json:"target"`
}

func NewTarget() *Target {
	return &Target{}
}

func (t *Target) Count(db *gorm.DB) (int, error) {
	counter := &TargetCounter{}
	db.Model(&Target{}).Select("count(ID) as total").Find(&counter)
	return counter.Total, nil
}

func (t *Target) FindOne(db *gorm.DB, id int) (*Target, error) {
	var tanga Target
	db.First(&tanga, id)
	if tanga.ID == 0 {
		return nil, &go4error.ModelError{
			ModelName: "Tanga",
			Code:      0,
			Message:   go4error.ModelError.MsgIDNotFound(id),
		}
	}
	return &tanga, nil
}

func (t *Target) FindAll(db *gorm.DB) ([]Target, error) {
	var tangas []Target
	db.Order("created_at ASC").Find(&tangas)
	if len(tangas) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Tanga",
			Code:      0,
			Message:   go4error.ModelError.MsgZeroRecordsFound(),
		}
	}
	return tangas, nil
}

func (t *Target) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Target, error) {
	tangas := []Target{}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&tangas)
	if len(tangas) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Tanga",
			Code:      0,
			Message:   go4error.ModelError.MsgZeroRecordsFound(),
		}
	}
	return &tangas, nil
}

func (t *Target) Create(db *gorm.DB, name string, comment string) (*Target, error) {
	tanga := Target{
		Name:    name,
		Comment: comment,
	}
	db.Create(&tanga)
	return &tanga, nil
}

func (t *Target) Update(db *gorm.DB, id int, name string, comment string) (*Target, error) {
	db.Model(&Target{}).Where("ID =?", id).Update("name", name).Update("comment", comment)
	return t, nil
}

func (t *Target) Delete(db *gorm.DB, id int) (*Target, error) {
	tanga, err := t.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&tanga)
	return tanga, nil
}

func (t *Target) GetIDAsString() string {
	return fmt.Sprintf("%d", t.ID)
}
