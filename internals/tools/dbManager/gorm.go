package dbManager

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type DBDialector gorm.Dialector
type DBMaxIdleConns int
type DBMaxOpenConns int
type ConnMaxLifeTimeMinutes int

type GORMDBMSetting struct {
	Dialector              DBDialector
	DBMaxIdleConns         DBMaxIdleConns
	DBMaxOpenConns         DBMaxOpenConns
	ConnMaxLifeTimeMinutes ConnMaxLifeTimeMinutes
}

type GormDBManager struct {
	Settings   *GORMDBMSetting
	SqlSession *gorm.DB
}

func (Manager *GormDBManager) Init() {
	sqlSession, err := gorm.Open(Manager.Settings.Dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	Manager.SqlSession = sqlSession

	if err != nil {
		panic("DataBase Connection Failed")
	}

	sqlDB, _ := sqlSession.DB()
	sqlDB.SetMaxIdleConns(int(Manager.Settings.DBMaxIdleConns))
	sqlDB.SetMaxOpenConns(int(Manager.Settings.DBMaxOpenConns))
	sqlDB.SetConnMaxLifetime(time.Duration(Manager.Settings.ConnMaxLifeTimeMinutes) * time.Minute)
}

func (Manager *GormDBManager) IsConnected() bool {
	_, err := Manager.SqlSession.DB()
	if err != nil {
		fmt.Println("get db failed:", err)
		return false
	}
	return true
}

func (Manager *GormDBManager) ProvideDBConnection() any {
	return Manager.SqlSession
}

func NewGORMDBMSetting(dialector DBDialector, dbMaxIdleConns DBMaxIdleConns,
	dbMaxOpenConns DBMaxOpenConns, ConnMaxLifeTimeMinutes ConnMaxLifeTimeMinutes) *GORMDBMSetting {
	return &GORMDBMSetting{Dialector: dialector,
		DBMaxIdleConns:         dbMaxIdleConns,
		DBMaxOpenConns:         dbMaxOpenConns,
		ConnMaxLifeTimeMinutes: ConnMaxLifeTimeMinutes}
}

func NewGormDBManager(dbmSetting *GORMDBMSetting) *GormDBManager {
	Manager := GormDBManager{Settings: dbmSetting}
	Manager.Init()

	return &Manager
}
