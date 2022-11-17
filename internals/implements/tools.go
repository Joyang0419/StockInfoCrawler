package implements

import (
	"StockInfoCrawler"
	"StockInfoCrawler/internals/tools/dbManager"
	"fmt"
	"gorm.io/driver/mysql"
)

var (
	sqlUrl = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s",
		StockInfoCrawler.Config.Db.UserName,
		StockInfoCrawler.Config.Db.Password,
		StockInfoCrawler.Config.Db.Host,
		StockInfoCrawler.Config.Db.Port,
		StockInfoCrawler.Config.Db.DbName,
	)
	gormSetting = dbManager.NewGORMDBMSetting(mysql.Open(sqlUrl),
		dbManager.DBMaxIdleConns(StockInfoCrawler.Config.Db.MaxIdleConnection),
		dbManager.DBMaxOpenConns(StockInfoCrawler.Config.Db.MaxOpenConnection),
		dbManager.ConnMaxLifeTimeMinutes(StockInfoCrawler.Config.Db.ConnMaxLifetimeMinutes))
	GormDBManager dbManager.InterfaceDBManger = dbManager.NewGormDBManager(gormSetting)
)
