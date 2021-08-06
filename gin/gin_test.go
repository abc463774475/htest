package gin

import (
	"github.com/abc463774475/bbtool/n_log"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	"log"
	"testing"
)

var (
	mAll = map[string]interface{}{}
)

func TestGin(t *testing.T)  {
	log.Printf("111")
	r := gin.Default()
	r.GET("/d1", func(context *gin.Context) {
		//rUrl := "https://www.baidu.com/"
		//q := url.Values{}
		//q.Set("wage", "123")
		//q.Set("amount", "13123")
		//location := url.URL{Path: rUrl, RawQuery: q.Encode()}
		//context.Redirect(http.StatusMovedPermanently, location.RequestURI())

		n_log.Info("11111")
		rUrl := "https://www.baidu.com/"
		context.Redirect(301, rUrl)
	})

	r.POST("/p2", func(context *gin.Context) {
		str, err := uuid.GenerateUUID()
		if err != nil {
			n_log.Erro("err  %v", err)
			return
		}
		l := 1024*1024*1024
		data := make([]byte, l)

		for i := range data {
			data[i] = '1'
		}

		mAll[str] = data
		context.String(200,str)
	})

	r.Run(":8081")
}