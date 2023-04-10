package models

import (
	"chapter3_2/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string `gorm:"not null" json:"full_name" form:"full_name" valid:"required"`
	Email    string `gorm:"not null" json:"email" form:"email" valid:"required"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required,minstringlength(6)"`
	Cars     []Car  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"cars"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
