package main

import (
	"fmt"
)
/// Tenho que passar o argumente TAG para as funções recebe-los e dar continuidade


func main() {
	tags := []string{}
	tags = GetWordsIO()

	go TenTwitterResp(GetWordsIO())

	fmt.Print("\n")
	fmt.Print("///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////")
	fmt.Print("\n")

	go NumTwitterResp(GetWordsIO())
}

func TenTwitterResp(tags){
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

func NumTwitterResp(tags){
	fmt.Print("Número de Twitees para cada hashtag da frase digitada:" + "\n")
	for i := 0; i < len(tags); i++ {
		if int((" " + tags[i])[len(tags[i])]) == 10 {
			count(("#" + RemoveIndex([]rune(tags[i]))))
		} else {
			count(("#" + tags[i]))
		}
	}
}