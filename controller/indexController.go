package ctrl

import (
	"github.com/flosch/pongo2"
	"golang-gin/csrf"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type IndexController struct {
}

func (ct *IndexController) Get(c *gin.Context) {
	csrfToken := csrf.GetToken(c)
	c.HTML(http.StatusOK, "index.html", pongo2.Context{"token": csrfToken})
}
