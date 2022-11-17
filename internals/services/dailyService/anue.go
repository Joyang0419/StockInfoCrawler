package dailyService

import (
	"StockInfoCrawler/internals/models"
	"StockInfoCrawler/internals/services"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"time"
)

type AuneService struct {
}

func (service *AuneService) ScrapeDaily(stockCode string, fromDate time.Time, toDate time.Time,
	channel chan models.DailyModel) (err error) {

	headers := services.Headers{ContentType: "application/json",
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) " +
			"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"}

	collector := services.NewCollector(&headers)

	collector.OnScraped(func(r *colly.Response) {
		var scrapedData map[string]any
		if err := json.Unmarshal(r.Body, &scrapedData); err != nil {
			log.Panic(fmt.Sprintf("error: %v", err))
		}

		responseData := scrapedData["data"].(map[string]any)
		timestampArray := responseData["t"].([]any)
		volumeArray := responseData["v"].([]any)
		closingArray := responseData["c"].([]any)
		openingArray := responseData["o"].([]any)
		highestArray := responseData["h"].([]any)
		lowestArray := responseData["l"].([]any)

		go func() {
			for index := len(timestampArray) - 1; index >= 0; index-- {
				dailyModel := models.DailyModel{
					StockCode:     stockCode,
					Volume:        decimal.NewFromFloat(volumeArray[index].(float64)),
					OpeningPrice:  decimal.NewFromFloat(openingArray[index].(float64)),
					ClosingPrice:  decimal.NewFromFloat(closingArray[index].(float64)),
					HighestPrice:  decimal.NewFromFloat(highestArray[index].(float64)),
					LowestPrice:   decimal.NewFromFloat(lowestArray[index].(float64)),
					CalculateDate: time.Unix(int64(int(timestampArray[index].(float64))), 0),
				}
				channel <- dailyModel
			}
			defer close(channel)
		}()
	})

	url := fmt.Sprintf("https://ws.api.cnyes.com/ws/api/v1/charting/history?"+
		"resolution=D&quote=1"+
		"&symbol=TWS:%s:STOCK"+
		"&from=%d&to=%d", stockCode, fromDate.Unix(), toDate.Unix())

	err = collector.Visit(url)

	return err
}

func NewAuneService() *AuneService {
	return &AuneService{}
}
