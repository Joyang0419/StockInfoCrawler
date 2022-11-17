package implements

import (
	"StockInfoCrawler/internals/repos/basicRepo"
	"StockInfoCrawler/internals/repos/dailyRepo"
)

var (
	DailyRepo dailyRepo.InterfaceDailyRepo = dailyRepo.NewGormDailyRepo(GormDBManager)
	BasicRepo basicRepo.InterfaceBasicRepo = basicRepo.NewGormBasicRepo(GormDBManager)
)
