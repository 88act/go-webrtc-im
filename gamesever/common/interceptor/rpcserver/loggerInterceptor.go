package rpcserver

import (
	"context"

	"go-game/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/**
* @Description rpc service logger interceptor
* @Author Mikael
* @Date 2021/1/9 13:35
* @Version 1.0
**/

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR-1】 %+v", err)
			//转成grpc err
			err = status.Error(codes.Code(e.GetErrCode()), e.GetErrMsg())
			//err = status.Error(codes.Code(e.GetErrCode()), err.Error())
		} else {
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR-2】 %+v", err)
		}
	}

	return resp, err
}
