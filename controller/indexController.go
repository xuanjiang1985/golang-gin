package ctrl

import (
	seelog "github.com/cihub/seelog"
	"github.com/flosch/pongo2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang-gin/csrf"
	"gopkg.in/gin-gonic/gin.v1"
	//"log"
	"math"
	"net/http"
	"strconv"
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

	p := []ArticlesUser{}

	//if has page param for articles
	page := c.Query("page")
	var skip int
	var current_page int
	if page == "" {
		skip = 0
		current_page = 1
	} else {
		b, ok := strconv.Atoi(page)
		if ok != nil || b < 1 {
			c.String(404, "404 page not found")
			return
		}
		skip = b*10 - 10
		current_page = b
	}
	//sql select 10 articles
	err = db.Select(&p, "SELECT articles.id as id,articles.user_id as user_id,articles.thanks as thanks,articles.comments as comments,articles.content as content,FROM_UNIXTIME(articles.created_at, '%Y-%m-%d %H:%i') as created_at, users.name as user_name, users.sex as sex FROM articles left join users on articles.user_id = users.id ORDER BY id DESC LIMIT ?,10", skip)
	if err != nil {
		seelog.Error("can't read db ", err)
		return
	}
	//find all pages
	var all int
	err = db.Get(&all, "SELECT count(*) FROM articles")
	if err != nil {
		seelog.Error("can't read db ", err)
		return
	}
	all_page := float64(all) / float64(10)
	allpage := math.Ceil(all_page)
	all = int(allpage)

	csrfToken := csrf.GetToken(c)
	//if auth
	authUser, _ := c.Get("authUser")
	//log.Println(authUser)
	c.HTML(http.StatusOK, "index.html", pongo2.Context{
		"token":        csrfToken,
		"articles":     &p,
		"current_page": current_page,
		"all_page":     all,
		"authUser":     authUser,
	})
}
