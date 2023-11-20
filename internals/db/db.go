package db

import (
	"log"

	"github.com/ozancakir/go-pokeapi-proxy/internals/db/entities"
	"github.com/ozancakir/go-pokeapi-proxy/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _db *gorm.DB

func GetDB() *gorm.DB {
	return _db
}

func Setup() {
	var err error
	_db, err = gorm.Open(sqlite.Open(utils.Path("./pokeapi.db")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalln("failed to connect database")
	}

	err = _db.AutoMigrate(
		&entities.Response{},
	)
	if err != nil {
		log.Fatalln("failed to migrate database")
	}

}
