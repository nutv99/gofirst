package orm

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Password string
	FullName string
	Avatar   string
}
