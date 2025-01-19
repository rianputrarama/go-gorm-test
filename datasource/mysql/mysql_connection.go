package mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Connect *gorm.DB

func Init() {
	//start := time.Now()
	var err error
	dsn := mysql.Open("root:@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local")
	//db, err := gorm.Open(dsn, &gorm.Config{})
	Connect, err = gorm.Open(dsn, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})

	if err != nil {
		log.Fatal("failed to initialize connection to mysq")
	} else {
		fmt.Println("initialize connection to mysql, connection success")
	}
}
