package main

import (
	"fmt"
	"go_jin_testing/models"
	"go_jin_testing/pkg/settings"
	"go_jin_testing/routers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	settings.Setup()
	models.Setup()
	// gredis.Setup()
	// util.Setup()
}

func main() {
	gin.SetMode(settings.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := settings.ServerSetting.ReadTimeout
	writeTimeout := settings.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", settings.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Print("[info] start http server on", endPoint)
	server.ListenAndServe()
}
