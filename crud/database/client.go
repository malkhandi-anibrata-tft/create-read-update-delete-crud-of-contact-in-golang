package database

import (
	"crud/entity"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {

		fmt.Println("error in connect database", err)
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}
func Migrate(table *entity.Person) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
