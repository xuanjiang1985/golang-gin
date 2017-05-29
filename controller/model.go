package ctrl

import (
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

//comments in table
type Comments struct {
	Id         int
	Article_id int
	User_id    int
	Comment    string
	Created_at string
	Updated_at string
}
