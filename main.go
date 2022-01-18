package main

import (
	"fmt"
	"go_web/config"
	"go_web/dao/mysql"
	"go_web/dao/redis"
	"go_web/logger"
	"go_web/routers"
)

//此处是主入口
func main() {

	//加载配置文件
	if err := config.LoadFromFile("./conf/web.ini"); err != nil {
		fmt.Printf("load config file err:%v\n", err) //当加载配置文件失败的时候，直接打印错误
		return                                       //退出执行,啥也不返回
	}
	//当上述都没错的时候，才会继续往下执行

	//加载logger日志
	if err := logger.Init(config.Conf.LogConfig); err != nil {
		fmt.Printf("logger config file err:%v\n", err)
		return
	}

	//连接数据库
	//连接mysql
	if err := mysql.Init(config.Conf.MysqlConfig); err != nil {
		fmt.Printf("connect mysql config err:%v\n", err)
		return
	}

	//连接redis
	if err := redis.Init(config.Conf.RedisConfig); err != nil {
		fmt.Printf("connect redis config err:%v\n", err)
		return
	}

	//返回一个路由实例
	r := routers.SetupRouter()

	//启动路由
	r.Run()
}
