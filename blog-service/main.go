package main

import (
	"blog-service/configs/settings"
	"blog-service/global"
	"blog-service/internal/model"
	"blog-service/internal/routers"
	"blog-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func setupSetting() error {
	setting, err := settings.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	log.Printf("global.ServerSetting: %+v", global.ServerSetting)
	log.Printf("global.AppSetting: %+v", global.AppSetting)
	log.Printf("global.DatabaseSetting: %+v", global.DatabaseSetting)

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

func setupLogger() error {
	w := &lumberjack.Logger{
		Filename: filepath.Join(global.AppSetting.LogSavePath, global.AppSetting.LogFileName) +
			global.AppSetting.LogFileExt,
		MaxSize:   1 << 0,
		MaxAge:    10,
		LocalTime: true,
		Compress:  true,
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(logger.GetEncoderConfig()),
		zapcore.AddSync(w),
		zap.NewAtomicLevelAt(zapcore.DebugLevel))
	global.Logger = zap.New(core, zap.AddCaller()).Sugar()
	return nil
}

func main() {
	global.Logger.Info("Init successfully")
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
