package yee

import (
	"github.com/abc463774475/bbtool/n_log"
	"github.com/cookieY/yee"
	_ "github.com/cookieY/yee"
	"net/http"
	"testing"
)

func TestYee(t *testing.T)  {
	e := yee.New()
	e.GET("/d5", func(context yee.Context) (err error) {
		n_log.Info("11111")
		rUrl := "https://www.baidu.com/"
		http.Redirect(context.Response(), context.Request(), rUrl, 301)
		return nil
	})

	e.Run(":8082")
}
