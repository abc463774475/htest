package tlstest

import (
	"crypto/tls"
	"github.com/abc463774475/bbtool/n_log"
	"log"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", ConnectAddr, conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		n_log.Erro(" %v  %v",n, err)
		return
	}
	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		n_log.Erro(" %v  %v",n, err)
		return
	}
	n_log.Debug(string(buf[:n]))


	time.Sleep(10*time.Second)
}