package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jbsmith7741/botter/apps/discord/bot"
	"github.com/mediaFORGE/ap-utils/config"
)

var (
	wg sync.WaitGroup
)

func main() {
	c := &bot.Config{}
	config.New(c).Parse()
	bot, err := bot.New(c)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	log.Printf("Success login: token:%v", bot.Token)

	wg.Add(1)

	go func() {
		for {
			time.Sleep(time.Second * 310)
			bot.Mine()
		}
	}()
	time.Sleep(5 * time.Second)
	wg.Add(1)
	go func() {
		for {
			time.Sleep(time.Hour)
			bot.Collect()
		}
	}()

	wg.Wait()
}
