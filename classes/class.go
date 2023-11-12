package FreiOrderBot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Структура для товаров из json
type Goods struct {
	Id            int    `json:"id"`
	InternalTitle string `json:"internal"`
	PublicTitle   string `json:"public"`
	Count         string `json:"count"`
}

type Bar struct {
	Cups   []Goods `json:"cups"`
	Others []Goods `json:"others"`
}

type Consumables struct {
	Chemical []Goods `json:"chemical"`
	Bar      Bar     `json:"bar"`
}

type OrderTree struct {
	Cons        Consumables `json:"consumables"`
	Packages    []Goods     `json:"packages"`
	AlContainer []Goods     `json:"al_container"`
	PContainer  []Goods     `json:"p_container"`
	Factory     []Goods     `json:"factory"`
}

// Пресеты инлайновых клавиатур
type KeyboardList struct {
	StartKeyboard       *tgbotapi.ReplyKeyboardMarkup
	EditKeyboard        *tgbotapi.ReplyKeyboardMarkup
	MainInlineKeyboard  *tgbotapi.InlineKeyboardMarkup
	ConsumablesKeyboard *tgbotapi.InlineKeyboardMarkup
	ChemicalKeyboard    *tgbotapi.InlineKeyboardMarkup
	BarKeyboard         *tgbotapi.InlineKeyboardMarkup
	CupsKeyboard        *tgbotapi.InlineKeyboardMarkup
	OthersKeyboard      *tgbotapi.InlineKeyboardMarkup
	PackagesKeyboard    *tgbotapi.InlineKeyboardMarkup
	AlContainerKeyboard *tgbotapi.InlineKeyboardMarkup
	PContainerKeyboard  *tgbotapi.InlineKeyboardMarkup
	FactoryKeyboard     *tgbotapi.InlineKeyboardMarkup
}

// Единая структура бота
type Bot struct {
	BotAPI      *tgbotapi.BotAPI
	Goods       *OrderTree
	UsersBucket map[int64]string
	PlainGoods  map[int]Goods
	Keyboards   KeyboardList
}

// Функция на команду /start
func (b *Bot) StartFunc(id int64) {
	msg := tgbotapi.NewMessage(id, "Приветики, это бот для помощи в составлении заявки")
	msg.ReplyMarkup = StartKeyboard()

	b.BotAPI.Send(msg)
}

// Инициализация корзины пользователя
func (b *Bot) ProductList(id int64) {
	msg := tgbotapi.NewMessage(id, "Выбери нужную категорию")
	msg.ReplyMarkup = b.Keyboards.MainInlineKeyboard

	b.UsersBucket[id] = ""

	b.BotAPI.Send(msg)
}
