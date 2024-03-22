package internal

import (
	"context"
	"fmt"
	"gin-demo/internal/config"
	"gin-demo/internal/router"
	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io"
	"os"
	"time"

	"strings"
)

func Exec() *gin.Engine {
	InitConfig()     // 初始化配置
	InitLogger()     // 初始化日志
	InitMysql()      // 初始化Mysql
	InitRedis()      // 初始化Redis
	InitLocalCache() // 初始化本地缓存
	return router.InitRouter()
}

func InitConfig() {
	configFileName := "./etc/config.yaml"
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf("读取配置文件失败%v", err)
	}
	c := config.GetServerConfig()
	if err := v.Unmarshal(&c); err != nil {
		zap.S().Panicf("解析配置文件失败 %v", err)
	}

	config.SetServerConfig(c)
	fmt.Printf("配置文件：%+v", c)
}

func InitMysql() {
	c := config.GetServerConfig().MysqlConf
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DB)
	//Gorm 有一个 默认 logger 实现，默认情况下，它会打印慢 SQL 和错误
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.serverConfig{
	//		SlowThreshold: time.Second,   // 慢 SQL 阈值
	//		LogLevel:      logger.Silent, // Log level
	//		Colorful:      true,          // 禁用彩色打印
	//	},
	//)

	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//Logger: newLogger,
	})
	config.SetDb(db)

	if err != nil {
		zap.S().Panicf("初始化数据库失败 err:%v", err)
	}
}

func InitRedis() {
	ctx := context.Background()
	c := config.GetServerConfig()
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    strings.Split(c.RedisConf.Host, ","),
		Password: c.RedisConf.Password,
		// To route commands by latency or randomly, enable one of the following.
		//RouteByLatency: true,
		//RouteRandomly: true,
	})

	err := redisClient.Ping(ctx).Err()
	if err != nil {
		zap.S().Panicf("初始化Redis失败 err:%+v", err)
	}
	config.SetRedis(redisClient)
}

func InitLocalCache() {
	interval := 8760 * 100 * time.Hour
	c := bigcache.DefaultConfig(interval)
	var err error
	ctx := context.Background()
	localCache, err := bigcache.New(ctx, c)
	if err != nil {
		zap.S().Panicf("初始化本地缓存失败 err:%+v", err)
	}
	config.SetLocalCache(localCache)
}

func InitLogger() {
	encoder := getEncoder()
	loggerInfo := getLogWriterInfo()
	logLevel := zapcore.InfoLevel
	switch config.GetServerConfig().LogConf.Level {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	}

	coreInfo := zapcore.NewCore(encoder, loggerInfo, logLevel)
	logger := zap.New(coreInfo)
	zap.ReplaceGlobals(logger)
}

func getEncoder() zapcore.Encoder {
	productionEncoderConfig := zap.NewProductionEncoderConfig()
	productionEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	productionEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(productionEncoderConfig)
}

func getLogWriterInfo() zapcore.WriteSyncer {
	logPath := config.GetServerConfig().LogConf.Path + "/" + config.GetServerConfig().Name + ".log"
	l := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    config.GetServerConfig().LogConf.MaxSize,    //最大MB
		MaxBackups: config.GetServerConfig().LogConf.MaxBackups, //最大备份
		MaxAge:     config.GetServerConfig().LogConf.MaxAge,     //保留7天
		Compress:   true,
	}

	var ws io.Writer
	if config.GetServerConfig().Mode == "release" {
		ws = io.MultiWriter(l)
	} else {
		//如果不是开发环境，那么会打印日志到日志文件和标准输出，也就是控制台
		ws = io.MultiWriter(l, os.Stdout)
	}

	return zapcore.AddSync(ws)
}
