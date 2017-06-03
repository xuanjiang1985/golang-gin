package ctrl

import (
	valid "github.com/asaskevich/govalidator"
	seelog "github.com/cihub/seelog"
	"github.com/flosch/pongo2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang-gin/csrf"
	"golang-gin/sessions"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
	"strings"
	"time"
)

type AuthController struct {
}

func init() {
	valid.SetFieldsRequiredByDefault(true)
}

func (ct *AuthController) GetRegister(c *gin.Context) {
	session := sessions.Default(c)
	if s, ok := session.Get("authUserName").(string); ok && len(s) != 0 {
		//log.Println(session.Get("authUserName"))
		c.Redirect(302, "/")
		return
	}
	csrfToken := csrf.GetToken(c)
	c.HTML(http.StatusOK, "auth/register.html", pongo2.Context{
		"token": csrfToken,
	})
}

func (ct *AuthController) PostRegister(c *gin.Context) {
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

	type Validator struct {
		Name             string `valid:"required~昵称：不能为空,length(4|15)~昵称：4至15个字符之间"`
		Email            string `valid:"required~邮箱：不能为空,email~邮箱：必须是email格式"`
		Password         string `valid:"required~密码：不能为空,length(6|150)~密码：至少6个字符"`
		Confirm_password string `valid:"required~确认密码：不能为空,length(6|150)~确认密码：至少6个字符"`
	}
	data := &Validator{
		Name:             c.PostForm("昵称"),
		Email:            c.PostForm("邮箱"),
		Password:         c.PostForm("密码"),
		Confirm_password: c.PostForm("密码确认"),
	}
	//true or false of validator
	csrfToken := csrf.GetToken(c)
	ok, err := valid.ValidateStruct(data)
	if err != nil {
		msg := strings.Trim(err.Error(), ";")
		message := strings.Split(msg, ";")
		//log.Println(data)
		c.HTML(200, "auth/register.html", pongo2.Context{
			"token":  csrfToken,
			"errors": message,
			"data":   data,
		})
		return
	}
	//store user
	if ok {

	}
	password := []byte(c.PostForm("密码"))
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	unix_time := time.Now().Unix()
	result, err := db.Exec(`INSERT INTO users (name,email,password,created_at,updated_at) VALUES (?,?,?,?,?)`, c.PostForm("昵称"), c.PostForm("邮箱"), hashedPassword, unix_time, unix_time)
	if err != nil {
		c.HTML(200, "auth/register.html", pongo2.Context{
			"token":  csrfToken,
			"errors": []string{"邮箱已存在"},
			"data":   data,
		})
		return
	}
	//session start
	userId, _ := result.LastInsertId()
	// userInfo := map[string]string{
	// 	"name": c.PostForm("昵称"),
	// 	"id":   string(userId),
	// }
	//log.Println(userId)
	session := sessions.Default(c)
	session.Set("authUserName", c.PostForm("昵称"))
	session.Set("authUserId", userId)
	session.Save()
	c.Redirect(302, "/")
}

func (ct *AuthController) GetLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("authUserName")
	session.Delete("authUserId")
	session.Save()
	c.Redirect(302, "/")
}

func (ct *AuthController) PostLogin(c *gin.Context) {
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

	type Validator struct {
		Email    string `valid:"required~邮箱：不能为空,email~邮箱：必须是email格式"`
		Password string `valid:"required~密码：不能为空,length(6|150)~密码：至少6个字符"`
	}
	data := &Validator{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}
	//true or false of validator
	ok, err := valid.ValidateStruct(data)
	if err != nil {
		//log.Println(data)
		c.JSON(200, gin.H{
			"status": "ok",
			"error":  err.Error(),
		})
		return
	}
	//store user
	if ok {

	}
	password := []byte(c.PostForm("password"))
	var user Users
	err = db.Get(&user, "SELECT * FROM users WHERE email=?", c.PostForm("email"))
	if err != nil {
		//log.Println(err)
		c.JSON(200, gin.H{
			"status": "ok",
			"error":  "账户或密码错误1",
		})
		return
	}
	errors := bcrypt.CompareHashAndPassword([]byte(user.Password), password)
	if errors != nil {
		log.Println(errors)
		log.Println(user.Password)
		c.JSON(200, gin.H{
			"status": "ok",
			"error":  "账户或密码错误2",
		})
		return
	}
	//session start
	session := sessions.Default(c)
	session.Set("authUserName", user.Name)
	session.Set("authUserId", user.Id)
	session.Save()
	c.JSON(200, gin.H{
		"status": "ok",
		"error":  "",
	})
}

func (ct *AuthController) GetSetting(c *gin.Context) {
	authUser, _ := c.Get("authUser")
	csrfToken := csrf.GetToken(c)
	c.HTML(http.StatusOK, "auth/setting.html", pongo2.Context{
		"authUser": authUser,
		"token":    csrfToken,
	})
}
