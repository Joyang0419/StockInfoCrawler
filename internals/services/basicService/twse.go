package basicService

import (
	"StockInfoCrawler/internals/helpers"
	"StockInfoCrawler/internals/models"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

type TwseService struct {
}

func NewTwseService() *TwseService {
	return &TwseService{}
}

func (service *TwseService) ScrapeCategory() (categories []models.CategoryModel, err error) {
	res, err := http.Get("https://isin.twse.com.tw/isin/C_public.jsp?strMode=2")
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	b, err := io.ReadAll(res.Body)
	// goquery 限定需為 UTF-8
	b, _ = helpers.DecodeBig5(b)
	reader := bytes.NewReader(b)

	document, err := goquery.NewDocumentFromReader(reader) // Load the HTML document

	if err != nil {
		log.Fatal(err)
	}
	keys := make(map[string]bool)
	document.Find("body > table.h4 > tbody > tr").EachWithBreak(func(i int, tableSelector *goquery.Selection) bool {
		if categorySelection := tableSelector.Find("td > b"); len(categorySelection.Nodes) > 0 {
			if categoryStr := strings.Replace(categorySelection.Text(), " ", "", -1); categoryStr == "上市認購(售)權證" {
				return false
			}
		}

		stockInfoSlice := strings.SplitN(tableSelector.Find("td:nth-child(1)").Text(), "　", 2)
		if len(stockInfoSlice) == 2 {
			stockCategory := strings.Replace(tableSelector.Find("td:nth-child(5)").Text(), " ", "", -1)
			if _, ok := keys[stockCategory]; !ok {
				keys[stockCategory] = true
				categories = append(categories, models.CategoryModel{Name: stockCategory})
			}
		}
		return true
	})

	return categories, nil
}
func (service *TwseService) ScrapeBasic(categoryIDMapping map[string]uint, channel chan models.BasicModel) (err error) {
	res, err := http.Get("https://isin.twse.com.tw/isin/C_public.jsp?strMode=2")
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	b, err := io.ReadAll(res.Body)
	// goquery 限定需為 UTF-8
	b, _ = helpers.DecodeBig5(b)
	reader := bytes.NewReader(b)

	document, err := goquery.NewDocumentFromReader(reader) // Load the HTML document

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		document.Find("body > table.h4 > tbody > tr").EachWithBreak(func(i int, tableSelector *goquery.Selection) bool {
			if categorySelection := tableSelector.Find("td > b"); len(categorySelection.Nodes) > 0 {
				if categoryStr := strings.Replace(categorySelection.Text(), " ", "", -1); categoryStr == "上市認購(售)權證" {
					return false
				}
			}

			stockInfoSlice := strings.SplitN(tableSelector.Find("td:nth-child(1)").Text(), "　", 2)
			if len(stockInfoSlice) == 2 {
				stockCode := strings.Replace(stockInfoSlice[0], " ", "", -1)
				stockName := strings.Replace(stockInfoSlice[1], " ", "", -1)
				stockCategory := strings.Replace(tableSelector.Find("td:nth-child(5)").Text(), " ", "", -1)
				basicModel := models.BasicModel{
					CategoryID: categoryIDMapping[stockCategory],
					StockCode:  stockCode,
					StockName:  stockName,
				}
				channel <- basicModel
			}
			return true
		})
		defer close(channel)
	}()
	return err
}
