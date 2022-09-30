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

func (bUser *BUser) ValidateLoginInfo(ctx context.Context, req *pb.LoginRequest) (*pb.LoginRes, error) {
	var user model.User
	err := bUser.MUser.QueryLoginByMobile(req, &user)
	if err != nil {
		bUser.ILogger.LogFailure(core.GetFuncName(), err)
		return nil, err
	}
	if user.Salt == EmptyStr {
		return nil, core.FormatError(702, nil)
	}
	if user.Password == EmptyStr {
		return nil, core.FormatError(703, nil)
	}
	bUser.ILogger.LogSuccess(core.GetFuncName())
	return &pb.LoginRes{EncodedPwd: user.Password, Salt: user.Salt}, nil
}

func (bUser *BUser) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.SuccessRes, error) {
	user := model.PB2Model(req)
	err := bUser.MUser.CreateUser(user)
	if err != nil {
		bUser.ILogger.LogFailure(core.GetFuncName(), err)
		return nil, err
	}
	return &pb.SuccessRes{Result: true}, nil
}
