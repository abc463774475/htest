package gateway

import (
	"github.com/abc463774475/bbtool/n_log"
	"log"
	"math"
	"strings"
	"testing"
	"unsafe"
)

func TestF1(t *testing.T)  {

	n_log.Info("\u003e")

	str := `
	# Select *from t1;
	select id from t2;
`
	//str = strings.Replace(str, "select", "select /*max(id)*/",1)
	str = Replace(str,"select", func(matchStr, str string) bool {
		return strings.ToLower(matchStr) == strings.ToLower(str)
	},1)

	log.Printf("11\n%v\n", str)
}

//  Replace
func Replace(str string,matchStr string, matchFun func(matchStr, str string) bool, n int) string {
	if n <= 0 {
		n = math.MaxInt32
	}

	ret := make([]byte,0, len(str))
	s := matchStr

	count := 0
	for i:=0; i < len(str);{
		if i + len(s) < len(str) {
			if count < n && matchFun(matchStr, str[i:i+len(matchStr)]){
				ret = append(ret, []byte("select /*MAX_EXECUTION_TIME(1000)*/ ")...)
				i += len(s)
				count++

				if count >= n {
					ret = append(ret, str[i+len(matchStr):]...)
					break
				}
				continue
			}

			ret = append(ret, str[i])
			i++
		}else {
			ret = append(ret, str[i:]...)
			break
		}
	}

	return *(*string)(unsafe.Pointer(&ret))
}