package ctrl

import (
	seelog "github.com/cihub/seelog"
	"github.com/flosch/pongo2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang-gin/csrf"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type IndexController struct {
}

func (ct *IndexController) Get(c *gin.Context) {
	//开启日志
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()

	//数据库连接
	db, err := sqlx.Connect("mysql", sqlconn)
	if err != nil {
		seelog.Error("can't connect db ", err)
		return
	}
	defer db.Close()

	p := []Articles{}
	err = db.Select(&p, "SELECT * FROM articles ORDER BY id DESC")
	if err != nil {
		seelog.Error("can't read db ", err)
		return
	}
	csrfToken := csrf.GetToken(c)
	c.HTML(http.StatusOK, "index.html", pongo2.Context{"token": csrfToken, "articles": &p})
}
