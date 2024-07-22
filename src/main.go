package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "net/http"
    "os"
)

// 構造体としてテーブル作成
type Task struct {
    gorm.Model
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
    DB.AutoMigrate(&Task{})
}

func main() {
    initDatabase()
    r := gin.Default()

    r.GET("/tasks", getTasks)
    r.POST("/tasks", createTask)
    r.PUT("/tasks/:id", updateTask)
    r.DELETE("/tasks/:id", deleteTask)
    r.GET("/tasks/:id", showTask)
    r.GET("/tasks/status", getTasksByStatus)

    r.Run(":8080")
}

func getTasks(c *gin.Context) {
    var task []Task
    DB.Find(&task)
    c.JSON(http.StatusOK, task)
}

func createTask(c *gin.Context) {
    var newTask Task
    if err := c.BindJSON(&newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    DB.Create(&newTask)
    c.JSON(http.StatusCreated, newTask)
}

func updateTask(c *gin.Context) {
    var task Task
    if err := DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    if err := c.BindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    DB.Save(&task)
    c.JSON(http.StatusOK, task)
}

func deleteTask(c *gin.Context) {
    var task Task
    if err := DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    DB.Delete(&task)
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func showTask(c *gin.Context) {
    var task Task
    if err := DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"task": task})
}

//getTasksの中にまとめることもできそう。
func getTasksByStatus(c *gin.Context) {
    var tasks []Task
    status := c.Query("status")

    if err := DB.Where("status = ?", status).Find(&tasks).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": tasks})
}