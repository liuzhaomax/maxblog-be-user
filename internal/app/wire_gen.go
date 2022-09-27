// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"maxblog-be-user/internal/core"
	"maxblog-be-user/src/model"
	"maxblog-be-user/src/service"
)

// Injectors from wire.go:

func InitInjector() (*Injector, func(), error) {
	db, cleanup, err := InitDB()
	if err != nil {
		return nil, nil, err
	}
	mUser := &model.MUser{
		DB: db,
	}
	trans := &core.Trans{
		DB: db,
	}
	logger := &core.Logger{}
	bUser := &service.BUser{
		MUser:   mUser,
		Tx:      trans,
		ILogger: logger,
	}
	injector := &Injector{
		DB:      db,
		Service: bUser,
	}
	return injector, func() {
		cleanup()
	}, nil
}
