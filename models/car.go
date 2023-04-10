package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Car struct {
	GormModel
	Title       string `json:"title" form:"title" valid:"required"`
	Brand       string `json:"brand" form:"brand" valid:"required"`
	Model       string `json:"model" form:"model" valid:"required"`
	Description string `json:"description" form:"description" valid:"required"`
	UserID      uint
	User        *User
}

func (c *Car) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return

}

func (c *Car) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
