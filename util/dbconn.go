package yizuutil

import (
	"yizu/conf"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {

	dsn := "host=" + conf.ServerConfig().PgConfig.Address + " " +
		"user=" + conf.ServerConfig().PgConfig.Username + " " +
		"password=" + conf.ServerConfig().PgConfig.Password + " " +
		"dbname=" + conf.ServerConfig().PgConfig.DBName + " " +
		"port=" + conf.ServerConfig().PgConfig.Port + " " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("PostgreSQL连接失败: %v", err)
	}
	db = db.Debug()
	return db, err
}

func GetRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.ServerConfig().RdConfig.Address,
		Password: conf.ServerConfig().RdConfig.Password,
		DB:       conf.ServerConfig().RdConfig.DB,
	})
	return rdb
}
