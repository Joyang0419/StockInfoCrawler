package StockInfoCrawler

import (
	"StockInfoCrawler/configs"
	"time"
)

var (
	TaiwanLocation, _    = time.LoadLocation("Asia/Taipei")
	GreenwichLocation, _ = time.LoadLocation("")
	Config, _            = configs.LoadConfig(".", "app", "env")
)

const (
	DateStringFormat       = "2006-01-02"
	Todate                 = "2018-01-01"
	DailyPriceSleepSeconds = 5
)
