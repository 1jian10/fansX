// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: like.proto

package likeservice

import (
	"context"

	"fansX/services/like/proto/likeRpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CancelLikeReq     = likeRpc.CancelLikeReq
	Empty             = likeRpc.Empty
	GetLikeNumsReq    = likeRpc.GetLikeNumsReq
	GetLikeNumsResp   = likeRpc.GetLikeNumsResp
	LikeReq           = likeRpc.LikeReq
	ListLikesReq      = likeRpc.ListLikesReq
	ListLikesResp     = likeRpc.ListLikesResp
	ListUserLikesReq  = likeRpc.ListUserLikesReq
	ListUserLikesResp = likeRpc.ListUserLikesResp

	LikeService interface {
		Like(ctx context.Context, in *LikeReq, opts ...grpc.CallOption) (*Empty, error)
		CancelLike(ctx context.Context, in *CancelLikeReq, opts ...grpc.CallOption) (*Empty, error)
		ListUserLikes(ctx context.Context, in *ListUserLikesReq, opts ...grpc.CallOption) (*ListUserLikesResp, error)
		ListLikes(ctx context.Context, in *ListLikesReq, opts ...grpc.CallOption) (*ListLikesResp, error)
		GetLikeNums(ctx context.Context, in *GetLikeNumsReq, opts ...grpc.CallOption) (*GetLikeNumsResp, error)
	}

	defaultLikeService struct {
		cli zrpc.Client
	}
)

func NewLikeService(cli zrpc.Client) LikeService {
	return &defaultLikeService{
		cli: cli,
	}
}

func (m *defaultLikeService) Like(ctx context.Context, in *LikeReq, opts ...grpc.CallOption) (*Empty, error) {
	client := likeRpc.NewLikeServiceClient(m.cli.Conn())
	return client.Like(ctx, in, opts...)
}

func (m *defaultLikeService) CancelLike(ctx context.Context, in *CancelLikeReq, opts ...grpc.CallOption) (*Empty, error) {
	client := likeRpc.NewLikeServiceClient(m.cli.Conn())
	return client.CancelLike(ctx, in, opts...)
}

func (m *defaultLikeService) ListUserLikes(ctx context.Context, in *ListUserLikesReq, opts ...grpc.CallOption) (*ListUserLikesResp, error) {
	client := likeRpc.NewLikeServiceClient(m.cli.Conn())
	return client.ListUserLikes(ctx, in, opts...)
}

func (m *defaultLikeService) ListLikes(ctx context.Context, in *ListLikesReq, opts ...grpc.CallOption) (*ListLikesResp, error) {
	client := likeRpc.NewLikeServiceClient(m.cli.Conn())
	return client.ListLikes(ctx, in, opts...)
}

func (m *defaultLikeService) GetLikeNums(ctx context.Context, in *GetLikeNumsReq, opts ...grpc.CallOption) (*GetLikeNumsResp, error) {
	client := likeRpc.NewLikeServiceClient(m.cli.Conn())
	return client.GetLikeNums(ctx, in, opts...)
}
