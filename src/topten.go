package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/tidwall/gjson"
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

	var j = []byte(string(enc))
	c := make(map[string]json.RawMessage) // a map container to decode the JSON structure into
	e := json.Unmarshal(j, &c)            // unmarschal JSON

	if e != nil { // panic on error
		panic(e)
	}

	k := make([]string, len(c)) // a string slice to hold the keys

	i := 0                // iteration counter
	for s, _ := range c { // copy c's keys into k
		k[i] = s
		i++
	}

	fmt.Print("twitees: " + text + "\n")

	ttTextList := []string{}
	for i := 0; i < len(k); i++ {

		ttText := gjson.Get(string(enc), string(k[i]+".Tweet.text"))
		ttName := gjson.Get(string(enc), string(k[i]+".Author.name"))
		ttDate := gjson.Get(string(enc), string(k[i]+".Tweet.created_at"))

		println("Nome do usuario: " + ttName.String())
		println("Texto: " + ttText.String())
		println("Data: " + ttDate.String() + "\n")

		println("- - - - - - - - - - - - - - - - - - - -\n")

		auxList := []string{}
		auxList = GetWords(ttText.Str)

		for j := 0; j < len(auxList); j++ {
			ttTextList = append(ttTextList, auxList[j])

		}

	}

	wordcloudText := make(map[string]int)
	for i := 0; i < len(ttTextList); i++ {

		wordcloudText[ttTextList[i]] = 1

		for j := 0; j < len(ttTextList); j++ {

			if ttTextList[i] == ttTextList[j] {
				wordcloudText[ttTextList[i]] = wordcloudText[ttTextList[i]] + 1
			}

		}
	}

	jsonString, err := json.Marshal(wordcloudText)

	if err := os.Truncate("C:/Users/Spalko/Documents/TT_News/src/wordcloud/input.json", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}

	file, err := os.OpenFile("C:/Users/Spalko/Documents/TT_News/src/wordcloud/input.json", os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		panic(err)
	}

	_, err2 := file.Write(jsonString)

	if err2 != nil {
		log.Fatal(err2)
	}

	// metaBytes, err := json.MarshalIndent(tweetResponse.Meta, "", "    ")
	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println(string(metaBytes))
	// fmt.Print("\n")
}
