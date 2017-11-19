package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var Token = ""

func main() {
	fmt.Println(Token)

	//init new Discord session
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	// Register messageCreate as a handler for messages
	dg.AddHandler(messageCreate)

	// Open a websocket connectiond and start listening
	err = dg.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Crypto Discord Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

//handle messages
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	//split command in string array
	command := strings.Split(m.Content, " ")

	if command[0] == "!status" {
		s.ChannelMessageSend(m.ChannelID, "I'm ok. Thanks for asking")
	}

	if command[0] == "!help" {
		s.ChannelMessageSend(m.ChannelID, "Available commands:\n !status !help !price COIN !info COIN")
	}

	if command[0] == "!info" {
		if len(command) == 2 {
			msg, err := handleCoinInfo(command[1])
			if err != nil {
				log.Println(err)
			}
			s.ChannelMessageSend(m.ChannelID, msg)
		} else {
			s.ChannelMessageSend(m.ChannelID, "Wrong input")

		}

	}

	if command[0] == "!price" {
		if len(command) == 2 {
			msg, err := handleCoinPrice(command[1])
			if err != nil {
				log.Println(err)
			}
			s.ChannelMessageSend(m.ChannelID, msg)
		} else {
			s.ChannelMessageSend(m.ChannelID, "Wrong input")
		}

	}
}

//handleCoinInfo
func handleCoinInfo(coin string) (coinInfo string, err error) {
	coinInfo, err = GetCoinInfo(coin)
	if err != nil {
		log.Println("GetCoinInfo: ", err)
	}

	return
}

//handleCoinPrice
func handleCoinPrice(coin string) (coinPrice string, err error) {
	coinPrice, err = GetCoinPrice(coin)
	if err != nil {
		log.Println("GetCoinInfo: ", err)
	}

	return
}
