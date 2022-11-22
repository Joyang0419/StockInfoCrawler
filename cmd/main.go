package main

import (
	"StockInfoCrawler"
	_ "StockInfoCrawler/inits"
	"StockInfoCrawler/internals/helpers"
	"StockInfoCrawler/internals/implements"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

func main() {
	if categories := implements.BasicRepo.GetCategories(); len(categories) == 0 {
		implements.BasicController.ScrapeCategory()
	}

	if stockCodes := implements.BasicRepo.GetStockCodes(); len(stockCodes) == 0 {
		implements.BasicController.ScrapeBasic()
	}

	c := cron.New(
		cron.WithLocation(StockInfoCrawler.TaiwanLocation),
		cron.WithSeconds(),
		cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))

	_, err := c.AddFunc("0 0 14 * * *", func() {
		stockCodes := implements.BasicRepo.GetStockCodes()
		dailyPriceCalculateDate := implements.DailyRepo.GetCalculateTimestamp("daily_price")
		var fromDate, toDate time.Time

		if dailyPriceCalculateDate.IsZero() {
			toDate, _ = time.ParseInLocation(
				StockInfoCrawler.DateStringFormat, StockInfoCrawler.Todate, StockInfoCrawler.GreenwichLocation)
			fromDate = helpers.GetDate(time.Now().In(StockInfoCrawler.GreenwichLocation))
		} else {
			fromDate = helpers.GetDate(time.Now().In(StockInfoCrawler.GreenwichLocation))
			toDate = dailyPriceCalculateDate.AddDate(0, 0, 1).In(StockInfoCrawler.GreenwichLocation)
		}
		implements.DailyController.ScrapeDaily(stockCodes, fromDate, toDate)
	})
	if err != nil {
		return
	}
	c.Run()
}
