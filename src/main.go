package main

import (
	"fmt"
	"sync"
	"time"
	"log"
)

var wg sync.WaitGroup

func main() {

	dt := time.Now()

	tags := []string{}
	tags = GetWordsIO()

	wg.Add(1)
	NumTwitterResp(tags...)


	fmt.Print("\n")
	fmt.Print("///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////")
	fmt.Print("\n")

	wg.Add(1)
	TenTwitterResp(tags...)

	dtFim := time.Now()
	log.Printf("%s",dtFim.Sub(dt))

	wg.Wait()

}

func NumTwitterResp(tags ...string){
	fmt.Print("Número de Twitees para cada hashtag da frase digitada:" + "\n")
	for i := 0; i < len(tags); i++ {

		if int((" " + tags[i])[len(tags[i])]) == 10 {
			go count(("#" + RemoveIndex([]rune(tags[i]))))

		} else {
			go count(("#" + tags[i]))	
		}
	}
}

func TenTwitterResp(tags ...string){
	fmt.Print("Top 10 Twitees para cada hashtag da frase digitada:" + "\n")
	for i := 0; i < len(tags); i++ {

		if int((" " + tags[i])[len(tags[i])]) == 10 {
			go topTen(("#" + RemoveIndex([]rune(tags[i]))))

		} else {
			go topTen(("#" + tags[i]))
		}
	}
	wordCloud();
}

