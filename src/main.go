package main

import (
	"fmt"
	"sync"
	"time"
	"log"
)
/// Tenho que passar o argumente TAG para as funções recebe-los e dar continuidade
/// dt e dtFim são paramentros para medir o tempo de execução

var wg sync.WaitGroup

func main() {

	dt := time.Now()

	tags := []string{}
	tags = GetWordsIO()

	//fmt.Print(tags[0])

	NumTwitterResp(tags...)

	dtFim := time.Now()
	log.Printf("%s",dtFim.Sub(dt))


	fmt.Print("\n")
	fmt.Print("///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////")
	fmt.Print("\n")

	TenTwitterResp(tags...)
}

func TenTwitterResp(tags ...string){
	fmt.Print("Top 10 Twitees para cada hashtag da frase digitada:" + "\n")
	for i := 0; i < len(tags); i++ {
		if int((" " + tags[i])[len(tags[i])]) == 10 {
			topTen(("#" + RemoveIndex([]rune(tags[i]))))
		} else {
			topTen(("#" + tags[i]))
		}
	}
	wordCloud();
}

func NumTwitterResp(tags ...string){
	fmt.Print("Número de Twitees para cada hashtag da frase digitada:" + "\n")
	for i := 0; i < len(tags); i++ {
		wg.Add(1)
		if int((" " + tags[i])[len(tags[i])]) == 10 {
			count(("#" + RemoveIndex([]rune(tags[i]))))
		} else {
			count(("#" + tags[i]))
		}
	}
}