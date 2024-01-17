package models

import (
	"fmt"

	"github.com/k23dev/pacifica/pkg/go4error.ModelError"
	"github.com/k23dev/pacifica/pkg/tango_debug"
	"gorm.io/gorm"
)

type Target_field struct {
	gorm.Model
	Name    string `json:"name" param:"name" query:"name" form:"name"`
	FValue  string `json:"fvalue" param:"fvalue" query:"fvalue" form:"fvalue"`
	TangaID uint
	Tanga   Target
}

type Target_fieldDTO struct {
	Name    string `json:"name" param:"name" query:"name" form:"name"`
	FValue  string `json:"fvalue" param:"fvalue" query:"fvalue" form:"fvalue"`
	TangaID uint   `json:"tanga_id" param:"tanga_id" query:"tanga_id" form:"tanga_id"`
}

type Target_fieldCounter struct {
	Total int
}

type Target_fieldFile struct {
	Name  string `toml:"name" json:"name"`
	Value string `toml:"value" json:"value"`
}

func NewTarget_field() *Target_field {
	return &Target_field{}
}

func (t *Target_field) Count(db *gorm.DB) (int, error) {
	counter := &Target_fieldCounter{}
	db.Model(&Target_field{}).Select("count(ID) as total").Find(&counter)
	return counter.Total, nil
}

func (t *Target_field) FindOne(db *gorm.DB, id int) (*Target_field, error) {
	var tanga_field Target_field
	db.Preload("Tanga").First(&tanga_field, id)
	if tanga_field.ID == 0 {
		return nil, &go4error.ModelError{
			ModelName: "Tanga_field",
			Code:      0,
			Message:   go4error.ModelError.MsgIDNotFound(id),
		}
	}
	tango_debug.Struct("tanga field", tanga_field)
	return &tanga_field, nil
}

func (t *Target_field) FindAllFiltered(db *gorm.DB, tangaID uint) (*Target_field, error) {
	var tanga_field Target_field
	db.Preload("Tanga").Where("tanga_id=?", tangaID).First(&tanga_field)
	if tanga_field.ID == 0 {
		return nil, &go4error.ModelError{
			ModelName: "Tanga_field",
			Code:      0,
			Message:   go4error.ModelError.MsgIDNotFound(int(tangaID)),
		}
	}
	tango_debug.Struct("tanga field", tanga_field)
	return &tanga_field, nil
}

func (t *Target_field) FindAll(db *gorm.DB) ([]Target_field, error) {
	var tanga_fields []Target_field
	db.Order("created_at ASC").Preload("Tanga").Find(&tanga_fields)
	if len(tanga_fields) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Tanga_field",
			Code:      0,
			Message:   go4error.ModelError.MsgZeroRecordsFound(),
		}
	}
	return tanga_fields, nil
}

func (t *Target_field) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Target_field, error) {
	tanga_fields := []Target_field{}

	db.Order("created_at ASC").Preload("Tanga").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&tanga_fields)
	if len(tanga_fields) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Tanga_field",
			Code:      0,
			Message:   go4error.ModelError.MsgZeroRecordsFound(),
		}
	}
	return &tanga_fields, nil
}

func (t *Target_field) Create(db *gorm.DB, dto Target_fieldDTO) (*Target_field, error) {
	t.SatinizeDTOCreate(&dto)
	tango_debug.Struct("dto", dto)
	tanga_field := Target_field{
		Name:    dto.Name,
		FValue:  dto.FValue,
		TangaID: dto.TangaID,
	}
	db.Create(&tanga_field)
	return &tanga_field, nil
}

func (t *Target_field) Update(db *gorm.DB, id int, dto Target_fieldDTO) (*Target_field, error) {
	t.SatinizeDTOUpdate(&dto)
	tango_debug.Struct("dto", dto)
	db.Model(&Target_field{}).Where("ID =?", id).Update("name", dto.Name).Update("fvalue", dto.FValue).Update("tanga_id", dto.TangaID)
	return t, nil
}

func (t *Target_field) Delete(db *gorm.DB, id int) (*Target_field, error) {
	tanga_field, err := t.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&tanga_field)
	return tanga_field, nil
}

func (t *Target_field) GetIDAsString() string {
	return fmt.Sprintf("%d", t.ID)
}

func (t *Target_field) SatinizeDTOCreate(dto *Target_fieldDTO) error {
	// TODO
	return nil
}

func (t *Target_field) SatinizeDTOUpdate(dto *Target_fieldDTO) error {
	// TODO
	return nil
}

func (s *Target_field) FindAllByTangaFieldID(db *gorm.DB, id uint) (*[]Target_field, error) {

	list := &[]Target_field{}
	db.Model(&Target_field{}).Preload("Tanga").Where("tanga_id=?", id).Find(&list)
	if len(*list) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Tanga_field",
			Code:      0,
			Message:   go4error.ModelError.MsgZeroRecordsFound(),
		}
	}
	return list, nil

}
