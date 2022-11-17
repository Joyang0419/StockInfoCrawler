package dailyController

import "time"

type InterfaceDailyController interface {
	ScrapeDaily(stockCodes []string, fromDate time.Time, toDate time.Time)
}
