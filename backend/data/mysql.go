package data

import (
	"fmt"
	"proyecto-integrador/models"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var DB *xorm.Engine

var (
	userName string = "root"
	password string = "password"
	address  string = "localhost"
	port     string = "3306"
	dbName   string = "db"
	charset  string = "utf8"
)

// ConnectDatabase connects to mysql db, gets and sets table schema
func ConnectDatabase() (err error) {
	// connect to database
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", userName, password, address, port, dbName, charset)
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if nil != err {
		fmt.Println("sql connection error")
		return err
	}
	fmt.Println("sql connected successfully")

	// set show SQL query in console
	engine.ShowSQL(true)

	// set name mapping to GonicMapper
	engine.SetMapper(names.GonicMapper{})

	// create or sincronize book table with struct
	// it is possible to sync more than one table
	// example: err := engine.Sync(new(models.Book), new(models.Customer))
	if err := engine.Sync2(new(models.Categories)); nil != err {
		fmt.Println("error with database schema synchronize")
	}

	// show all the tables in console
	ts, _ := engine.DBMetas()
	for _, v := range ts {
		fmt.Printf("v: %v\n", v.Name)
	}

	DB = engine
	return nil
}
