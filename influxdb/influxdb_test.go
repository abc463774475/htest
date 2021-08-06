package influxdb

import (
	"context"
	"testing"
	"time"

	"github.com/influxdata/influxdb-client-go/domain"

	"github.com/abc463774475/bbtool/n_log"
	_ "github.com/influxdata/influxdb-client-go"
	influxdb2 "github.com/influxdata/influxdb-client-go"
)

func TestConn(t *testing.T) {
	client := influxdb2.NewClient("http://localhost:8086", "")
	if client == nil {
		n_log.Panic("client nil")
	}

	queryAPI := client.QueryAPI("my-org")
	queryAPI.QueryRaw(context.Background(), "show databases ;", &domain.Dialect{
		Annotations:    nil,
		CommentPrefix:  nil,
		DateTimeFormat: nil,
		Delimiter:      nil,
		Header:         nil,
	})

	writeAPI := client.WriteAPIBlocking("my-org", "my-bucket")

	p := influxdb2.NewPoint("stat", map[string]string{
		"unit": "temperature",
	}, map[string]interface{}{
		"avg": 24.5, "max": 45.0,
	}, time.Now())

	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		n_log.Panic("err  %v", err)
	}

	return

	err = writeAPI.WritePoint(context.Background(),
		influxdb2.NewPoint(
			"ms1", map[string]string{
				"r1": "xxxx",
			}, map[string]interface{}{
				"av1": 60,
				"sum": 10.99,
			}, time.Now(),
		))
	if err != nil {
		n_log.Panic("err  %v", err)
	}

	time.Sleep(100 * time.Second)
}

func TestQuery(t *testing.T) {
	client := influxdb2.NewClient("http://localhost:8086", "")
	if client == nil {
		n_log.Panic("client nil")
	}

	queryAPI := client.QueryAPI("my-org")
	str, err := queryAPI.Query(context.Background(), "show databases")

	n_log.Info("err  %v\n%v", err, str)
}
