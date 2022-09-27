package app

import (
	"context"
	"flag"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"maxblog-be-template/internal/conf"
	"maxblog-be-template/internal/core"
	"maxblog-be-template/src/pb"
	"maxblog-be-template/src/service"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type options struct {
	ConfigDir  string
	ConfigFile string
}

type Option func(*options)

func SetConfigFile(configFile string) Option {
	return func(opts *options) {
		opts.ConfigFile = configFile
	}
}

func InitConfig(opts *options) {
	cfg := conf.GetInstanceOfConfig()
	cfg.Load(opts.ConfigFile)
	conf.InitLogger()
	logger.WithFields(logger.Fields{
		"path": opts.ConfigFile,
	}).Info(core.FormatInfo(101))
}

func InitDB() (*gorm.DB, func(), error) {
	logger.Info(core.FormatInfo(108))
	cfg := conf.GetInstanceOfConfig()
	db, clean, err := cfg.NewDB()
	if err != nil {
		logger.WithFields(logger.Fields{
			"失败方法": core.GetFuncName(),
		}).Fatal(core.FormatError(801, err).Error())
		return nil, clean, err
	}
	err = cfg.AutoMigrate(db)
	if err != nil {
		logger.WithFields(logger.Fields{
			"失败方法": core.GetFuncName(),
		}).Fatal(core.FormatError(802, err).Error())
		return nil, clean, err
	}
	return db, clean, err
}

func InitServer(ctx context.Context, service *service.BData) func() {
	cfg := conf.GetInstanceOfConfig()
	host := flag.String("host", cfg.Server.Host, "Enter host")
	port := flag.Int("port", cfg.Server.Port, "Enter port")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	server := grpc.NewServer()
	pb.RegisterDataServiceServer(server, service)
	go func() {
		listen, err := net.Listen("tcp", addr)
		if err != nil {
			logger.WithFields(logger.Fields{
				"失败方法": core.GetFuncName(),
			}).Fatal(core.FormatError(904, err).Error())
		}
		logger.WithContext(ctx).Infof("Server is running at %s", addr)
		err = server.Serve(listen)
		if err != nil {
			logger.WithFields(logger.Fields{
				"失败方法": core.GetFuncName(),
			}).Fatal(core.FormatError(903, err).Error())
		}
	}()
	return func() {
		logger.Info(core.FormatInfo(103))
		_, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.Server.ShutdownTimeout))
		defer cancel()
		server.Stop()
		logger.Info(core.FormatInfo(107))
	}
}

func Init(ctx context.Context, opts ...Option) func() {
	// initialising options
	options := options{}
	for _, opt := range opts {
		opt(&options)
	}
	// init config
	InitConfig(&options)
	logger.Info(core.FormatInfo(102))
	// init injector and DB
	injector, injectorClean, _ := InitInjector()
	cfg := conf.GetInstanceOfConfig()
	logger.WithFields(logger.Fields{
		"db_type":   cfg.DB.Type,
		"db_name":   cfg.Mysql.DBName,
		"user_name": cfg.Mysql.UserName,
		"host":      cfg.Mysql.Host,
		"port":      cfg.Mysql.Port,
	}).Info(core.FormatInfo(109))
	// init server
	serverClean := InitServer(ctx, injector.Service)
	return func() {
		serverClean()
		injectorClean()
	}
}

func Launch(ctx context.Context, opts ...Option) {
	clean := Init(ctx, opts...)
	cfg := conf.GetInstanceOfConfig()
	logger.WithFields(logger.Fields{
		"app_name": cfg.App.AppName,
		"version":  cfg.App.Version,
		"pid":      os.Getpid(),
		"host":     cfg.Server.Host,
		"port":     cfg.Server.Port,
	}).Info(core.FormatInfo(106))
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
LOOP:
	for {
		sig := <-sc
		logger.WithContext(ctx).Infof("%s [%s]", core.FormatInfo(105), sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break LOOP
		case syscall.SIGHUP:
		default:
			break LOOP
		}
	}
	defer logger.WithContext(ctx).Infof(core.FormatInfo(104))
	defer time.Sleep(time.Second)
	defer os.Exit(state)
	defer clean()
}
