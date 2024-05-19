package apiongo

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
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
	Code    string  `json:"key"`
	Balance float64 `json:"balance"`
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})

	r.Use(c)

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

	r.POST("/create-user", func(c *gin.Context) {
		var req LoginRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		token, err := jwt.ParseWithClaims(req.Key, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		var user User

		if err := db.FirstOrCreate(&user, User{Key: claims.Id}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created or updated successfully"})
	})

	r.POST("/balance", func(c *gin.Context) {
		var req BalanceRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		if err := db.Model(&User{}).Where("key = ?", req.Key).Update("balance", req.Balance).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "balance updated"})
	})

	r.POST("/apply-promocode", func(c *gin.Context) {
		var req PromoRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		var promo Promocode
		if err := db.Where("code = ? AND used = ?", req.Code, false).First(&promo).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or used promo code"})
			return
		}

		var user User
		if err := db.Where("key = ?", req.Code).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}

		user.Balance += promo.Value
		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}

		promo.Used = true
		if err := db.Save(&promo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "promocode applied"})
	})

	r.Run(":8080")

}
