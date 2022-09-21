package commons

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
)

type Response struct {
	Code int
	Data interface{}
	Msg  string
	Url  string
}

func Fail(c *context.Context, message string, data interface{}, url string) {
	Result(c, http.StatusInternalServerError, data, message, url)
}

func Success(c *context.Context, data interface{}, message string, url string) {
	Result(c, http.StatusOK, data, message, url)
}

func Result(c *context.Context, code int, data interface{}, msg string, url string) {
	if c.Input.IsPost() || c.Input.IsPut() || c.Input.IsDelete() {
		result := Response{
			Code: code,
			Msg:  msg,
			Data: data,
			Url:  url,
		}
		c.Output.JSON(result, false, false)
		panic(beego.ErrAbort)
	}
	if url == "" {
		url = c.Request.Referer()
		if url == "" {
			url = "/"
		}
	}
	c.Redirect(http.StatusFound, url)
}
