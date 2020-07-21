package global

import (
	"blog-service/configs/settings"
	"go.uber.org/zap"
)

var (
	ServerSetting   *settings.ServerSettings
	AppSetting      *settings.AppSettings
	DatabaseSetting *settings.DatabaseSettings
	//
	//Logger *logger.Logger
	Logger *zap.SugaredLogger
)
