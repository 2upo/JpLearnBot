package main

import (
	"fmt"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// katakana := []string{
	// 	"ア", "イ", "ウ", "エ", "オ",
	// 	"カ", "キ", "ク", "ケ", "コ",
	// 	"サ", "シ", "ス", "セ", "ソ",
	// 	"タ", "チ", "ツ", "テ", "ト",
	// 	"ナ", "ニ", "ヌ", "ネ", "ノ",
	// 	"ハ", "ヒ", "フ", "ヘ", "ホ",
	// 	"マ", "ミ", "ム", "メ", "モ",
	// 	"ヤ", "ユ", "ヨ",
	// 	"ラ", "リ", "ル", "レ", "ロ",
	// 	"ワ", "ヰ", "ヱ", "ヲ",
	// 	"ン",
	// 	"ガ", "ギ", "グ", "ゲ", "ゴ",
	// 	"ザ", "ジ", "ズ", "ゼ", "ゾ",
	// 	"ダ", "ヂ", "ヅ", "デ", "ド",
	// 	"バ", "ビ", "ブ", "ベ", "ボ",
	// 	"パ", "ピ", "プ", "ペ", "ポ",
	// 	"ァ", "ィ", "ゥ", "ェ", "ォ",
	// 	"ー",
	// }
	// hiragana := []string{
	// 	"あ", "い", "う", "え", "お",
	// 	"か", "き", "く", "け", "こ",
	// 	"さ", "し", "す", "せ", "そ",
	// 	"た", "ち", "つ", "て", "と",
	// 	"な", "に", "ぬ", "ね", "の",
	// 	"は", "ひ", "ふ", "へ", "ほ",
	// 	"ま", "み", "む", "め", "も",
	// 	"や", "ゆ", "よ",
	// 	"ら", "り", "る", "れ", "ろ",
	// 	"わ", "ゐ",
	// }
	hiraganaPic := make([]string, 46)
	for i := 0; i < 46; i++ {
		hiraganaPic[i] = fmt.Sprint(i+1) + ".jpg"
	}
	katakanaPic := make([]string, 51)
	for i := 0; i < 51; i++ {
		katakanaPic[i] = "katakana/" + fmt.Sprint(i+1) + ".jpg"
	}
	hiraganaTrans := []string{
		"a", "i", "u", "e", "o",
		"ka", "ki", "ku", "ke", "ko",
		"sa", "shi", "su", "se", "so",
		"ta", "chi", "tsu", "te", "to",
		"na", "ni", "nu", "ne", "no",
		"ha", "hi", "hu", "he", "ho",
		"ma", "mi", "mu", "me", "mo",
		"ya", "yu", "yo",
		"ra", "ri", "ru", "re", "ro",
		"wa", "wo", "n",
	}
	katakanaTrans := []string{
		"a", "i", "u", "e", "o",
		"ka", "ki", "ku", "ke", "ko",
		"sa", "shi", "su", "se", "so",
		"ta", "chi", "tsu", "te", "to",
		"na", "ni", "nu", "ne", "no",
		"ha", "hi", "hu", "he", "ho",
		"ma", "mi", "mu", "me", "mo",
		"ya", "yi", "yu", "ye", "yo",
		"ra", "ri", "ru", "re", "ro",
		"wa", "wi", "wu", "we", "wo", "n",
	}

	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	counter := 0
	randK := shuffleIndex(len(katakanaPic))

	for update := range updates {
		rand.Seed(time.Now().UnixNano())
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		randH := rand.Intn(len(hiraganaPic))
		if update.Message.Text == "/katakana" {
			if counter > len(katakanaPic)-1 {
				counter = 0
				randK = shuffleIndex(len(katakanaPic))
			}
			msg.Entities = []tgbotapi.MessageEntity{{Type: "spoiler", Offset: 0, Length: len(katakanaTrans[randK[counter]]) + 8}}
			msg.Text = "Answer: " + katakanaTrans[randK[counter]]
			file := tgbotapi.FilePath(katakanaPic[randK[counter]])
			msg2 := tgbotapi.NewPhoto(update.Message.Chat.ID, file)
			bot.Send(msg2)
			bot.Send(msg)
			counter++
		} else if update.Message.Text == "/hiragana" {
			msg.Entities = []tgbotapi.MessageEntity{{Type: "spoiler", Offset: 0, Length: len(hiraganaTrans[randH]) + 8}}
			msg.Text = "Answer: " + hiraganaTrans[randH]
			file := tgbotapi.FilePath("hiragana/" + hiraganaPic[randH])
			msg2 := tgbotapi.NewPhoto(update.Message.Chat.ID, file)
			bot.Send(msg2)
			bot.Send(msg)

		} else {
			msg.Text = "Chose /katakana or /hiragana"
			bot.Send(msg)
		}
	}
}

func shuffleIndex(len int) []int {
	result := make([]int, len)
	for i := 0; i < len; i++ {
		result[i] = i
	}
	//shuffle
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		shuffleIndex := rand.Intn(len)
		result[i], result[shuffleIndex] = result[shuffleIndex], result[i]
	}
	return result
}
