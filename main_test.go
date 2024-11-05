// main_test.go
package main

import (
	"os"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for the Telegram bot
type MockBotAPI struct {
	mock.Mock
}

func (m *MockBotAPI) Send(msg tgbotapi.MessageConfig) (tgbotapi.Message, error) {
	args := m.Called(msg)
	return args.Get(0).(tgbotapi.Message), args.Error(1)
}

func (m *MockBotAPI) GetUpdatesChan(config tgbotapi.UpdateConfig) (<-chan tgbotapi.Update, error) {
	args := m.Called(config)
	return args.Get(0).(chan tgbotapi.Update), args.Error(1)
}

func TestLoadEnv(t *testing.T) {
	// Ensure the .env file is loaded and the token is retrieved
	err := godotenv.Load()
	assert.NoError(t, err, "Error loading .env file")

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	assert.NotEmpty(t, botToken, "TELEGRAM_BOT_TOKEN should not be empty")
}

func TestBotInitialization(t *testing.T) {
	// Mock bot initialization
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	assert.NoError(t, err, "Failed to initialize bot")
	assert.NotNil(t, bot, "Bot should not be nil")
}

func TestSendMessage(t *testing.T) {
	// Set up mock bot
	mockBot := new(MockBotAPI)

	// Create a message to send
	msg := tgbotapi.NewMessage(12345, "Hello! This is a test message from my Golang bot.")

	// Prepare a response message
	expectedResponse := tgbotapi.Message{
		MessageID: msg.ReplyToMessageID,
		From:      &tgbotapi.User{ID: 67890, UserName: "bot_user"},
		Chat:      &tgbotapi.Chat{ID: msg.ChatID},
		Text:      msg.Text,
	}

	// Mock the Send method to return the expected response
	mockBot.On("Send", msg).Return(expectedResponse, nil)

	// Call Send method
	response, err := mockBot.Send(msg)
	assert.NoError(t, err, "Should not return an error when sending a message")
	assert.Equal(t, expectedResponse, response, "Response should match the expected response")
	mockBot.AssertExpectations(t)
}

func TestHandleUpdate(t *testing.T) {
	// Simulate an incoming update
	mockBot := new(MockBotAPI)
	updates := make(chan tgbotapi.Update)

	// Mock the GetUpdatesChan method
	mockBot.On("GetUpdatesChan", mock.Anything).Return(updates, nil)

	// Simulate a received message
	go func() {
		updates <- tgbotapi.Update{Message: &tgbotapi.Message{
			From: &tgbotapi.User{UserName: "test_user"},
			Text: "Hello!",
			Chat: &tgbotapi.Chat{ID: 12345},
		}}
	}()

	// Your logic to handle updates would go here, which should use mockBot.Send
	// Ensure that Send is called with the expected message
}

func TestMain(m *testing.M) {
	// Set up any necessary testing environment
	// ...

	// Run tests
	exitVal := m.Run()

	// Clean up
	// ...

	os.Exit(exitVal)
}
