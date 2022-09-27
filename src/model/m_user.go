package model

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"maxblog-be-user/internal/core"
	"maxblog-be-user/src/pb"
)

var UserSet = wire.NewSet(wire.Struct(new(MUser), "*"))

type MUser struct {
	DB *gorm.DB
}

func (mUser *MUser) QueryUserById(req *pb.IdRequest, user *User) error {
	result := mUser.DB.First(user, req.Id)
	if result.RowsAffected == 0 {
		return core.FormatError(803, nil)
	}
	return nil
}

func (mUser MUser) QueryLoginByMobile(req *pb.LoginRequest, user *User) error {
	result := mUser.DB.First(user, req.Mobile)
	if result.RowsAffected == 0 {
		return core.FormatError(803, nil)
	}
	return nil
}
