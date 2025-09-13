package logic

import (
	"context"
	"mall/service/order/model"
	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/order"
	"mall/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *order.ListRequest) (*order.ListResponse, error) {
	//查询用户是否存在
	_,err := l.svcCtx.UserRpc.UserInfo(l.ctx,&userclient.UserInfoRequest{
		Id: in.Uid,
	})
	if err!=nil{
		return nil,err
	}
	//注意此处需要自行实现分组查询
	list,err:= l.svcCtx.OrderModel.FindAllByUid(l.ctx,in.Uid)
	if err!=nil{
		if err==model.ErrNotFound{
			return nil,status.Error(100,"订单不存在")
		}
		return nil,status.Error(500,err.Error())
	}
	orderList:=make([]*order.DetailResponse,0)
	for _,item:=range list{
		orderList = append(orderList,&order.DetailResponse{
			Id: int64(item.Id),
			Uid:int64(item.Uid),
			Pid:int64(item.Pid),
			Amount: int64(item.Amount),
			Status: int64(item.Status),
		})
	}

	return &order.ListResponse{
		Data: orderList,
	}, nil
}
