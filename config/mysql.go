package config

import (
	"fmt"
	"os"
	"strconv"
)

type Mysql struct {
	Host     string
	Port     uint16
	Username string
	Password string
	Database string
}

func (cfg *Mysql) ToDsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
}

func (cfg *Mysql) LoadFromEnv() error {
	if val, ok := os.LookupEnv("MYSQL_HOST"); ok {
		cfg.Host = val
	}

	if val, ok := os.LookupEnv("MYSQL_PORT"); ok {
		port, err := strconv.Atoi(val)
		if err != nil {
			return err
		}

		cfg.Port = uint16(port)
	}

	if val, ok := os.LookupEnv("MYSQL_USERNAME"); ok {
		cfg.Username = val
	}

	if val, ok := os.LookupEnv("MYSQL_PASSWORD"); ok {
		cfg.Password = val
	}

	if val, ok := os.LookupEnv("MYSQL_DATABASE"); ok {
		cfg.Database = val
	}

	return nil
}
