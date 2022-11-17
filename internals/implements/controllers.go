package implements

import (
	"StockInfoCrawler/internals/controllers/basicController"
	"StockInfoCrawler/internals/controllers/dailyController"
)

var (
	DailyController dailyController.InterfaceDailyController = dailyController.NewControllerDaily(
		DailyService, DailyRepo)

	BasicController basicController.InterfaceDailyController = basicController.NewControllerBasic(
		BasicService, BasicRepo)
)
