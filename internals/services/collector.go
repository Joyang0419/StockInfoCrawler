package services

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

type Headers struct {
	ContentType string `json:"Content-Type"`
	UserAgent   string `json:"User-Agent"`
}

func NewCollector(headers *Headers) (collector *colly.Collector) {
	collector = colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		var headersMap map[string]string
		headersJson, _ := json.Marshal(headers)
		err := json.Unmarshal(headersJson, &headersMap)

		if err != nil {
			log.Error("Headers Unmarshal failed")
		}

		for key, value := range headersMap {
			r.Headers.Set(key, value)
		}
	})

	collector.OnError(func(r *colly.Response, err error) {
		log.Error(fmt.Sprintf("Error %s: %v", r.Request.URL, err))
	})
	return collector
}
