package main

import (
	"fmt"
)

func main() {
	tags := []string{}
	tags = GetWordsIO()
	fmt.Print("Número de Twitees para cada hashtag da frase digitada:" + "\n")
	for i := 0; i < len(tags); i++ {
		if int((" " + tags[i])[len(tags[i])]) == 10 {
			count(("#" + RemoveIndex([]rune(tags[i]))))
		} else {
			count(("#" + tags[i]))
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
	wordCloud();
}
