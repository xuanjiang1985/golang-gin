package ctrl

import (
	"database/sql"
	"golang-gin/conf"
)

//db and log settings
var sqlconn string = conf.Conn
var logger = conf.Logger

//articles in table
type Articles struct {
	Id         int
	User_id    int
	Thanks     int
	Comments   int
	Content    string
	Created_at string
	Updated_at string
}

//articles with UserName in table
type ArticlesUser struct {
	Id         int
	User_id    int
	Thanks     int
	Comments   int
	Content    string
	Created_at string
	User_name  sql.NullString
	Sex        sql.NullInt64
}

//comments in table
type Comments struct {
	Id         int
	Article_id int
	User_id    int
	Comment    string
	Created_at string
	Updated_at string
}

//comments in table
type Users struct {
	Id             int
	Name           string
	Email          string
	Password       string
	Remember_token string
	Sex            int
	Admin          int
	Created_at     int
	Updated_at     int
}
