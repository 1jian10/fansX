// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: publiccontent.proto

package server

import (
	"context"

	"fansX/services/content/public/internal/logic"
	"fansX/services/content/public/internal/svc"
	"fansX/services/content/public/proto/publicContentRpc"
)

type PublicContentServiceServer struct {
	svcCtx *svc.ServiceContext
	publicContentRpc.UnimplementedPublicContentServiceServer
}

func NewPublicContentServiceServer(svcCtx *svc.ServiceContext) *PublicContentServiceServer {
	return &PublicContentServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *PublicContentServiceServer) GetUserContentList(ctx context.Context, in *publicContentRpc.GetUserContentListReq) (*publicContentRpc.GetUserContentListResp, error) {
	l := logic.NewGetUserContentListLogic(ctx, s.svcCtx)
	return l.GetUserContentList(in)
}
