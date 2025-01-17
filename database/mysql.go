package database

import (
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type mysql struct {
	DB *gorm.DB
}

// Init 初始化 MySQL 数据库
func (m *mysql) Init(source string) error {
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
	m.DB, err = gorm.Open(gmysql.New(gmysql.Config{
		DSN:                       source, // DSN data source name
		DefaultStringSize:         256,    // string 类型字段的默认长度
		DisableDatetimePrecision:  true,   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,   // 重命名索引时删除旧索引，然后创建一个新索引
		DontSupportRenameColumn:   true,   // 重命名列时使用 change 重命名
		SkipInitializeWithVersion: true,   // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return err // 连接数据库错误
	}

	if m.DB.Error != nil {
		return m.DB.Error // 连接数据库错误
	}

	sqlDB, err := m.DB.DB() // 获取底层 sql.DB 对象
	if err != nil {
		return err // 获取底层 sql.DB 对象错误
	}

	sqlDB.SetMaxIdleConns(10)           // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接可复用的最大时间

	return nil
}
