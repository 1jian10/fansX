package svc

import (
	"context"
	bigcache "fansX/internal/middleware/cache"
	"fansX/internal/middleware/lua"
	"fansX/internal/util"
	"fansX/services/content/public/proto/publicContentRpc"
	"fansX/services/feed/internal/config"
	"fansX/services/feed/internal/script"
	"fansX/services/relation/proto/relationRpc"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	etcd "go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
)

type ServiceContext struct {
	Config         config.Config
	RClient        *redis.Client
	DB             *gorm.DB
	Logger         *slog.Logger
	Executor       *lua.Executor
	Cache          *bigcache.Cache
	RelationClient relationRpc.RelationServiceClient
	ContentClient  publicContentRpc.PublicContentServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := "root:@tcp(linux.1jian10.cn:4000)/test?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	r := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   1,
	})

	if err := r.Ping(context.Background()).Err(); err != nil {
		panic(err.Error())
	}

	logger, err := util.InitLog("FeedService", slog.LevelDebug)
	if err != nil {
		panic(err.Error())
	}

	e := lua.NewExecutor(r)
	_, err = e.Load(context.Background(), []*lua.Script{
		script.RangeByScore,
		script.RevRange,
	})
	if err != nil {
		panic(err.Error())
	}
	eClient, err := etcd.New(etcd.Config{
		Endpoints: []string{"1jian10.cn:4379"},
	})
	if err != nil {
		panic(err.Error())
	}

	cache := bigcache.NewCache(eClient)

	zClient := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:4379"},
		},
	})
	relationClient := relationRpc.NewRelationServiceClient(zClient.Conn())
	contentClient := publicContentRpc.NewPublicContentServiceClient(zClient.Conn())

	svc := &ServiceContext{
		Config:         c,
		DB:             db,
		RClient:        r,
		Logger:         logger,
		Executor:       e,
		Cache:          cache,
		RelationClient: relationClient,
		ContentClient:  contentClient,
	}

	return svc
}
