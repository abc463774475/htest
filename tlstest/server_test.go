package tlstest

import (
	"bufio"
	"crypto/tls"
	"github.com/abc463774475/bbtool/n_log"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		n_log.Erro("%v",err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		n_log.Erro("%v",err)
		return
	}
	defer ln.Close()
	n_log.Info("start")
	for {
		conn, err := ln.Accept()
		if err != nil {
			n_log.Erro("%v",err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	n_log.Info("recv one accept")
	defer conn.Close()
	r := bufio.NewReader(conn)

	tlcConn := conn.(*tls.Conn)
	if err := tlcConn.Handshake(); err != nil {
		n_log.Erro("err   %v", err)
		return
	}

	n_log.Debug("handleshark sucs")

	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			n_log.Erro("%v",err)
			return
		}
		n_log.Info(msg)
		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			n_log.Erro("n  %v %v", n, err)
			return
		}
	}
}