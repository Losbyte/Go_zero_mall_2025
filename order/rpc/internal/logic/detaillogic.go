package logic

import (
	"context"
	"mall/service/order/model"
	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *order.DetailRequest) (*order.DetailResponse, error) {
	res,err:=l.svcCtx.OrderModel.FindOne(l.ctx,uint64(in.Id))
	if err!=nil{
		if err==model.ErrNotFound{
			return nil,status.Error(100,"订单不存在")
		}
		return nil,status.Error(500,err.Error())
	}
	return &order.DetailResponse{
		Id: int64(res.Id),
		Uid: int64(res.Uid),
		Pid: int64(res.Pid),
		Amount: int64(res.Amount),
		Status: int64(res.Status),
	}, nil
}
