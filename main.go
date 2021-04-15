package main

import (
	"ceph/gin/handler"
	"ceph/gin/router"
	"ceph/service"
)

func main() {

	sv := service.NewService()
    rt:=router.NewRouter(handler.NewHandler(sv))
    rt.Router()
}
