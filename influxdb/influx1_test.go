package influxdb

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/abc463774475/bbtool/n_log"
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client"
	_ "github.com/jinzhu/gorm"
)

var conn *client.Client

func Init(ip string, port int) {
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", ip, port))
	if err != nil {
		log.Fatal(err)
	}

	conf := client.Config{
		URL: *host,
		// Username: os.Getenv("INFLUX_USER"),
		// Password: os.Getenv("INFLUX_PWD"),
	}
	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	conn = con
}

func GetDataBases() ([]string, error) {
	res, err := conn.Query(client.Query{
		Command:         "show databases",
		Database:        "",
		RetentionPolicy: "",
		Chunked:         false,
		ChunkSize:       0,
		NodeID:          0,
	})
	if err != nil || res == nil {
		return nil, err
	}

	errNotExit := errors.New("not exist")

	ret := []string{}
	if len(res.Results) == 0 || len(res.Results[0].Series) == 0 {
		return nil, errNotExit
	}
	vals := res.Results[0].Series[0].Values
	for i := range vals {
		if len(vals[i]) > 0 {
			name, ok := vals[i][0].(string)
			if ok && name != "_internal"{
				ret = append(ret, name)
			}
		}
	}

	return ret, nil
}

func GetMeasurements(database string) ([]string, error) {
	res, err := conn.Query(client.Query{
		Command:         "show measurements",
		Database:        database,
		RetentionPolicy: "",
		Chunked:         false,
		ChunkSize:       0,
		NodeID:          0,
	})

	if err != nil || res == nil {
		return nil, err
	}

	errNotExit := errors.New("not exist")

	ret := []string{}
	if len(res.Results) == 0 || len(res.Results[0].Series) == 0 {
		return nil, errNotExit
	}
	vals := res.Results[0].Series[0].Values
	for i := range vals {
		if len(vals[i]) > 0 {
			name, ok := vals[i][0].(string)
			if ok {
				ret = append(ret, name)
			}
		}
	}

	return ret, nil
}

type FiledType string

const (
	FiledType_Integer FiledType = "integer"
	FiledType_Float FiledType = "float"
	FiledType_String FiledType = "string"
)

type FiledInfo struct {
	Type FiledType
	Name string
}

func GetFileds(database string, measurement string) ([]FiledInfo, error) {
	res, err := conn.Query(client.Query{
		Command:         fmt.Sprintf("show field keys from %v", measurement),
		Database:        database,
		RetentionPolicy: "",
		Chunked:         false,
		ChunkSize:       0,
		NodeID:          0,
	})

	if err != nil || res == nil {
		return nil, err
	}

	errNotExit := errors.New("not exist")

	ret := []FiledInfo{}
	if len(res.Results) == 0 || len(res.Results[0].Series) == 0 {
		return nil, errNotExit
	}
	vals := res.Results[0].Series[0].Values
	for i := range vals {
		if len(vals[i]) > 0 {
			fieldType, ok1 := vals[i][1].(string)
			name, ok2 := vals[i][0].(string)
			if ok1 && ok2 {
				ret = append(ret, FiledInfo{
					Type: FiledType(fieldType),
					Name: name,
				})
			}
		}
	}

	return ret, nil
}


func TestClient(t *testing.T) {
	Init("localhost", 8086)
	//res, err := conn.Query(client.Query{
	//	Command:         "show databases",
	//	Database:        "",
	//	RetentionPolicy: "",
	//	Chunked:         false,
	//	ChunkSize:       0,
	//	NodeID:          0,
	//})

	res, err := GetDataBases()

	n_log.Info("res   %v   \n%v", err, res)
}

func TestMesures(t *testing.T) {
	Init("localhost", 8086)

	//res, err := conn.Query(client.Query{
	//	Command:         "show measurements",
	//	Database:        "my-bucket",
	//	RetentionPolicy: "",
	//	Chunked:         false,
	//	ChunkSize:       0,
	//	NodeID:          0,
	//})

	res,err := GetMeasurements("my-bucket")


	n_log.Info("res   %v   \n%v", err, res)
}

func TestFileds(t *testing.T)  {
	Init("localhost", 8086)
	//res, err := conn.Query(client.Query{
	//	Command:         fmt.Sprintf("show field keys from %v", "ms1"),
	//	Database:        "my-bucket",
	//	RetentionPolicy: "",
	//	Chunked:         false,
	//	ChunkSize:       0,
	//	NodeID:          0,
	//})
	res, err := GetFileds("my-bucket", "ms1")

	n_log.Info("err  %v  res\n%v", err, res)
}

func TestPing(t *testing.T)  {
	Init("localhost",8086)
	t1, str, err := conn.Ping()
	n_log.Info("t1  %v  str  %v  err  %v", t1, str, err)
}

func TestSelect(t *testing.T)  {
	strs := strings.Replace("insert ms1,r1=pp av1=99i,sum=8.88","insert","",1)

	n_log.Info("strs \n%v",strs)

	return
	Init("localhost", 8086)
	pts := make([]client.Point, 1)
	pts[0] = client.Point{
		//Measurement: "ms1",
		Tags:        nil,
		Time:        time.Time{},
		//Fields: map[string]interface{}{
		//	"av1":99,
		//	"sum": 8.88,
		//},
		Precision:   "",
		Raw:    "ms1,r1=pp av1=99i,sum=8.88",    // "insert ms1,r1=pp av1=99i,sum=8.88",
	}

	res, err := conn.Write(client.BatchPoints{
		Database:        "my-bucket",
		RetentionPolicy: "",
		Points: pts,
	})

	n_log.Info("err  %v  res  %v", err, res)
}