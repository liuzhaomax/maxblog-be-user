//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"maxblog-be-template/internal/core"
	dataModel "maxblog-be-template/src/model"
	dataService "maxblog-be-template/src/service"
)

func InitInjector() (*Injector, func(), error) {
	wire.Build(
		InitDB,
		core.TransSet,
		core.LoggerSet,
		dataModel.ModelSet,
		dataService.ServiceSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
