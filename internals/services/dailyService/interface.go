package dailyService

import (
	"StockInfoCrawler/internals/models"
	"time"
)

type InterfaceDailyService interface {
	ScrapeDaily(stockCode string, fromDate time.Time, toDate time.Time, channel chan models.DailyModel) (err error)
}
