package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/g8rswimmer/go-twitter/v2"
	"golang.org/x/exp/slices"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

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

func RemoveIndex(s []rune) string {
	ret := []rune{}
	for i := 0; i < len(s)-1; i++ {
		ret = append(ret, s[i])
	}
	return string(ret)
}

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entre sua frase:")
	tags := []string{}
	city, _ := reader.ReadString('\n')
	phrase := strings.Split(city, " ")
	stopwords := []string{"de", "a", "o", "que", "e", "do", "da", "em", "um", "para", "é", "com", "não", "uma", "os", "no", "se", "na", "por", "mais", "as", "dos", "como", "mas", "foi", "ao", "ele", "das", "tem", "à", "seu", "sua", "ou", "ser", "quando", "muito", "há", "nos", "já", "está", "eu", "também", "só", "pelo", "pela", "até", "isso", "ela", "entre", "era", "depois", "sem", "mesmo", "aos", "ter", "seus", "quem", "nas", "me", "esse", "eles", "estão", "você", "tinha", "foram", "essa", "num", "nem", "suas", "meu", "às", "minha", "têm", "numa", "pelos", "elas", "havia", "seja", "qual", "será", "nós", "tenho", "lhe", "deles", "essas", "esses", "pelas", "este", "fosse", "dele", "tu", "te", "vocês", "vos", "lhes", "meus", "minhas", "teu", "tua", "teus", "tuas", "nosso", "nossa", "nossos", "nossas", "dela", "delas", "esta", "estes", "estas", "aquele", "aquela", "aqueles", "aquelas", "isto", "aquilo", "estou", "está", "estamos", "estão", "estive", "esteve", "estivemos", "estiveram", "estava", "estávamos", "estavam", "estivera", "estivéramos", "esteja", "estejamos", "estejam", "estivesse", "estivéssemos", "estivessem", "estiver", "estivermos", "estiverem", "hei", "há", "havemos", "hão", "houve", "houvemos", "houveram", "houvera", "houvéramos", "haja", "hajamos", "hajam", "houvesse", "houvéssemos", "houvessem", "houver", "houvermos", "houverem", "houverei", "houverá", "houveremos", "houverão", "houveria", "houveríamos", "houveriam", "sou", "somos", "são", "era", "éramos", "eram", "fui", "foi", "fomos", "foram", "fora", "fôramos", "seja", "sejamos", "sejam", "fosse", "fôssemos", "fossem", "for", "formos", "forem", "serei", "será", "seremos", "serão", "seria", "seríamos", "seriam", "tenho", "tem", "temos", "tém", "tinha", "tínhamos", "tinham", "tive", "teve", "tivemos", "tiveram", "tivera", "tivéramos", "tenha", "tenhamos", "tenham", "tivesse", "tivéssemos", "tivessem", "tiver", "tivermos", "tiverem", "terei", "terá", "teremos", "terão", "teria", "teríamos", "teriam"}
	for i := 0; i < len(phrase); i++ {
		if !slices.Contains(stopwords, phrase[i]) {
			tags = append(tags, phrase[i])
		}
	}
	fmt.Print("Número de Twitees para cada hashtag da frase digitada:" + "\n")
	for i := 0; i < len(tags); i++ {
		if int((" " + tags[i])[len(tags[i])]) == 10 {
			count(("#" + RemoveIndex([]rune(tags[i]))))
			topTen(("#" + RemoveIndex([]rune(tags[i]))))
		} else {
			count(("#" + tags[i]))
			topTen(("#" + tags[i]))
		}
	}
	fmt.Print("\n")
	fmt.Print("///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////")
	fmt.Print("\n")
	fmt.Print("Top 10 Twitees para cada hashtag da frase digitada:" + "\n")
	for i := 0; i < len(tags); i++ {
		if int((" " + tags[i])[len(tags[i])]) == 10 {
			topTen(("#" + RemoveIndex([]rune(tags[i]))))
		} else {
			topTen(("#" + tags[i]))
		}
	}

}
