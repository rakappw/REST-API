package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Province struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type APIResponse struct {
	Code    int        `json:"code"`
	Data    []Province `json:"data"`
	Message string     `json:"message"`
	Status  string     `json:"status"`
}

var DB *gorm.DB

func initDB() {
	dsn := "root:@tcp(localhost:3306)/wilayahs?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&Province{})
}

func fetchAndStoreProvinces() {
	client := resty.New()
	resp, err := client.R().Get("https://emsifa.github.io/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		log.Println("Failed to fetch API:", err)
		return
	}

	var result APIResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		log.Println("Failed to decode JSON:", err)
		return
	}

	var provinces []Province
	for _, p := range result.Data {
		id, err := strconv.Atoi(strings.TrimSpace(fmt.Sprintf("%v", p.ID)))
		if err != nil {
			log.Println("Invalid ID format:", p.ID)
			continue
		}
		provinces = append(provinces, Province{
			ID:   uint(id),
			Code: p.Code,
			Name: p.Name,
		})
	}

	DB.Create(&provinces)
	log.Println("Successfully stored provinces")
}

func getProvinces(c *gin.Context) {
	var provinces []Province
	DB.Find(&provinces)
	c.JSON(http.StatusOK, gin.H{"data": provinces})
}

func main() {
	initDB()
	fetchAndStoreProvinces()

	r := gin.Default()
	r.GET("/provinces", getProvinces)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
