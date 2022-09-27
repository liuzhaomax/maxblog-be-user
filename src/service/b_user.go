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
	var user model.User
	err := bUser.Tx.ExecTrans(ctx, func(ctx context.Context) error {
		err := bUser.MUser.QueryUserById(req, &user)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		bUser.ILogger.LogFailure(core.GetFuncName(), err)
		return nil, err
	}
	res := model.Model2PB(&user)
	bUser.ILogger.LogSuccess(core.GetFuncName())
	return res, nil
}

func (bUser *BUser) PostLogin(ctx context.Context, req *pb.LoginRequest) (*pb.LoginRes, error) {
	var user model.User
	err := bUser.MUser.QueryLoginByMobile(req, &user)
	if err != nil {
		bUser.ILogger.LogFailure(core.GetFuncName(), err)
		return nil, err
	}
	if req.Password != user.Password {
		bUser.ILogger.LogFailure(core.GetFuncName(), core.FormatError(701, nil))
		return nil, core.FormatError(701, nil)
	}
	var res *pb.LoginRes
	res.Result = true
	bUser.ILogger.LogSuccess(core.GetFuncName())
	return res, nil
}
