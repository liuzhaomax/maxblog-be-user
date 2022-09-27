package service

import (
	"context"
	"github.com/google/wire"
	"maxblog-be-user/internal/core"
	"maxblog-be-user/src/model"
	"maxblog-be-user/src/pb"
)

var UserSet = wire.NewSet(wire.Struct(new(BUser), "*"))

type BUser struct {
	MUser   *model.MUser
	Tx      *core.Trans
	ILogger core.ILogger
}

func (bUser *BUser) GetUserById(ctx context.Context, req *pb.IdRequest) (*pb.UserRes, error) {
	var data model.User
	err := bUser.Tx.ExecTrans(ctx, func(ctx context.Context) error {
		err := bUser.MUser.QueryUserById(req, &data)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		bUser.ILogger.LogFailure(core.GetFuncName(), err)
		return nil, err
	}
	res := model.Model2PB(&data)
	bUser.ILogger.LogSuccess(core.GetFuncName())
	return res, nil
}
