package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	FreiOrderBot "github.com/yunn92/frei-order-bot/classes"
)

var bot FreiOrderBot.Bot

// Инициализация в глобальную структуру бота
func init() {
	bot.BotAPI = bot_init()
	bot.Goods = data_init()
	bot.UsersBucket = make(map[int64]FreiOrderBot.Bucket)
	bot.PlainGoods = *FreiOrderBot.TreeToPlain(*bot.Goods)
	keyboard_init(&bot)
}

func main() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.BotAPI.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {

			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					bot.StartFunc(update.Message.Chat.ID) //class.go
				case "cancel":
					log.Println("This function will avaliable soon")
				}
			}

			//Придумать как обработать данные только после ввода количества товара
			if update.Message.Text == "Создать заявку" {
				bot.ProductList(update.Message.Chat.ID) //class.go --> Инициализация корзины пользователя
			} else if count, err := strconv.Atoi(update.Message.Text); err == nil { //Проверка что в качестве количества введено число
				msg := tgbotapi.NewMessage(update.Message.From.ID, "Добавлено успешно")

				temp := bot.UsersBucket[update.Message.From.ID].Text
				temp += fmt.Sprintf(" %d", count)

				if _, err := bot.BotAPI.Send(msg); err != nil {
					log.Println(err)
				}

				msg = tgbotapi.NewMessage(update.Message.From.ID, "Выбери следующий товар")
				msg.ReplyMarkup = bot.Keyboards.MainInlineKeyboard

				if _, err := bot.BotAPI.Send(msg); err != nil {
					log.Println(err)
				}
			} else if _, err := strconv.Atoi(update.Message.Text); err != nil {
				msg := tgbotapi.NewMessage(update.Message.From.ID, "Введено не число")

				if _, err := bot.BotAPI.Send(msg); err != nil {
					log.Println(err)
				}

				msg = tgbotapi.NewMessage(update.Message.From.ID, "Выбери следующий товар")
				msg.ReplyMarkup = bot.Keyboards.MainInlineKeyboard

				if _, err := bot.BotAPI.Send(msg); err != nil {
					log.Println(err)
				}
			}

			log.Printf("%s[%v]:%s", update.Message.From.FirstName, update.Message.Chat.ID, update.Message.Text)
		} else if update.CallbackQuery != nil {
			log.Printf("%s[%v]:%s", update.CallbackQuery.From.FirstName, update.CallbackQuery.From.ID, update.CallbackQuery.Data)

			var msg tgbotapi.EditMessageReplyMarkupConfig

			data := update.CallbackQuery.Data
			user_id := update.CallbackQuery.From.ID

			switch data {
			case "consumables":
				msg = tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, *bot.Keyboards.ConsumablesKeyboard)
			case "bar":
				msg = tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, *bot.Keyboards.BarKeyboard)
			case "chemical":
				msg = tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, *bot.Keyboards.ChemicalKeyboard)
			case "cups":
				msg = tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, *bot.Keyboards.CupsKeyboard)
			case "others":
				msg = tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, *bot.Keyboards.OthersKeyboard)
			case "packages":
				msg = tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, *bot.Keyboards.PackagesKeyboard)
			case "al_container":
				msg = tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, *bot.Keyboards.AlContainerKeyboard)
			case "p_container":
				msg = tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, *bot.Keyboards.PContainerKeyboard)
			case "factory":
				msg = tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, *bot.Keyboards.FactoryKeyboard)
			default:
				msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "Сколько?")

				if _, err := bot.BotAPI.Send(msg); err != nil {
					log.Println(err)
				}
			}

			if goodsId, err := strconv.Atoi(data); err != nil {
				temp := bot.UsersBucket[user_id].Text

				temp += bot.PlainGoods[goodsId].PublicTitle
			}

			if _, err := bot.BotAPI.Send(msg); err != nil {
				log.Println(err)
			}
		}
		continue
	}
}

func data_init() *FreiOrderBot.OrderTree {
	var goods FreiOrderBot.OrderTree

	file, err := os.Open("./temp/goods.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	jsonErr := json.Unmarshal(data, &goods)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return &goods
}

func bot_init() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI("6542249200:AAEOg_Eo9KEC5BHThnTIWvmo9BvMGAAGUoI")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

func keyboard_init(bot *FreiOrderBot.Bot) {
	bot.Keyboards.MainInlineKeyboard = FreiOrderBot.InlineProducts(bot.Goods)
	bot.Keyboards.EditKeyboard = FreiOrderBot.EditKeyboard()
	bot.Keyboards.ConsumablesKeyboard = FreiOrderBot.InlineConsumables()
	bot.Keyboards.BarKeyboard = FreiOrderBot.InlineBar()
	bot.Keyboards.ChemicalKeyboard = FreiOrderBot.ParseJSONToKeyboard(bot.Goods.Cons.Chemical)
	bot.Keyboards.CupsKeyboard = FreiOrderBot.ParseJSONToKeyboard(bot.Goods.Cons.Bar.Cups)
	bot.Keyboards.OthersKeyboard = FreiOrderBot.ParseJSONToKeyboard(bot.Goods.Cons.Bar.Others)
	bot.Keyboards.PackagesKeyboard = FreiOrderBot.ParseJSONToKeyboard(bot.Goods.Packages)
	bot.Keyboards.AlContainerKeyboard = FreiOrderBot.ParseJSONToKeyboard(bot.Goods.AlContainer)
	bot.Keyboards.PContainerKeyboard = FreiOrderBot.ParseJSONToKeyboard(bot.Goods.PContainer)
	bot.Keyboards.FactoryKeyboard = FreiOrderBot.ParseJSONToKeyboard(bot.Goods.Factory)
}
