package handler

import (
	"encoding/json"
	"fmt"
	"just-shopping-around-server/internal/model"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)





func GetNews(c echo.Context) error{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		log.Fatal("API_KEY is not set")
	}
	
	resp, err := http.Get(fmt.Sprintf("https://newsapi.org/v2/everything?q=groceries&sortBy=publishedAt&apiKey=%s", apiKey))
		if  err != nil {
			return err	
		}
	defer resp.Body.Close()
	
	var newsFeed []model.NewsAPIResponse
	
	if err := json.NewDecoder(resp.Body).Decode(&newsFeed); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to parse news API response",
		})
	}
	
	
	return c.JSON(http.StatusOK, newsFeed)
}

func Auth(c echo.Context)error{
	message := &model.MessageText{}

	if err := c.Bind(message); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, message)
}