package main

import (
	"behappy/bot"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

// User Предоставляет модель пользователя в бд
type User struct {
	ID      uint    `gorm:"primaryKey"`
	Key     string  `gorm:"unique;not null"`
	Balance float64 `gorm:"default:0"`
	Admin   bool    `gorm:"default:false"`
}

// Promocode представляет модель промокода в бд
type Promocode struct {
	ID    uint    `gorm:"primaryKey"`
	Code  string  `gorm:"unique;not null"`
	Value float64 `gorm:"not null"`
	Used  int64   `gorm:"not null;default:0"`
	Max   int64   `gorm:"not null"`
}

// BalanceRequest залупня для баланса
type BalanceRequest struct {
	Key     string  `json:"key"`
	Balance float64 `json:"balance"`
}

// PromoRequest для получения блядского промо чтобы менять блядский баланс
type PromoRequest struct {
	Code string `json:"code"`
}

type CreatePromoRequest struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
	Max   int64   `json:"max"`
}

func main() {
	go bot.Run()

	r := gin.Default()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	}))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect", err)
	}

	err = db.AutoMigrate(&User{}, &Promocode{})
	if err != nil {
		log.Fatal("Failed to migrate", err)
	}

	r.POST("/balance", func(c *gin.Context) {
		address := c.GetHeader("Authorization")
		if address == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		var req BalanceRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		var user User
		if err := db.Where("key = ?", address).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}

		user.Balance = req.Balance
		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "balance updated"})
	})

	r.POST("/apply-promocode", func(c *gin.Context) {
		address := c.GetHeader("Authorization")
		if address == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		var req PromoRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
			return
		}

		var promo Promocode
		if err := db.Where("code = ?", req.Code).First(&promo).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid promocode"})
			return
		}

		if promo.Used >= promo.Max {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Promocode is used"})
			return
		}

		var user User
		if err := db.Where("key = ?", address).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}

		user.Balance += promo.Value
		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}

		promo.Used++
		if err := db.Save(&promo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "promocode applied"})

	})

	r.POST("/create-promocode", func(c *gin.Context) {
		address := c.GetHeader("Authorization")
		if address == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		var req CreatePromoRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
			return
		}

		var user User
		if err := db.Where("key = ? and admin = true", address).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found or not admin"})
			return
		}

		promo := Promocode{
			Code:  req.Code,
			Value: req.Value,
			Max:   req.Max,
		}

		if err := db.Save(&promo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "promocode created"})
	})

	r.GET("/get-balance", func(c *gin.Context) {
		address := c.GetHeader("Authorization")
		if address == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		var user User
		if err := db.Where("key = ?", address).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// если юзера ебаного нет то создаем нового
				newUser := User{Key: address, Balance: 0.0}
				if createErr := db.Create(&newUser).Error; createErr != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
					return
				}
				c.JSON(http.StatusOK, gin.H{"balance": newUser.Balance})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"balance": user.Balance})
	})

	r.GET("/", func(c *gin.Context) {
		c.File("./web/index.html")
	})

	r.GET("/home", func(c *gin.Context) {
		c.File("./web/home.html")
	})

	r.GET("/admin", func(c *gin.Context) {
		c.File("./web/admin.html")
	})

	r.Static("/static", "./web/static")
	r.StaticFile("/tonconnect-manifest.json", "./web/tonconnect-manifest.json")

	log.Fatal(r.Run(os.Getenv("HOST")))
}
