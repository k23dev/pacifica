package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name" param:"name" query:"name" form:"name"`
}

type CategoryDTO struct {
	Name string `json:"name" param:"name" query:"name" form:"name"`
}

type CategoryCounter struct {
	Total int
}

func NewCategory() *Category {
	return &Category{}
}

func (c *Category) Count(db *gorm.DB) (int, error) {
	counter := &CategoryCounter{}
	db.Model(&Category{}).Select("count(ID) as total").Find(&counter)
	return counter.Total, nil
}

func (c *Category) FindOne(db *gorm.DB, id int) (*Category, error) {
	var category Category
	db.First(&category, id)
	if category.ID == 0 {
		return nil, &go4error.ModelError{
			ModelName: "Category",
			Code:      0,
			Message:   go4error.MsgIDNotFound(id),
		}
	}
	return &category, nil
}

func (c *Category) FindAll(db *gorm.DB) ([]Category, error) {
	var categories []Category
	db.Order("created_at ASC").Find(&categories)
	if len(categories) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Category",
			Code:      0,
			Message:   go4error.ModelError.MsgZeroRecordsFound(),
		}
	}
	return categories, nil
}

func (c *Category) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Category, error) {
	categories := []Category{}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&categories)
	if len(categories) <= 0 {
		return nil, &go4error.ModelError{
			ModelName: "Category",
			Code:      0,
			Message:   go4error.ModelError.MsgZeroRecordsFound(),
		}
	}
	return &categories, nil
}

func (c *Category) Create(db *gorm.DB, name string) (*Category, error) {
	category := Category{
		Name: name,
	}
	db.Create(&category)
	return &category, nil
}

func (c *Category) Update(db *gorm.DB, id int, name string) (*Category, error) {
	db.Model(&Category{}).Where("ID =?", id).Update("name", name)
	return c, nil
}

func (c *Category) Delete(db *gorm.DB, id int) (*Category, error) {
	category, err := c.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&category)
	return category, nil
}

func (c *Category) GetIDAsString() string {
	return fmt.Sprintf("%d", c.ID)
}
