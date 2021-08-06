package main

import (
	_ "embed"
	"log"
)

//go:embed hhhv.txt
var hhhv string

func main()  {
	log.Printf("embed   is  %v\n",hhhv)
}
