package connection

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

const (
	engine   = "postgres"
	username = "postgres"
	password = "root"
	name     = "restful_app"
)

func GetConnection() *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, name)
	db, err := gorm.Open(engine, dsn)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
