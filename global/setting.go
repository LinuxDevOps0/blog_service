package global

import (
	"github.com/jinzhu/gorm"
	"practice.com/blog_service/pkg/logger"
	"practice.com/blog_service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
)
