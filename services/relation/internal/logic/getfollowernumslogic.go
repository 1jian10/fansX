package logic

import (
	"bilibili/common/util"
	"bilibili/model"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"

	"bilibili/services/relation/internal/svc"
	"bilibili/services/relation/proto/relationRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerNumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerNumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerNumsLogic {
	return &GetFollowerNumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerNumsLogic) GetFollowerNums(in *relationRpc.GetFollowerNumsReq) (*relationRpc.GetFollowerNumsResp, error) {
	db := l.svcCtx.DB
	logger := util.SetTrace(context.Background(), l.svcCtx.Logger)
	client := l.svcCtx.RClient

	logger.Info("GetFollowerNums", "userid", in.UserId)
	timeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	key := "FollowerNums:" + strconv.FormatInt(in.UserId, 10)

	res, err := client.Get(timeout, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		logger.Error("get follower nums from redis:" + err.Error())
		return nil, err
	}
	if err == nil {
		nums, err := strconv.ParseInt(res, 10, 64)
		if err != nil {
			logger.Error("parse follower nums:" + err.Error())
			return nil, err
		}
		logger.Debug("get follower nums from redis")
		return &relationRpc.GetFollowerNumsResp{Nums: nums}, nil
	}

	logger.Debug("not found follower nums from redis")

	record, err := l.svcCtx.Single.Do("GetFollowerNums:"+strconv.FormatInt(in.UserId, 10), func() (interface{}, error) {
		record := &model.FollowerNums{}
		err = db.Take(record, in.UserId).Error
		if err != nil {
			return 0, err
		}
		client.Set(context.Background(), key, record.Nums, time.Minute*5)
		return record, nil
	})

	if err != nil {
		logger.Error("get follower nums from tidb:" + err.Error())
		return nil, err
	}

	return &relationRpc.GetFollowerNumsResp{Nums: record.(*model.FollowerNums).Nums}, nil
}
