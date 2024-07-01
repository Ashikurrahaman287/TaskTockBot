package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	todoList []string
	chatID   int64
	botToken = "7410641542:AAFjqnwXQ5334lv_aw8Oi8rOhLOg-OeomZg"
)

func main() {
	// HTTP client with increased timeout and retry logic
	httpClient := &http.Client{
		Timeout: 10 * time.Second, // Increased timeout to 10 seconds
		Transport: &http.Transport{
			MaxIdleConns:          10,
			IdleConnTimeout:       30 * time.Second,
			DisableCompression:    true,
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 10 * time.Second, // Increased timeout for response headers
		},
	}

	bot, err := tgbotapi.NewBotAPIWithClient(botToken, tgbotapi.APIEndpoint, httpClient)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			switch {
			case strings.HasPrefix(update.Message.Text, "/start"):
				startCommand(bot, update.Message)
			case strings.HasPrefix(update.Message.Text, "/add"):
				addCommand(bot, update.Message)
			case strings.HasPrefix(update.Message.Text, "/done"):
				doneCommand(bot, update.Message)
			}
		}
	}
}

func startCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	chatID = message.Chat.ID
	welcomeMsg := tgbotapi.NewMessage(chatID, "Welcome to TaskTockBot! Use /add <task> to add tasks and /done <task> to mark them as done.")
	bot.Send(welcomeMsg)
	go startReminders(bot, chatID)
}

func addCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	task := strings.TrimPrefix(message.Text, "/add ")
	if task != "" {
		todoList = append(todoList, task)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Task added: "+task)
		bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Usage: /add <task>")
		bot.Send(msg)
	}
}

func doneCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	task := strings.TrimPrefix(message.Text, "/done ")
	taskFound := false
	for i, t := range todoList {
		if t == task {
			todoList = append(todoList[:i], todoList[i+1:]...)
			msg := tgbotapi.NewMessage(message.Chat.ID, "Task marked as done: "+task)
			bot.Send(msg)
			taskFound = true
			break
		}
	}
	if !taskFound {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Task not found: "+task)
		bot.Send(msg)
	}
}

func startReminders(bot *tgbotapi.BotAPI, chatID int64) {
	for {
		time.Sleep(1 * time.Hour)
		for _, task := range todoList {
			msg := tgbotapi.NewMessage(chatID, "Reminder: "+task)
			bot.Send(msg)
		}
	}
}
