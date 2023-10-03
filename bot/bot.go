package bot

import (
	"bot/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"net/url"
)

var count = 0
var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong 🏓!!")
	} else {
		var msg = sendMessage(m.Content)
		_, _ = s.ChannelMessageSend(m.ChannelID, msg)
	}
}
func sendMessage(message string) string {

	resp, err := http.PostForm("http://localhost:8080/process",
		url.Values{"message": {message}})
	if err != nil {
		fmt.Println(err)
		return "there is an err"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "there is an err"
	}

	return string(body)
}
