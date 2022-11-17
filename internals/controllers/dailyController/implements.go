package dailyController

import (
	"StockInfoCrawler"
	"StockInfoCrawler/internals/models"
	"StockInfoCrawler/internals/repos/dailyRepo"
	"StockInfoCrawler/internals/services/dailyService"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type ControllerDaily struct {
	Service dailyService.InterfaceDailyService
	Repo    dailyRepo.InterfaceDailyRepo
}

func NewControllerDaily(Service dailyService.InterfaceDailyService, Repo dailyRepo.InterfaceDailyRepo) (
	controller *ControllerDaily) {
	controller = &ControllerDaily{
		Service: Service,
		Repo:    Repo,
	}
	return controller
}

func (controller *ControllerDaily) ScrapeDaily(stockCodes []string, fromDate time.Time, toDate time.Time) {
	stockCodeLength := len(stockCodes)
	log.Info(fmt.Sprintf("StockCodeLength: %d, FromDate: %s, ToDate: %s", stockCodeLength, fromDate, toDate))
	currentTime := time.Now().In(StockInfoCrawler.GreenwichLocation)

	if fromDate.After(currentTime) {
		log.Error("fromDate more than currentTime, Stop Scraping.")
		return
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < len(stockCodes); i++ {
		DailyModelChannel := make(chan models.DailyModel)
		stockCode := stockCodes[i]
		if i%10 == 0 {
			log.Info(fmt.Sprintf("休息: %d秒", StockInfoCrawler.DailyPriceSleepSeconds))
			time.Sleep(StockInfoCrawler.DailyPriceSleepSeconds * time.Second)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := controller.Service.ScrapeDaily(stockCode, fromDate, toDate, DailyModelChannel)
			if err != nil {
				log.Panic(err)
			}
			controller.Repo.CreateDailyPrice(DailyModelChannel)
		}()
	}

	if dailyPriceCalculateDate :=
		controller.Repo.GetCalculateTimestamp("daily_price"); dailyPriceCalculateDate.IsZero() {
		controller.Repo.CreateCalculateTimestamp(models.CalculateTimestampModel{Table: "daily_price",
			CalculateTimestamp: fromDate.Unix()})
	} else {
		controller.Repo.UpdateDailyPriceCalculateTimestamp(fromDate.Unix())
	}

	log.Info(fmt.Sprintf("dailyPriceCalculateDate: %s", fromDate))
	wg.Wait()
	log.Info("Done.")
}
