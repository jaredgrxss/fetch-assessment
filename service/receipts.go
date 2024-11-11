package service

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// structs
type item struct {
	ShortDescription string `json:"shortDescription"`
	Price string `json:"price"`
}

type receipt struct {
	Retailer string  `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items []item `json:"items"`
	Total string `json:"total"`
	Points int
}

// in memory store to hold receipts, 
// can be lost if server stops
var receipts map[string]receipt = map[string]receipt{}

func GetReceipt(c *gin.Context) {
	receiptId := c.Param("id")
	if receipt, exists := receipts[receiptId]; exists {
		c.JSON(http.StatusOK, gin.H{"points": receipt.Points})
		return 
	} else {
		c.JSON(http.StatusNotFound, gin.H{"description": "No receipt found for that id"})
		return
	}
}

func PostReceipt(c *gin.Context) {
	// generate a uuid for storage
	uuid := uuid.New()
	// bind request body 
	var requestBody receipt
	if err := c.BindJSON(&requestBody); err != nil {
		log.Println("Error in parsing request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"description": "The receipt is invalid"})
		return
	}
	// caculate points
	newReceipt, err := calculatePoints(requestBody)
	if err != nil {
		log.Println("Error in calculating total points of receipt:", err)
		c.JSON(http.StatusBadRequest, gin.H{"description": "The receipt is invalid"})
		return
	}
	// store the receipt with points
	receipts[uuid.String()] = newReceipt 
	// respond with uuid
	c.JSON(http.StatusOK, gin.H{"uuid": uuid})
}

func calculatePoints(data receipt) (receipt, error) {
	// start with 0
	var points int = 0
	// add points for retailer name
	for _, c := range data.Retailer {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			points++ 
		}
	}
	totalFloat, err := strconv.ParseFloat(data.Total, 64) 
	if err != nil {
		return data, err
	}
	totalInt := int(totalFloat)
	if totalFloat == float64(totalInt) { // check if total is a whole number
		points += 50
	}
	if math.Mod(totalFloat, .25) == 0 { // check for divisibilty 
		points += 25
	}
	points += 5 * (len(data.Items) / 2); // every 2 items in receipt gets 5 points
	for _, item := range data.Items {
		if len(strings.TrimSpace(item.ShortDescription)) % 3 == 0 { // check for description length
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return data, err
			}
			points += int(math.Ceil(.2 * price)) // add points after multiplication
		}
	}
	date, err := time.Parse("2006-01-02", data.PurchaseDate)
	if err != nil {
		return data, err
	}
	if date.Day() & 1 == 1 { 
		points += 6 // check for odd day
	}
	if data.PurchaseTime >= "14:00" && data.PurchaseTime <= "16:00" {
		points += 10 // check for promotion time between 2PM and 4PM
	}
	data.Points = points // set points on obj
	return data, nil
}