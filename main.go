package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve the bot token from environment variables
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is required but not found in the environment")
	}

	// Initialize the bot with the token
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Failed to initialize bot: %v", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	log.Println("Bot started and listening for updates...")

	// Set up an update configuration to receive updates from Telegram
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	// Start receiving updates
	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Listen for incoming updates (messages)
	for update := range updates {
		if update.Message != nil { // Check if the update contains a message
			log.Printf("Received message from [%s]: %s", update.Message.From.UserName, update.Message.Text)

			// Respond with a simple message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! This is a test message from my Golang bot.")
			if _, err := bot.Send(msg); err != nil {
				log.Printf("Failed to send message: %v", err)
			}
		}
	}
}
