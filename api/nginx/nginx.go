package nginx

import (
	"net/http"

	"github.com/0xJacky/Nginx-UI/internal/nginx"
	"github.com/gin-gonic/gin"
	"github.com/uozi-tech/cosy"
)

func BuildNginxConfig(c *gin.Context) {
	var ngxConf nginx.NgxConfig
	if !cosy.BindAndValid(c, &ngxConf) {
		return
	}
	content, err := ngxConf.BuildConfig()
	if err != nil {
		cosy.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"content": content,
	})
}

func TokenizeNginxConfig(c *gin.Context) {
	var json struct {
		Content string `json:"content" binding:"required"`
	}

	if !cosy.BindAndValid(c, &json) {
		return
	}

	ngxConfig, err := nginx.ParseNgxConfigByContent(json.Content)
	if err != nil {
		cosy.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, ngxConfig)
}

func FormatNginxConfig(c *gin.Context) {
	var json struct {
		Content string `json:"content"`
	}

	if !cosy.BindAndValid(c, &json) {
		return
	}
	content, err := nginx.FmtCode(json.Content)
	if err != nil {
		cosy.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"content": content,
	})
}
