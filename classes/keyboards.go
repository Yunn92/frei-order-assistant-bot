package FreiOrderBot

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	var Keyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Создать заявку"),
			tgbotapi.NewKeyboardButton("Последняя заявка"),
		),
	)

	return &Keyboard
}

func EditKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	var Keyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Проверить заказ"),
		),
	)

	return &Keyboard
}

func InlineProducts(data *OrderTree) *tgbotapi.InlineKeyboardMarkup {
	var Keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Расходники", "consumables"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Крафт пакеты", "packages"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Алюм. контейнеры", "al_container"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Пласт. контейнеры", "p_container"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Кухня и кондитерка", "factory"),
		),
	)

	return &Keyboard
}

func InlineConsumables() *tgbotapi.InlineKeyboardMarkup {
	var Keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Химия", "chemical"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Барные", "bar"),
		),
	)

	return &Keyboard
}

func InlineBar() *tgbotapi.InlineKeyboardMarkup {
	var Keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Стаканы", "cups"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Другие", "others"),
		),
	)

	return &Keyboard
}

func ParseJSONToKeyboard(goods []Goods) *tgbotapi.InlineKeyboardMarkup {
	var Keyboard = tgbotapi.NewInlineKeyboardMarkup()

	for _, i := range goods {
		button_text := fmt.Sprintf("%s (%s)", i.InternalTitle, i.Count)
		temp := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(button_text, strconv.Itoa(i.Id)))

		Keyboard.InlineKeyboard = append(Keyboard.InlineKeyboard, temp)
	}

	return &Keyboard
}
