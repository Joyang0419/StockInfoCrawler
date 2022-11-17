package models

import (
	"github.com/shopspring/decimal"
	"time"
)

// DailyModel 每日個股價格資訊
type DailyModel struct {
	StockCode     string          // 股票代碼
	Volume        decimal.Decimal //交易量(張)
	OpeningPrice  decimal.Decimal // 開盤價
	ClosingPrice  decimal.Decimal // 收盤價
	HighestPrice  decimal.Decimal // 最高價
	LowestPrice   decimal.Decimal // 最低價
	CalculateDate time.Time       // 價格日期
}

func (DailyModel) TableName() string {
	return "daily_price"
}

// BasicModel 股票基本
type BasicModel struct {
	ID         uint
	CategoryID uint   // 分類代碼
	StockCode  string // 股票代碼
	StockName  string // 股票名稱
}

func (BasicModel) TableName() string {
	return "basic"

}

// CategoryModel 股票分類
type CategoryModel struct {
	ID   uint
	Name string // 分類名稱
}

func (CategoryModel) TableName() string {
	return "category"
}

type CalculateTimestampModel struct {
	ID                 uint
	Table              string
	CalculateTimestamp int64
}

func (CalculateTimestampModel) TableName() string {
	return "calculate_timestamps"
}
