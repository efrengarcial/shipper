package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewRandom()
	return scope.SetColumn("Id", uuid.String())
}