package main

import (
	"fmt"
	"sync"

	"os"

	"github.com/jbsmith7741/botter/apps/discord/bot"
	"github.com/jbsmith7741/botter/internal/version"
	"github.com/mediaFORGE/ap-utils/config"
)

var (
	wg sync.WaitGroup
)

func main() {
	c := &bot.Config{}

	version.ShowVersion()
	if err := config.New(c).Parse(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bot, err := bot.New(c)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	} else {
		fmt.Printf("%v has connected successfully\n", bot.UserName)
	}

	wg.Add(1)

	go bot.Run()

	wg.Wait()
}
