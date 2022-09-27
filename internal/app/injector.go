package app

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"maxblog-be-template/src/service"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	DB      *gorm.DB
	Service *service.BData
}
