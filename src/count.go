package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/g8rswimmer/go-twitter/v2"
)

// Função que recebe uma string e retorna o numero de tweets feitos com a hashtag
func count(text string) {
	count := []string{}
	token := "AAAAAAAAAAAAAAAAAAAAAHqTbwEAAAAAwm73WtWFdTK4m0wPh3nlaTMvBCI%3D7v2x6p9N7HWv7v5xhjxBlGepC16oF2xPiBrqHxCQR6OI9Vlotq"
	query := ""
	query = string(text)
	flag.Parse()

	client := &twitter.Client{
		Authorizer: authorize{
			Token: token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
	opts := twitter.TweetRecentCountsOpts{
		Granularity: twitter.GranularityHour,
	}

	tweetResponse, err := client.TweetRecentCounts(context.Background(), query, opts)
	if err != nil {
		log.Panicf("tweet recent counts error: %v", err)
	}
	metaBytes, err := json.MarshalIndent(tweetResponse.Meta, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Print("twitees: " + text + "\n")
	for i := 0; i < len(metaBytes); i++ {
		if string(metaBytes[i]) >= "0" && string(metaBytes[i]) <= "9" {
			count = append(count, string(metaBytes[i]))
		}
	}
	fmt.Print(strings.Join(count, "") + "\n")
}
