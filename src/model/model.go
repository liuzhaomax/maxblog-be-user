package model

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"maxblog-be-template/src/pb"
)

var ModelSet = wire.NewSet(
	DataSet,
)

type Data struct {
	gorm.Model
	Mobile string `gorm:"index:idx_mobile;unique;varchar(11);not null"`
}

func Model2PB(data *Data) *pb.DataRes {
	dataRes := &pb.DataRes{
		Id:     int32(data.ID),
		Mobile: data.Mobile,
	}
	return dataRes
}
