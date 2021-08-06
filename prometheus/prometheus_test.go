package prometheus

import (
	"context"
	"github.com/abc463774475/bbtool/n_log"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"log"
	"testing"
	"time"
)

func TestPrometheus(t *testing.T) {
	client, err := api.NewClient(api.Config{
		Address:      "http://172.31.2.24:9090",
	})
	if err != nil {
		log.Panicln(err)
	}

	v2 := v1.NewAPI(client)

	//str1 := `100 - ((node_filesystem_avail_bytes{mountpoint="/",fstype!="rootfs"} * 100) /            node_filesystem_size_bytes{mountpoint="/",fstype!="rootfs"})`
	str1 := `100 - node_filesystem_free_bytes{
          fstype!~"rootfs|selinuxfs|autofs|rpc_pipefs|tmpfs|udev|none|devpts|sysfs|debugfs|fuse.*"} / node_filesystem_size_bytes{fstype!~"rootfs|selinuxfs|autofs|rpc_pipefs|tmpfs|udev|none|devpts|sysfs|debugfs|fuse.*"} * 100
`
	// node_filesystem_size_bytes node_filesystem_free_bytes

	m, err := v2.Query(context.Background(), "node_filesystem_size_bytes",time.Now())
	if err != nil {
		panic(err)
	}

	log.Printf("mm %v\n%v\n", m.Type(),m)
	{
		m1,ok := m.(model.Vector)
		if !ok {
			log.Printf("m  %v", ok)
		}
		d := int64(m1[0].Value)
		n_log.Info("d  %v", d)
	}

	{
		m, err := v2.Query(context.Background(), str1, time.Now())
		if err != nil {
			panic(err)
		}

		log.Printf("mm %v\n%v\n", m.Type(), m)
	}
}


