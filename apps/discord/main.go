package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"os"

	"github.com/jbsmith7741/botter/apps/discord/bot"
	"github.com/jbsmith7741/botter/internal/version"
	"github.com/mediaFORGE/ap-utils/config"
	"github.com/davecgh/go-spew/spew"
)

var (
	wg sync.WaitGroup
)

func main() {
	c := &bot.Config{Email: "",
		Password: "",
		RoomId:   ""}

	version.ShowVersion()
	if err := config.New(c).Parse(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
