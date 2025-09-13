package logic

import (
	"context"
	"mall/service/order/model"
	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/order"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *order.UpdateRequest) (*order.UpdateResponse, error) {
	res,err:=l.svcCtx.OrderModel.FindOne(l.ctx,uint64(in.Id))
	if err!=nil{
		if err == model.ErrNotFound{
			return nil,status.Error(100,"订单不存在")
		}
		return nil,status.Error(500,err.Error())
	}
	if in.Uid!=0{
		res.Uid = uint64(in.Uid)
	}
	if in.Pid!=0{
		res.Pid = uint64(in.Pid)
	}
	if in.Amount!=0{
		res.Amount = uint64(in.Amount)
	}
	if in.Status!=0{
		res.Status = uint64(in.Status)
	}
	err = l.svcCtx.OrderModel.Update(l.ctx,res)
	if err!=nil{
		return nil,status.Error(500,err.Error())
	}
	return &order.UpdateResponse{}, nil
}
