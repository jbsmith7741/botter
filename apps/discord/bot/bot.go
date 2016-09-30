package bot

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

type Config struct {
	Email    string
	Password string
	RoomId   string
}

type Bot struct {
	RoomID string
	Token string
	client *discordgo.Session
}

func New(c *Config) (b *Bot, err error) {
	b = &Bot{RoomID: c.RoomId}
	b.client, err = discordgo.New(c.Email, c.Password)
	b.Token = b.client.Token
	return b, err
}


func (b *Bot) Mine() {
	if _, err := b.client.ChannelMessageSend(b.RoomID, "!mine"); err != nil {
		log.Print("Mine Error")
	}
	log.Print("mine\t")

}
func (b *Bot) Collect() {
	if _, err := b.client.ChannelMessageSend(b.RoomID, "!collect"); err != nil {
		log.Print("Collect Error")
	}
	log.Print("collect\n")
}