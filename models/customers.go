package models

import (
	"gorm.io/gorm"
)

type Customers struct {
	ID    uint64 `gorm:"primary key"`
	Name  string `gorm:"type:varchar" json:"name"`
	Age   uint64 `gorm:"type:int" json:"age"`
	Email string `gorm:"type:varchar" json:"email"`
}

func MigrateCustomerData(db *gorm.DB) error {
	err := db.AutoMigrate(&Customers{})
	return err
}
