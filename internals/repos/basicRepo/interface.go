package basicRepo

import "StockInfoCrawler/internals/models"

type InterfaceBasicRepo interface {
	CreateCategories(categories []models.CategoryModel) bool
	CreateBasic(channel chan models.BasicModel) bool
	GetCategories() []models.CategoryModel
	GetCategoryIDMapping() map[string]uint
	GetStockCodes() []string
}
