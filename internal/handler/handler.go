package handler

import (
	"encoding/json"
	"fmt"
	"io"
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

	apiURL := os.Getenv("API_URL")

	if apiURL == "" {
		log.Fatal("API_URL is not set")
	}

	
	
	resp, err := http.Get(fmt.Sprintf("%s/v2/everything?q=groceries&sortBy=publishedAt&apiKey=%s",apiURL, apiKey))
		if  err != nil {
			return err	
		}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body);if err != nil {
		return err
	}

	var newsApiResponse model.NewsAPIResponse
	
	err = json.Unmarshal(body, &newsApiResponse )

	if err != nil {
		return err
	}

	
	
	return c.JSON(http.StatusOK, newsApiResponse )
}

func Auth(c echo.Context)error{
	message := &model.MessageText{}

	if err := c.Bind(message); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, message)
}