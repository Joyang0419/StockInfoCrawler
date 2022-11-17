package basicService

import (
	"StockInfoCrawler/internals/models"
)

type InterfaceBasicService interface {
	ScrapeCategory() (categories []models.CategoryModel, err error)
	ScrapeBasic(categoryIDMapping map[string]uint, channel chan models.BasicModel) (err error)
}
