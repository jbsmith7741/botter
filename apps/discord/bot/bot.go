package bot

import (
	"fmt"
	"log"

	"time"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	Email    string
	Password string
	RoomId   string
}

type Bot struct {
	*Config
	client   *discordgo.Session
	UserName string
}

func New(c *Config) (b *Bot, err error) {
	b = &Bot{Config: c}
	b.client, err = discordgo.New(c.Email, c.Password)
	if b.client.Token == "" {
		return nil, fmt.Errorf("error: could not connect to client with %s", b.Email)
	}

	b.UserName, _ = b.client.User("@me")

	return b, err
}

func (b *Bot) Run() {
	mine := time.NewTicker(310 * time.Second)
	collect := time.NewTicker(time.Hour)
	for {
		select {
		case <-mine.C:
			b.Mine()
		case <-collect.C:
			b.Collect()
		}
	}
}

func (b *Bot) Mine() {
	if _, err := b.client.ChannelMessageSend(b.RoomId, "!mine"); err != nil {
		log.Printf("%v mine error", b.UserName)
	} else {
		log.Printf("%v mine\t", b.UserName)
	}
}

func (b *Bot) Collect() {
	if _, err := b.client.ChannelMessageSend(b.RoomId, "!collect"); err != nil {
		log.Printf("%v collect error", b.UserName)
	} else {
		log.Printf("%v collect\n", b.UserName)
	}
}

func (b *Bot) ChangeUserName(name string) {
	if _, err := b.client.UserUpdate(b.Email, b.Password, name, "", b.Password); err != nil {
		fmt.Println(err)
	}
}
