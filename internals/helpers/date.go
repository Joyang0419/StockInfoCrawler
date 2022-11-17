package helpers

import (
	"StockInfoCrawler"
	"time"
)

func GetDate(t time.Time) time.Time {
	currentDateStr := t.Format(StockInfoCrawler.DateStringFormat)
	currentDate, _ := time.ParseInLocation(
		StockInfoCrawler.DateStringFormat, currentDateStr, StockInfoCrawler.GreenwichLocation)
	return currentDate
}
