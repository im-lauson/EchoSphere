package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var RedisClient *redis.Client
var Logger *log.Logger

type InitConfig struct {
	// TODO: 实现初始化配置的逻辑
}

func (init *InitConfig) InitLogger() {
	// 创建一个日志文件，如果文件不存在则创建，追加写入方式
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	// 创建一个日志记录器，并设置输出到文件
	Logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (init *InitConfig) InitDB() {
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(127.0.0.1:3306)/web_scrapy?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// 设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to get database instance: " + err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

func (init *InitConfig) InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 访问密码，如果没有密码则留空
		DB:       0,                // 使用的 Redis 数据库，默认为 0
	})

	// 测试连接是否成功
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %s", err))
	}
}
