package svc

import (
	"fansX/internal/util"
	leaf "fansX/pkg/leaf-go"
	"fansX/services/content/meta/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Logger  *slog.Logger
	Creator leaf.Core
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := "root:@tcp(linux.1jian10.cn:4000)/test?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	logger, err := util.InitLog("ContentService", slog.LevelDebug)
	if err != nil {
		panic(err.Error())
	}
	creator, err := leaf.Init(&leaf.Config{
		Model: leaf.Snowflake,
		SnowflakeConfig: &leaf.SnowflakeConfig{
			CreatorName: "metacontent.rpc",
			Addr:        "1jian10.cn:23020",
			EtcdAddr:    []string{"1jian10.cn:4379"},
		},
	})
	if err != nil {
		panic(err.Error())
	}

	svc := &ServiceContext{
		DB:      db,
		Logger:  logger,
		Creator: creator,
	}

	return svc
}
