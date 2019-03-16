package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-log"
	"gomicro_warmhome/homeweb/handler"
	_ "gomicro_warmhome/homeweb/models"
	"gomicro_warmhome/homeweb/utils"
	"net/http"

	"github.com/micro/go-web"
)

func main() {
	//创建服务
	service := web.NewService(
		web.Name("go.micro.web.homeweb"),
		web.Version("latest"),
		web.Address(":"+utils.G_server_port),
	)

	//初始化服务
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	rou := httprouter.New()
	//映射静态页面
	rou.NotFound = http.FileServer(http.Dir("html"))
	//获取地区数据
	rou.GET("/api/v1.0/areas", handler.GetArea)
	//获取session
	rou.GET("/api/v1.0/session", handler.GetSession)
	//获取首页轮播图
	rou.GET("/api/v1.0/house/index", handler.GetIndex)

	//路由初始化
	service.Handle("/", rou)

	//运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}