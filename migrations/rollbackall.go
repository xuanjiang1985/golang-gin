package main

import (
	"database/sql"
	"fmt"
	"github.com/DavidHuie/gomigrate"
	_ "github.com/go-sql-driver/mysql"
	"golang-gin/conf"
)

var sqlconn string = conf.Conn

func main() {
	//connect mysql
	db, err := sql.Open("mysql", sqlconn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//starting migration
	migrator, _ := gomigrate.NewMigrator(db, gomigrate.Mysql{}, "mysql/")
	err = migrator.RollbackAll()
	//err = migrator.Migrate()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("RollBackAll successfully")
}
