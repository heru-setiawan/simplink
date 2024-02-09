package database

import (
	"simplink/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlInit(cfg config.Mysql) (*gorm.DB, error) {
	conn, err := gorm.Open(mysql.Open(cfg.ToDsn()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func MysqlMigrate(conn *gorm.DB, models ...interface{}) error {
	if err := conn.AutoMigrate(models...); err != nil {
		return err
	}

	return nil
}
