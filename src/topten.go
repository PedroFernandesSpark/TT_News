package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/g8rswimmer/go-twitter/v2"
)

func topTen(text string) {
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
	opts := twitter.TweetRecentSearchOpts{
		Expansions:  []twitter.Expansion{twitter.ExpansionEntitiesMentionsUserName, twitter.ExpansionAuthorID},
		TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt, twitter.TweetFieldConversationID, twitter.TweetFieldAttachments},
	}

	fmt.Println("Callout to tweet recent search callout")

	tweetResponse, err := client.TweetRecentSearch(context.Background(), query, opts)

	dictionaries := tweetResponse.Raw.TweetDictionaries()

	enc, err := json.MarshalIndent(dictionaries, "", "    ")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(enc))

	metaBytes, err := json.MarshalIndent(tweetResponse.Meta, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Print("twitees: " + text + "\n")
	fmt.Println(string(metaBytes))
	fmt.Print("\n")
}
