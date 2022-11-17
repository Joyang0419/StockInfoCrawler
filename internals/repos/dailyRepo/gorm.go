package dailyRepo

import (
	"StockInfoCrawler/internals/models"
	"StockInfoCrawler/internals/tools/dbManager"
	"gorm.io/gorm"
	"time"
)

type GormDailyRepo struct {
	DBManager dbManager.InterfaceDBManger
}

func NewGormDailyRepo(DBManager dbManager.InterfaceDBManger) *GormDailyRepo {
	return &GormDailyRepo{
		DBManager: DBManager,
	}
}

func (repo *GormDailyRepo) CreateDailyPrice(channel chan models.DailyModel) bool {
	db := repo.DBManager.ProvideDBConnection().(*gorm.DB)
	for dailyModel := range channel {
		db.Create(dailyModel)
	}
	return true
}

func (repo *GormDailyRepo) CreateCalculateTimestamp(model models.CalculateTimestampModel) bool {
	db := repo.DBManager.ProvideDBConnection().(*gorm.DB)
	db.Create(&model)
	return true
}

func (repo *GormDailyRepo) GetCalculateTimestamp(table string) time.Time {
	db := repo.DBManager.ProvideDBConnection().(*gorm.DB)
	var calculateTimestamp int64
	db.Model(models.CalculateTimestampModel{Table: table}).Pluck("calculate_timestamp", &calculateTimestamp)

	if calculateTimestamp == 0 {
		return time.Time{}
	}
	return time.Unix(calculateTimestamp, 0)
}

func (repo *GormDailyRepo) UpdateDailyPriceCalculateTimestamp(calculateTimestamp int64) bool {
	db := repo.DBManager.ProvideDBConnection().(*gorm.DB)
	var calculateTimestampModel = models.CalculateTimestampModel{
		Table: "daily_price"}
	db.First(&calculateTimestampModel)
	calculateTimestampModel.CalculateTimestamp = calculateTimestamp
	db.Save(&calculateTimestampModel)
	return true
}
