package database

import (
    "log"
    "os"
    "todo-app/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func InitDatabase() {
		// データソース名の構築
	// 環境変数を取得するGetenv関数。Go言語の標準ライブラリに含まれる os パッケージの関数で、環境変数の値を取得するために使用される。環境変数はdocker-compose.yml内で設定
    dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"

    var err error
    models.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    models.DB.AutoMigrate(&models.Task{})
}
