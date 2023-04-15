package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"log"
	"xorm.io/xorm"
)

var Engine = initMysql()
var RDB = initRedis()

func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
func initMysql() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:Rhw000427@tcp(localhost:3306)/cloud_disk")
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}
	return engine

}
