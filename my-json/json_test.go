package my_json

import (
	"encoding/json"
	"fmt"
	"github.com/abc463774475/bbtool/n_log"
	"reflect"
	"testing"
)

func TestJson(t *testing.T)  {
	data,err := json.Marshal(struct {
		FF []int
		AA float64
		BB int
		CC bool
		DD struct{
			EE  int
		}

	}{FF:[]int{10,23},AA:1.5,BB:10,CC:false,DD: struct{ EE int }{EE: 22}},)

	n_log.Info("err  %v data\n%v", err,string(data))

	//
	//str := `{"dd": 1.5,"aa":10,"cc":false,"ee";"{\"e1\":123},"ff":["xx","yy"]"}`
	//m := map[string]interface{}{}
	//
	//err := json.Unmarshal([]byte(str), &m)
	//n_log.Info("err  %v", err)


}

func TestJson1(t *testing.T)  {
	str := `{"FF":[10,23],"HH":"hhhhh","AA":1.5,"BB":10,"CC":false,"DD":{"EE":22}}`
	m := map[string]interface{}{}

	json.Unmarshal([]byte(str), &m)

	for k,v := range m {
		s,e := anyToString(v)

		n_log.Info("k %v e %v s %v",k, e, s)
	}

}


func anyToString(i interface{}) (string, error) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Bool,reflect.Int, reflect.Float64:
		return fmt.Sprintf("%v",i), nil
	default:
		data, err := json.Marshal(i)
		if err != nil {
			return "", err
		}

		return string(data), err
	}
}