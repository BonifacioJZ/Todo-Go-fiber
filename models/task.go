package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	*gorm.Model

	Name        string `json:"name"  gorm:"type:varchar(200);size:200"`
	Description string `json:"description" gorm:"type:text"`
	Done        bool   `json:"done" gorm:"type:boolean"`
}
