package main

import (
	"log"
	"time"
)

func f1() int {
	l := 200*1024*1024
	var data []byte
	// data = make([]byte,0, l)
	for i:=0; i < l; i++{
		data = append(data, 'h')
	}

	log.Println("1111111   %v    %v", len(data), cap(data) - len(data))
	time.Sleep(10*time.Second)
	// runtime.GC()

	log.Println("2222222")
	time.Sleep(120*time.Second)

	// n_log.Info("data len  %v    %v",len(data), data[l-1])

	//log.Println("data len  %v  ",len(data), )
	return len(data)
}
