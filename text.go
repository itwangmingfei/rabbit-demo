package main

import (
	"log"
	"os"
	"strings"
)

func main(){
	log.Println(bodyFrom(os.Args))
}
//监听传入参数
func bodyFrom(args []string) string {
	var s string
	log.Println(len(args),os.Args[0])
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}