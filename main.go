package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"practice.com/blog_service/global"
	"practice.com/blog_service/internal/model"
	"practice.com/blog_service/internal/router"
	"practice.com/blog_service/pkg/logger"
	"practice.com/blog_service/pkg/setting"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	global.Logger.WithCaller(1).Infof("%s: go-programming-tour-book/%s", "yif", "blog_service")
	gin.SetMode(global.ServerSetting.RunMode)
	routerHandler := router.NewRouter()

	server := http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           routerHandler,
		ReadTimeout:       global.ServerSetting.ReadTimeout,
		WriteTimeout:      global.ServerSetting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}

func setupSetting() error {
	projectSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = projectSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	err = projectSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = projectSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}


func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:   global.AppSetting.LogSavePath + "/" +global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:    600,
		MaxAge:     10,
		MaxBackups: 0,
		LocalTime:  true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

