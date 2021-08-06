package tlstest

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"github.com/abc463774475/bbtool/n_log"
	"io/ioutil"
	"log"
	"net"
	"testing"
)

func TestTwoEndPointsServer(t *testing.T) {
	cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		log.Println(err)
		return
	}
	certBytes, err := ioutil.ReadFile("client.pem")
	if err != nil {
		panic("Unable to read cert.pem")
	}
	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("failed to parse root certificate")
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    clientCertPool,
	}
	ln, err := tls.Listen("tcp", ConnectAddr, config)
	if err != nil {
		n_log.Erro("%v",err)
		return
	}

	defer ln.Close()
	n_log.Debug("start")

	for {
		conn, err := ln.Accept()
		if err != nil {
			n_log.Erro("%v",err)
			continue
		}
		go handleConnte(conn)
	}
}

func handleConnte(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)

	if err := conn.(*tls.Conn).Handshake(); err != nil {
		n_log.Erro("err  %v",err)
		return
	}

	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			n_log.Erro("%v",err)
			return
		}
		println(msg)
		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			n_log.Erro("%v %v",n,err)
			return
		}
	}
}
