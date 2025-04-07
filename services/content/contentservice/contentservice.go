// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: content.proto

package contentservice

import (
	"context"

	"bilibili/services/content/proto/contentRpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Empty      = contentRpc.Empty
	PublishReq = contentRpc.PublishReq

	ContentService interface {
		Publish(ctx context.Context, in *PublishReq, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultContentService struct {
		cli zrpc.Client
	}
)

func NewContentService(cli zrpc.Client) ContentService {
	return &defaultContentService{
		cli: cli,
	}
}

func (m *defaultContentService) Publish(ctx context.Context, in *PublishReq, opts ...grpc.CallOption) (*Empty, error) {
	client := contentRpc.NewContentServiceClient(m.cli.Conn())
	return client.Publish(ctx, in, opts...)
}
