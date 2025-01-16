package gosqllearn

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() (*gorm.DB, error) {
	dsn := mysql.Open("root:@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dsn, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

var db, err = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db) // Assert that the DB connection is not nil
}
func TestExecuteSQL(t *testing.T) {
	err := db.Exec("insert into sample(id, name) values (?, ?)", "1", "Tiara").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values (?, ?)", "2", "Budi").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values (?, ?)", "1", "Raisa").Error
	assert.Nil(t, err)
}
