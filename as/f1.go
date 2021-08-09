package main

import (
	"github.com/abc463774475/bbtool/n_log"
	"runtime"
	"time"
)

func f1()  {
	l := 200*1024*1024
	var data []byte
	// data = make([]byte,0, l)
	for i:=0; i < l; i++{
		data = append(data, 'h')
	}

	n_log.Info("1111111   %v    %v", len(data), cap(data) - len(data))
	time.Sleep(10*time.Second)
	runtime.GC()

	n_log.Info("2222222")
	time.Sleep(120*time.Second)

	// n_log.Info("data len  %v    %v",len(data), data[l-1])

	n_log.Info("data len  %v  ",len(data), )
}

func f2()  {
	l := 200*1024*1024
	var data []byte
	data = make([]byte,0, l)
	for i:=0; i < l; i++{
		data = append(data, 'h')
	}

	n_log.Info("1111111   %v    %v", len(data), cap(data) - len(data))
	time.Sleep(10*time.Second)
	//runtime.GC()

	n_log.Info("2222222")
	time.Sleep(120*time.Second)

	// n_log.Info("data len  %v    %v",len(data), data[l-1])

	n_log.Info("data len  %v   ",len(data), )
}