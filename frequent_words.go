package main

import (
	"fmt"
	"sort"
	"strings"
)

type WordOps struct {
	Word  string
	Count int
}

func GetWordFrequency() {
	text := "India is my country and all Indians are my brothers and sisters I love my country and I am proud of its rich and varied heritage. I shall always strive to be worthy of it I shall give respect to my parents teachers and elders and treat everyone with courtesy"
	wordCount := CountFrequency(text)
	fmt.Println(wordCount)
}

func CountFrequency(text string) (result []WordOps) {
	wordCounts := make(map[string]int)
	strArray := strings.Fields(text)

	// mapping of words and their no of occurence in given string
	for _, word := range strArray {
		_, ok := wordCounts[word]
		if ok {
			wordCounts[word] += 1
		} else {
			wordCounts[word] = 1
		}
	}

	// preparing string array to sort by word's count
	words := make([]string, 0, len(wordCounts))
	for word := range wordCounts {
		words = append(words, word)
	}

	sort.SliceStable(words, func(i, j int) bool {
		return wordCounts[words[i]] > wordCounts[words[j]]
	})

	//return top 10 frequent words
	for i := 0; i < 10; i++ {
		if i < len(words) {
			result = append(result, WordOps{Word: words[i], Count: wordCounts[words[i]]})
		}
	}

	return
}

func main() {
	GetWordFrequency()
}
