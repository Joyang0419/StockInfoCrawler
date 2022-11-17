package basicRepo

import (
	"StockInfoCrawler/internals/models"
	"StockInfoCrawler/internals/tools/dbManager"
	"gorm.io/gorm"
)

type GormBasicRepo struct {
	DBManager     dbManager.InterfaceDBManger
	BasicModel    models.BasicModel
	CategoryModel models.CategoryModel
}

func NewGormBasicRepo(DBManager dbManager.InterfaceDBManger) *GormBasicRepo {
	return &GormBasicRepo{
		DBManager: DBManager,
	}
}

func (repo *GormBasicRepo) CreateCategories(categories []models.CategoryModel) bool {
	db := repo.DBManager.ProvideDBConnection().(*gorm.DB)
	db.Create(categories)
	return true
}

func (repo *GormBasicRepo) CreateBasic(channel chan models.BasicModel) bool {
	db := repo.DBManager.ProvideDBConnection().(*gorm.DB)
	for basicModel := range channel {
		db.Create(&basicModel)
	}

	return true
}

func (repo *GormBasicRepo) GetCategories() (categories []models.CategoryModel) {
	db := repo.DBManager.ProvideDBConnection().(*gorm.DB)
	db.Find(&categories)
	return categories
}

func (repo *GormBasicRepo) GetCategoryIDMapping() map[string]uint {
	mapping := make(map[string]uint)
	categories := repo.GetCategories()
	for _, model := range categories {
		mapping[model.Name] = model.ID
	}
	return mapping
}

func (repo *GormBasicRepo) GetStockCodes() []string {
	db := repo.DBManager.ProvideDBConnection().(*gorm.DB)
	var stockCodes []string
	db.Model(&models.BasicModel{}).Pluck("stock_code", &stockCodes)
	return stockCodes
}
