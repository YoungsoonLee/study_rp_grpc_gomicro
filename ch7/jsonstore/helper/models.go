package helper

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Shipment struct {
	gorm.Model
	Package []Package
	Data    string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

type Package struct {
	gorm.Model
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

func (Shipment) TableName() string {
	return "Shipment"
}

func (Package) TableName() string {
	return "Package"
}

func InitDB() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open("postgres", "postgres://ivsxlgmu:TAoE09ZolUGIolTkZIeI-jAifWBkvvjy@ruby.db.elephantsql.com:5432/ivsxlgmu")
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Shipment{}, &Package{})
	return db, nil
}
