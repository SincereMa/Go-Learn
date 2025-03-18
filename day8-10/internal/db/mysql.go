package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
	"github.com/spf13/viper"
)

type Config struct {
	DSN string
}

func NewDBConfig() (Config, error) {
	cfg := Config{
		DSN: viper.GetString("db.dsn"), // 从 Viper 读取配置
	}
	fmt.Println("db_cfg, ", cfg)
	return cfg, nil
}

func NewDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
