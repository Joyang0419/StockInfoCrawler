package implements

import (
	"StockInfoCrawler/internals/services/basicService"
	"StockInfoCrawler/internals/services/dailyService"
)

var (
	DailyService dailyService.InterfaceDailyService = dailyService.NewAuneService()
	BasicService basicService.InterfaceBasicService = basicService.NewTwseService()
)
