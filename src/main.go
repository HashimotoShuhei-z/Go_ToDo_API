package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "net/http"
    "os"
)

// TODO：envが読み込めていない

// 構造体としてテーブル作成
type Todo struct {
    ID     uint   `json:"id" gorm:"primaryKey"` //JSONシリアル化時に、このフィールドが id として表現されることを指定
    Title  string `json:"title"`
    Status string `json:"status"`
}

var DB *gorm.DB

// データベース接続を初期化し、テーブルのマイグレーションを行う
func initDatabase() {
	// データソース名の構築
	// 環境変数を取得するGetenv関数。Go言語の標準ライブラリに含まれる os パッケージの関数で、環境変数の値を取得するために使用される。環境変数はdocker-compose.yml内で設定
    dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"

    var err error
	// gorm.Open: GORMを使ってデータベース接続を開く
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// エラーハンドリング: 接続に失敗した場合、panic でプログラムを停止する
    if err != nil {
        panic("failed to connect database")
    }
	// GORMの自動マイグレーション機能を使用。これにより、構造体のフィールドに対応するカラムがデータベーステーブルに自動的に反映
    DB.AutoMigrate(&Todo{})
}

func main() {
    // err := godotenv.Load()
    // if err != nil {
    //     log.Fatalf("Error loading .env file")
    // }

    initDatabase()
    r := gin.Default()

    r.GET("/todos", getTodos)
    r.POST("/todos", createTodo)
    r.PUT("/todos/:id", updateTodo)
    r.DELETE("/todos/:id", deleteTodo)

    r.Run(":8080")
}

func getTodos(c *gin.Context) {
    var todos []Todo
    DB.Find(&todos)
    c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
    var newTodo Todo
    if err := c.BindJSON(&newTodo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    DB.Create(&newTodo)
    c.JSON(http.StatusCreated, newTodo)
}

func updateTodo(c *gin.Context) {
    var todo Todo
    if err := DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    if err := c.BindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    DB.Save(&todo)
    c.JSON(http.StatusOK, todo)
}

func deleteTodo(c *gin.Context) {
    var todo Todo
    if err := DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    DB.Delete(&todo)
    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}