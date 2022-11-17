package dailyRepo

import (
	"StockInfoCrawler/internals/models"
	"time"
)

type InterfaceDailyRepo interface {
	CreateDailyPrice(channel chan models.DailyModel) bool
	CreateCalculateTimestamp(model models.CalculateTimestampModel) bool
	GetCalculateTimestamp(table string) time.Time
	UpdateDailyPriceCalculateTimestamp(calculateTimestamp int64) bool
}
