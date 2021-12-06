package cmd

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/go-kit/kit/log"
	"github.com/infraboard/keyauth/logs"
	"github.com/infraboard/mcube/bus"
	"github.com/infraboard/mcube/bus/broker/kafka"
	"github.com/infraboard/mcube/bus/broker/nats"
	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/cache/memory"
	"github.com/infraboard/mcube/cache/redis"
	_ "github.com/infraboard/mcube/logger"
	_ "github.com/infraboard/mcube/logger/zap"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/searKing/golang/go/os/signal"
	zlog "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"strings"
	"syscall"
	"time"
	//"go.uber.org/zap/zapcore"
	"github.com/spf13/cobra"
	"os"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/protocol"
	// 加载所有服务
	_ "github.com/infraboard/keyauth/pkg/all"
)

var (
	// pusher service config option
	confType string
	confFile string
	Logs     = make(map[string]*zlog.Logger)
)
var CommonLog = logs.Logger{}
var ServerLog = logs.Logger{}

// startCmd represents the start command
var serviceCmd = &cobra.Command{
	Use:   "start",
	Short: "权限中心服务",
	Long:  `权限中心服务`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化全局变量
		if err := loadGlobalConfig(confType); err != nil {
			return err
		}

		// 初始化全局组件
		if err := loadGlobalComponent(); err != nil {
			return err
		}

		// 初始化服务层
		if err := pkg.InitService(); err != nil {
			return err
		}

		conf := conf.C()
		// 启动服务
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

		//初始化服务
		svr, err := newService(conf)
		if err != nil {
			return err
		}

		// 等待信号处理
		go svr.waitSign(ch)

		// 启动服务
		if err := svr.start(); err != nil {
			if !strings.Contains(err.Error(), "http: Server closed") {
				return err
			}
		}

		return nil
	},
}

func newService(cnf *conf.Config) (*service, error) {
	log := CommonLog.Named("CLI")
	http := protocol.NewHTTPService()
	grpc := protocol.NewGRPCService()

	// 初始化总线
	bm, err := newBus()
	if err != nil {
		log.Errorf("new bus error, %s", err)
	}

	svr := &service{
		http: http,
		grpc: grpc,
		bm:   bm,
		log:  log,
	}

	return svr, nil
}

type service struct {
	http *protocol.HTTPService
	grpc *protocol.GRPCService
	bm   bus.Manager

	log  *logs.Logger
	stop context.CancelFunc
}

func (s *service) start() error {
	s.log.Infof("loaded domain pkg: %v", pkg.LoadedService())
	s.log.Infof("loaded http service: %s", pkg.LoadedHTTP())

	if s.bm != nil {
		if err := s.bm.Connect(); err != nil {
			s.log.Errorf("connect bus error, %s", err)
		}
	}

	go s.grpc.Start()
	return s.http.Start()
}

// config 为全局变量, 只需要load 即可全局可用户
func loadGlobalConfig(configType string) error {
	// 配置加载
	switch configType {
	case "file":
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			return err
		}
	case "env":
		err := conf.LoadConfigFromEnv()
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("unknown config type")
	}

	return nil
}

func loadGlobalComponent() error {
	// 初始化全局日志配置
	if err := loadGlobalLogzap(); err != nil {
		return err
	}
	// 加载缓存
	if err := loadCache2(); err != nil {
		return err
	}
	return nil
}

// 将原本的infraboard/mcube/logger 替换为 logrus
func loadGlobalLogzap() error {
	//var (
	//	logInitMsg string
	//	level      zap.Level
	//)

	lc := conf.C().Logs
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		TimeKey:    "ts",
		//CallerKey:      "file",
		CallerKey:     "caller",
		StacktraceKey: "trace",
		LineEnding:    zapcore.DefaultLineEnding,
		// EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		//EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	for key, logConfig := range *lc {
		//lv, _ := logs.NewLevel(logConfig.Level)
		infoLevel := zlog.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.DebugLevel
		})
		hook := getWriter(logConfig.PathDir, logConfig.File)
		Logs[key] = zlog.New(zapcore.NewCore(encoder, zapcore.AddSync(hook), infoLevel),
			zlog.AddCaller(), zlog.AddStacktrace(zlog.ErrorLevel))
	}
	ServerLog = logs.Logger{Logs["serverlog"], Logs["serverlog"].Sugar()}
	CommonLog = logs.Logger{Logs["commonlog"], Logs["commonLog"].Sugar()}
	//defer logger.Sync()
	return nil
}

//func (log *zlog.Logger) Comment(msg string, fields ...zlog.Field) {
//	if ce := log.check(InfoLevel, msg); ce != nil {
//		ce.Write(fields...)
//	}
//}

func getWriter(path string, filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYYY-mm-dd
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		// 没有使用go风格反人类的format格式
		path+"/"+filename+".%Y-%m-%d",
		rotatelogs.WithLinkName(path+"/"+filename),
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(1)),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}

// log 为全局变量, 只需要load 即可全局可用户, 依赖全局配置先初始化
//func loadGlobalLogger() error {
//	var (
//		logInitMsg string
//		level      zap.Level
//	)
//
//	lc := conf.C().Log
//	lv, err := zap.NewLevel(lc.Level)
//	if err != nil {
//		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
//		level = zap.InfoLevel
//	} else {
//		level = lv
//		logInitMsg = fmt.Sprintf("log level: %s", lv)
//	}
//
//	zapConfig := zap.DefaultConfig()
//	zapConfig.Level = level
//
//	switch lc.To {
//	case conf.ToStdout:
//		zapConfig.ToStderr = true
//		zapConfig.ToFiles = false
//	case conf.ToFile:
//		// only one log file
//		zapConfig.Files.Name = "api2.log"
//		zapConfig.Files.Path = lc.PathDir
//	}
//
//	switch lc.Format {
//	case conf.JSONFormat:
//		zapConfig.JSON = true
//	}
//
//	if err := zap.Configure(zapConfig); err != nil {
//		return err
//	}
//
//	zap.L().Named("INIT").Error(logInitMsg)
//	zap.L().Named("TEST").Info(logInitMsg)
//	return nil
//}
func loadCache2() error {
	l := CommonLog.Named("INIT")
	c := conf.C()
	// 设置全局缓存
	switch c.Cache.Type {
	case "memory", "":
		ins := memory.NewCache(c.Cache.Memory)
		cache.SetGlobal(ins)
		l.Info("use cache in local memory")
	case "redis":
		ins := redis.NewCache(c.Cache.Redis)
		cache.SetGlobal(ins)
		l.Info("use redis to cache")
	default:
		return fmt.Errorf("unknown cache type: %s", c.Cache.Type)
	}

	return nil
}

//func loadCache() error {
//	l := zap.L().Named("INIT")
//	c := conf.C()
//	// 设置全局缓存
//	switch c.Cache.Type {
//	case "memory", "":
//		ins := memory.NewCache(c.Cache.Memory)
//		cache.SetGlobal(ins)
//		l.Info("use cache in local memory")
//	case "redis":
//		ins := redis.NewCache(c.Cache.Redis)
//		cache.SetGlobal(ins)
//		l.Info("use redis to cache")
//	default:
//		return fmt.Errorf("unknown cache type: %s", c.Cache.Type)
//	}
//
//	return nil
//}

func newBus() (bus.Manager, error) {
	c := conf.C()
	if c.Nats != nil {
		ns, err := nats.NewBroker(c.Nats)
		if err != nil {
			return nil, err
		}
		bus.SetPublisher(ns)
		return ns, nil
	}

	if c.Kafka != nil {
		ks, err := kafka.NewPublisher(c.Kafka)
		if err != nil {
			return nil, err
		}
		bus.SetPublisher(ks)
		return ks, nil
	}

	return nil, fmt.Errorf("bus not config, nats or kafka required")
}

func (s *service) waitSign(sign chan os.Signal) {
	for {
		select {
		case sg := <-sign:
			switch v := sg.(type) {
			default:
				s.log.Infof("receive signal '%v', start graceful shutdown", v.String())
				if err := s.grpc.Stop(); err != nil {
					s.log.Errorf("grpc graceful shutdown err: %s, force exit", err)
				}
				s.log.Info("grpc service stop complete")
				if err := s.http.Stop(); err != nil {
					s.log.Errorf("http graceful shutdown err: %s, force exit", err)
				}
				s.log.Infof("http service stop complete")
				return
			}
		}
	}
}

func init() {
	serviceCmd.Flags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env/etcd]")
	serviceCmd.Flags().StringVarP(&confFile, "config-file", "f", "etc/keyauth.toml", "the service config from file")
	RootCmd.AddCommand(serviceCmd)
}
