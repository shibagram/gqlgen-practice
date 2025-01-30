package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// マイグレーション
	DB.AutoMigrate(&User{})

	// サンプルデータの挿入
	seedUser()
}

type User struct {
	ID    int    `gorm:"primaryKey"`
	Name  string
	Email string
}

// seedUser関数でデータを1件追加
func seedUser() {
	user := User{Name: "Alice", Email: "alice@example.com"}

	// レコードがすでに存在するか確認し、無ければ挿入する
	if err := DB.FirstOrCreate(&user, User{Name: "Alice"}).Error; err != nil {
		log.Println("Failed to insert user:", err)
	} else {
		log.Println("User successfully inserted:", user)
	}
}