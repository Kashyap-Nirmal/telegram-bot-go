# Telegram Bot in Go

A simple Telegram bot built using Go and the `go-telegram-bot-api` package. This bot sends a welcome message to any user who messages it.

## Prerequisites
- [Golang](https://golang.org/doc/install) (latest version recommended)
- [Telegram bot token](https://core.telegram.org/bots#botfather) (get this from [@BotFather](https://t.me/BotFather) on Telegram)
- [Go Telegram Bot API](https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api/v5)
- Basic knowledge of environment variables

## Setup

### 1. Clone the Repository
```
git clone https://github.com/your-username/telegram-bot-go.git
cd telegram-bot-go
```

### 2. Install Dependencies

Install the necessary Go packages:

```
go get github.com/joho/godotenv
go get github.com/go-telegram-bot-api/telegram-bot-api/
```

### 3. Create a .env File

Create a .env file in the project root directory and add your Telegram bot token:

```
TELEGRAM_BOT_TOKEN=your-telegram-bot-token
```
## Start the bo

### 4. Run the Bot

Run the bot using the following command:

```bash
go run main.go
```

### 5. Interact with Your Bot

- Open Telegram and search for your bot by name.
- Send it a message, and it should reply with a test message.

## Project Structure

`main.go`: Main file containing bot setup and logic.

`.env`: Environment file containing sensitive information like bot token.

#### Debugging

- The bot logs all incoming messages and responses.

- Errors related to message sending or bot setup are also logged.

### Troubleshooting

- Bot token not found: Make sure your .env file is correctly configured and includes your Telegram bot token.

- Environment file not loading: Check that the godotenv package is installed and loaded correctly in main.go.

### Additional Information

- For detailed documentation on the `go-telegram-bot-api` package, visit the package documentation.

- For more information on creating and managing Telegram bots, check the `Telegram Bot API` documentation.