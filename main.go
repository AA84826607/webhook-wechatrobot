package main

import (
	"ceph/gin/handler"
	"ceph/gin/router"
	"ceph/service"
	"flag"
	"fmt"
)

func main() {
    var url string
    var port string
	flag.StringVar(&url, "u", "", "url for wechat work robot.")
    flag.StringVar(&port,"p","8989","port")
	flag.CommandLine.Usage = func() { fmt.Println("-u :the robot of the wechat API\n-p :expose the port") }
    flag.Parse()
    if url ==""{
    	panic("must add -u(the robot of the wechat API) and -p(port)")
	}
	sv := service.NewService(url)
    rt:=router.NewRouter(handler.NewHandler(sv),port)
    rt.Router()
}
