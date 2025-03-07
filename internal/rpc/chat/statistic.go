package chat

import (
	"context"
	"github.com/OpenIMSDK/chat/pkg/proto/chat"
	"github.com/OpenIMSDK/tools/errs"
	"time"
)

func (o *chatSvr) NewUserCount(ctx context.Context, req *chat.NewUserCountReq) (*chat.NewUserCountResp, error) {
	resp := &chat.NewUserCountResp{}
	if req.Start > req.End {
		return nil, errs.ErrArgs.Wrap("start > end")
	}
	total, err := o.Database.NewUserCountTotal(ctx, nil)
	if err != nil {
		return nil, err
	}
	start := time.UnixMilli(req.Start)
	before, err := o.Database.NewUserCountTotal(ctx, &start)
	if err != nil {
		return nil, err
	}
	end := time.UnixMilli(req.End)
	count, err := o.Database.NewUserCountRangeEverydayTotal(ctx, &start, &end)
	if err != nil {
		return nil, err
	}
	resp.Total = total
	resp.Before = before
	resp.Count = count
	return resp, nil
}
func (o *chatSvr) UserLoginCount(ctx context.Context, req *chat.UserLoginCountReq) (*chat.UserLoginCountResp, error) {
	resp := &chat.UserLoginCountResp{}
	if req.Start > req.End {
		return nil, errs.ErrArgs.Wrap("start > end")
	}
	total, err := o.Database.UserLoginCountTotal(ctx, nil)
	if err != nil {
		return nil, err
	}
	start := time.UnixMilli(req.Start)
	before, err := o.Database.UserLoginCountTotal(ctx, &start)
	if err != nil {
		return nil, err
	}
	end := time.UnixMilli(req.End)
	count, err := o.Database.UserLoginCountRangeEverydayTotal(ctx, &start, &end)
	if err != nil {
		return nil, err
	}
	resp.Total = total
	resp.Before = before
	resp.Count = count
	return resp, nil
}
