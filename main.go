package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN") // Use the correct environment variable name
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panicf("Failed to create bot: %v", err)
	}

	bot.Debug = true
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panicf("Failed to get updates channel: %v", err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		handleUpdate(bot, update)
	}
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		startCommand(bot, update.Message)
	case "help":
		helpCommand(bot, update.Message)
	case "lyrics": // Handle the /lyrics command
		lyricsCommand(bot, update.Message)
	default:
		defaultMessage(bot, update.Message)
	}
}

func lyricsCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	log.Printf("Received message text: %s", message.Text)
	// Extract the song title from the message text
	songTitle := message.CommandArguments()
	log.Printf("Received song title: %s", songTitle) // Debugging output
	if songTitle == "" {
		log.Println("No song title provided.")
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Please provide a song title after the /lyrics command."))
		return
	}

	lyrics := getLyrics(songTitle)
	msg := tgbotapi.NewMessage(message.Chat.ID, lyrics)
	bot.Send(msg)
}

func startCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome! I am your friendly bot. How can I assist you today?")
	bot.Send(msg)
}

func helpCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Here are some commands you can use:\n/start - Start the bot\n/help - Get help information")
	bot.Send(msg)
}

func defaultMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	responseText := fmt.Sprintf("You said: %s", message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, responseText)
	bot.Send(msg)
}

// Placeholder function to simulate lyrics retrieval

// Placeholder function to simulate lyrics retrieval
func getLyrics(songTitle string) string {
	// Convert the input to lowercase for case-insensitive comparison
	songTitle = strings.ToLower(songTitle)
	if songTitle == "serah yamesegenehal" {
		return `ዘማሪ ተስፋዬ ጋቢሶ ርዕስ ሥራህ ያመሰግንሃል
ሥራህ ያመሰግንሃል ቅዱሳን ያከብሩሃል ጌትነትህም ገብቷቸው ተንበርክኳል ጉልበታቸው
አዝ ሃሌሉያ ተመስገን ጌታ ተመስገን ተመስገን (፫x) ኢየሱስ ተመስገን
እውነትህን ጠብቀሃል ለተበደሉት ፈርደሃል የተጠቃውን ታድገሃል በማዳንህ ከፍ ብለሃል
አዝ ሃሌሉያ ተመስገን ጌታ ተመስገን ተመስገን (፫x) ኢየሱስ ተመስገን
ደሃ አደጉን ተቀብለህ ለተራበ ምግብን ሰጥተህ ምጻተኛ አስጠግተሃል ግዞተኛውን ፈተሃል
አዝ ሃሌሉያ ተመስገን ጌታ ተመስገን ተመስገን (፫x) ኢየሱስ ተመስገን
የተቀደሰ ሥምህን በዜማ ቅኔ ላመስግን በአንተ ሆኖልኝ የመዳን ቀን ስለወጣሁ ከሰቀቀን
አዝ ሃሌሉያ ተመስገን ጌታ ተመስገን ተመስገን (፫x) ኢየሱስ ተመስገን`
	}
	return "Sorry, I couldn't find the lyrics for that song."
}
