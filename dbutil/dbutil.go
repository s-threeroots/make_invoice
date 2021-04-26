package dbutil

import (
	"make_invoice/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var drvier = "postgres"
var connect = "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"

func Open() *gorm.DB {

	db, err := gorm.Open(drvier, connect)
	if err != nil {
		panic("db error")
	}
	return db
}

func InitDb() {

	db := Open()
	db.AutoMigrate(model.Invoice{}, model.Detail{}, model.Item{})
	defer db.Close()
}

func Insert(value interface{}) error {

	db := Open()
	defer db.Close()

	result := db.Create(value)
	err := result.Error

	return err

}
