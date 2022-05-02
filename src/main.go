package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	dt := time.Now()

	tags := []string{}
	tags = GetWordsIO()

	wg.Add(1)
	go NumTwitterResp(tags...)
	fmt.Print("\n")
	fmt.Print("///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////")
	fmt.Print("\n")
	wg.Done()
	dtFim := time.Now()
	log.Printf("%s", dtFim.Sub(dt))
	wg.Wait()
	TenTwitterResp(tags...)
	wordCloud()

}

func NumTwitterResp(tags ...string) {
	fmt.Print("Número de Twitees para cada hashtag da frase digitada:" + "\n")
	for i := 0; i < len(tags); i++ {

		if int((" " + tags[i])[len(tags[i])]) == 10 {
			count(("#" + RemoveIndex([]rune(tags[i]))))

		} else {
			count(("#" + tags[i]))
		}
	}
}

func TenTwitterResp(tags ...string) {
	fmt.Print("Top 10 Twitees para cada hashtag da frase digitada:" + "\n")
	for i := 0; i < len(tags); i++ {

		if int((" " + tags[i])[len(tags[i])]) == 10 {
			topTen(("#" + RemoveIndex([]rune(tags[i]))))

		} else {
			topTen(("#" + tags[i]))
		}
	}

}
