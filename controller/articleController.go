package ctrl

import (
	seelog "github.com/cihub/seelog"
	"github.com/flosch/pongo2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang-gin/sessions"
	"gopkg.in/gin-gonic/gin.v1"
	//"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

type ArticleController struct {
}

func (ct *ArticleController) Get(c *gin.Context) {
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
	err = db.Select(&p, "SELECT * FROM articles")
	if err != nil {
		seelog.Error("can't read db ", err)
		return
	}
	seelog.Debug(&p)
	c.HTML(http.StatusOK, "article.html", pongo2.Context{"data": &p})
}

func (ct *ArticleController) Store(c *gin.Context) {
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
	var userId string
	session := sessions.Default(c)
	if s, ok := session.Get("authUserId").(string); ok && len(s) != 0 {
		//log.Println(session.Get("authUserName"))
		userId = s
	} else {
		userId = "0"
	}
	content := c.PostForm("content")
	unix_time := time.Now().Unix()
	_, err = db.Exec(`INSERT INTO articles (user_id,content,created_at,updated_at) VALUES (?,?,?,?)`, userId, content, unix_time, unix_time)
	if err != nil {
		c.String(200, "提交失败。")
		//log.Println(err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

func (ct *ArticleController) AddThank(c *gin.Context) {
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
	article_id := c.Param("id")
	ip := c.ClientIP()
	var haveId int
	err = db.Get(&haveId, "SELECT id FROM thanks WHERE article_id=? AND ip=?", article_id, ip)
	if err != nil {
		db.MustExec(`INSERT INTO thanks (article_id,ip) VALUES (?,?)`, article_id, ip)
		db.MustExec(`UPDATE articles SET thanks=thanks+1 WHERE id=?`, article_id)
	}
	return
}

func (ct *ArticleController) AddComment(c *gin.Context) {
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
	comment := c.PostForm("comment")
	article_id := c.PostForm("article_id")
	unix_time := time.Now().Unix()
	natural_time := time.Unix(unix_time, 0)
	natural_time_str := natural_time.Format("2006-01-02 03:04")
	db.MustExec(`INSERT INTO comments (article_id,comment,created_at,updated_at) VALUES (?,?,?,?)`, article_id, comment, unix_time, unix_time)
	db.MustExec(`UPDATE articles SET comments=comments+1 WHERE id=?`, article_id)
	c.JSON(200, gin.H{
		"status":     "ok",
		"id":         article_id,
		"comment":    comment,
		"created_at": natural_time_str,
	})
}

func (ct *ArticleController) GetComments(c *gin.Context) {
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

	article_id := c.Param("id")
	p := []Comments{}
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

	//sql select 10 comments
	err = db.Select(&p, "SELECT *, FROM_UNIXTIME(created_at, '%Y-%m-%d %H:%i') as created_at FROM comments WHERE article_id=? ORDER BY id LIMIT ?,10", article_id, skip)
	if err != nil {
		seelog.Error("can't read db ", err)
		return
	}

	//find all pages
	var all int
	err = db.Get(&all, "SELECT comments FROM articles WHERE id=?", article_id)
	if err != nil {
		seelog.Error("can't read db ", err)
		return
	}

	if all == 0 {
		all = 1
	} else {
		all_page := float64(all) / float64(10)
		allpage := math.Ceil(all_page)
		all = int(allpage)
	}

	c.JSON(200, gin.H{
		"status":       "ok",
		"id":           article_id,
		"comments":     &p,
		"current_page": current_page,
		"all_page":     all,
	})
}
