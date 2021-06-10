package main

import (
	"ceph/gin/handler"
	"ceph/gin/router"
	"ceph/service"
	"flag"
	"fmt"
)

func main() {
	var key string
	var port string
	flag.StringVar(&key, "k", "", "Key for wechat work robot.")
	flag.StringVar(&port, "p", "8989", "port")
	flag.CommandLine.Usage = func() { fmt.Println("-k :the robot of the wechat Key\n-p :expose the port") }
	flag.Parse()
	if key == "" {
		panic("must add -u(the robot of the wechat API) and -p(port)")
	}
	sv := service.NewService(key)
	rt := router.NewRouter(handler.NewHandler(sv), port)
	rt.Router()
}
