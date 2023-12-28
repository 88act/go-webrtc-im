package logic

import (
	"context"

	"go-game/app/basic/rpc/internal/svc"
	"go-game/app/basic/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证码
func (l *SendCodeLogic) SendCode(in *pb.SendCodeReq) (*pb.SendCodeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SendCodeResp{}, nil
}
