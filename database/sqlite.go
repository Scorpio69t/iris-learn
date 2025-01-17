package database

import (
	gsqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type sqlite struct {
	DB *gorm.DB
}

// Init 初始化 SQLite 数据库
func (s *sqlite) Init(source string) error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // Log level
			Colorful:                  false,         // 禁用彩色打印
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
		},
	)

	var err error
	s.DB, err = gorm.Open(gsqlite.New(gsqlite.Config{
		DSN: source,
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return err
	}

	if s.DB.Error != nil {
		return s.DB.Error // 连接数据库错误
	}

	sqlDB, err := s.DB.DB() // 获取底层 sql.DB 对象
	if err != nil {
		return err // 获取底层 sql.DB 对象错误
	}

	sqlDB.SetMaxIdleConns(10)           // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接可复用的最大时间

	return nil
}
