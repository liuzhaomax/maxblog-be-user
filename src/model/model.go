package model

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"maxblog-be-user/src/pb"
)

var ModelSet = wire.NewSet(
	UserSet,
)

type User struct {
	gorm.Model
	Mobile   string `gorm:"index:idx_mobile;unique;varchar(11);not null"`
	Password string `gorm:"type:varchar(32);not null"`
	NickName string `gorm:"type:varchar(32);unique"`
	Salt     string `gorm:"type:varchar(16)"`
	Role     uint32 `gorm:"type:int;default:1;comment:'1-普通用户，2-管理员'"`
}

func Model2PB(user *User) *pb.UserRes {
	userRes := &pb.UserRes{
		Id:       uint32(user.ID),
		Mobile:   user.Mobile,
		NickName: user.NickName,
		Role:     user.Role,
	}
	return userRes
}

func PB2Model(userReq *pb.CreateUserRequest) *User {
	var user User
	user.Mobile = userReq.Mobile
	user.Password = userReq.Password
	user.NickName = userReq.NickName
	user.Salt = userReq.Salt
	return &user
}
