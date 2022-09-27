package service

import (
	"context"
	"github.com/google/wire"
	"maxblog-be-template/internal/core"
	"maxblog-be-template/src/model"
	"maxblog-be-template/src/pb"
)

var DataSet = wire.NewSet(wire.Struct(new(BData), "*"))

type BData struct {
	MData   *model.MData
	Tx      *core.Trans
	ILogger core.ILogger
}

func (bData *BData) GetDataById(ctx context.Context, req *pb.IdRequest) (*pb.DataRes, error) {
	var data model.Data
	err := bData.Tx.ExecTrans(ctx, func(ctx context.Context) error {
		err := bData.MData.QueryDataById(req, &data)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		bData.ILogger.LogFailure(core.GetFuncName(), err)
		return nil, err
	}
	res := model.Model2PB(&data)
	bData.ILogger.LogSuccess(core.GetFuncName())
	return res, nil
}
