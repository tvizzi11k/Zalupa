package apiongo

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var jwtSecret []byte

func generateSecretKey() string {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(key)
}

// Config содержит настройки для подкл к бд
type Config struct {
	MySQLUsername string `json:"mysqlUsername"`
	MySQLPassword string `json:"mysqlPassword"`
	MySQLDbname   string `json:"mysqlDbname"`
	MySQLHost     string `json:"mysqlHost"`
	MySQLPort     string `json:"mysqlPort"`
}

// User Предоставляет модель пользователя в бд
type User struct {
	ID      uint    `gorm:"primaryKey"`
	Key     string  `gorm:"unique;not null"`
	Balance float64 `gorm:"default:0"`
}

// Promocode представляет модель промокода в бд
type Promocode struct {
	ID    uint    `gorm:"primaryKey"`
	Code  string  `gorm:"unique;not null"`
	Value float64 `gorm:"not null"`
	Used  bool    `gorm:"default:false"`
}

// LoginRequest запрос на вход
type LoginRequest struct {
	Key string `json:"key"`
}

// BalanceRequest залупня для баланса
type BalanceRequest struct {
	Key     string  `json:"key"`
	Balance float64 `json:"balance"`
}

// PromoRequest для получения блядского промо чтобы менять блядский баланс
type PromoRequest struct {
	Key     string  `json:"key"`
	Balance float64 `json:"balance"`
}

func init() {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		secret = generateSecretKey()
	}

	jwtSecret = []byte(secret)
}

func main() {
	var config Config
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("Not found json", err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal("Cant read json", err)
	}

	r := gin.Default()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySQLUsername, config.MySQLPassword, config.MySQLHost, config.MySQLPort, config.MySQLDbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect", err)
	}

	err = db.AutoMigrate(&User{}, &Promocode{})
	if err != nil {
		log.Fatal("Failed to migrate", err)
	}

	r.POST("/login", func(c *gin.Context) {
		var req LoginRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}
	})

}
