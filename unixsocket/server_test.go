package unixsocket

import (
	"bufio"
	"net"
	"testing"
	"time"

	"github.com/abc463774475/bbtool/n_log"
)

func TestServere(t *testing.T) {
	fileName := "test.sock"

	lis, err := net.Listen("unix", fileName)
	if err != nil {
		n_log.Panic("err  %v", err)
	}

	n_log.Info("create unix domain socket suc")

	defer lis.Close()

	for {
		conn, err := lis.Accept()
		n_log.Info("err  %v", err)

		if err != nil {
			n_log.Erro("err  %v", err)
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()

			reader := bufio.NewReader(conn)
			for {
				//data := make([]byte, 0, 1024)
				//n, err := conn.Read(data)
				//n_log.Info("n  %v err  %v", n, err)
				//
				//if n == 0 || err != nil {
				//	time.Sleep(60*time.Second)
				//	return
				//}

				msg, err := reader.ReadString('\n')
				if err != nil {
					n_log.Erro("err  %v",err)
					return
				}

				n_log.Info("msg  %v", msg,)
			}
		}(conn)
	}
}

func TestClient(t *testing.T) {
	fileName := "test.sock"
	conn, err := net.Dial("unix", fileName)
	if err != nil {
		n_log.Panic("err  %v", err)
	}

	defer conn.Close()

	time.Sleep(2*time.Second)
	n, err := conn.Write([]byte("1111111\r\n"))
	n_log.Info("n  %v  err  %v", n, err)

	time.Sleep(600 * time.Second)
}
