package main

import (
	"fmt"

	"gin_practice/conf"
	"gin_practice/handler"
	mysql2 "gin_practice/model/mysql"
	"gin_practice/service/item"
	"gin_practice/service/user"
	"gin_practice/tool/infra/mysql"
	"gin_practice/tool/logger"
	"gin_practice/tool/viperx"

	"gin_practice/router"
)

func main() {
	if err := viperx.EnvSetIntoConfig("env", "yaml", "./conf", &conf.Config); err != nil {
		logger.Fatal("viperx.EnvSetIntoConfig err: %v", err)
	}

	// region Infra Conn
	mysqlConn := mysql.SetupConn(
		mysql.Config{
			Host:            conf.Config.MySQL.Host,
			Port:            conf.Config.MySQL.Port,
			Username:        conf.Config.MySQL.Username,
			Password:        conf.Config.MySQL.Password,
			Database:        conf.Config.MySQL.Database,
			MaxIdleConns:    conf.Config.MySQL.MaxIdleConns,
			MaxOpenConns:    conf.Config.MySQL.MaxOpenConns,
			ConnMaxLifeTime: conf.Config.MySQL.ConnMaxLifeTime,
		},
		nil,
	)
	// endregion

	// region query
	userQuery := user.NewQuery(mysqlConn)
	itemQuery := item.NewQuery(mysqlConn)
	// endregion

	// region command
	userCmd := user.NewCMD(mysqlConn)
	itemCmd := item.NewCMD(mysqlConn)
	// endregion

	// region service
	userSvc := user.NewService(userQuery, userCmd)
	itemSvc := item.NewService(itemQuery, itemCmd)
	// endregion

	// TODO Delete, test proj, 先快速驗證
	// 自動遷移架構
	if err := mysqlConn.AutoMigrate(new(mysql2.Item), new(mysql2.User)); err != nil {
		logger.Fatal("[main]mysqlConn.AutoMigrate err: %v", err)
	}

	r := router.NewGinRouter(
		nil,
		router.Handler{
			Auth: handler.NewAuth(userSvc),
			Item: handler.NewItem(itemSvc),
		},
	)

	logger.Info("[main]success on port: %d", conf.Config.Server.Port)
	if err := r.Run(fmt.Sprintf(":%d", conf.Config.Server.Port)); err != nil {
		logger.Fatal("[main]r.Run err: %v", err)
	}
}
