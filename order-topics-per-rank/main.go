package main

import (
	"fmt"
	"sort"
	"strconv"
)

// Author : Jerome AMON [Go Backend Developer]

// This is a nice coding challenge where you have to display a filtered list of books topics based on a given rank value.
// You are provided a list of books topics. And another list of rankings (where each rank value at index x corresponds to  the topic at index x).
// Below is an example : topics = []string{"maths", "science", "physics", "sports", "games"} .... rankings = []string{"4", "7", "5", "3", "4"}
// You need to return an ascending ordered list of topics which has equal or less rank than a given base rannking value.
// Example : baseRanking = 5  . . . => output = ["sports", "maths", "games", "physics"]
// Notes : ranking values are between 0 <= rank <= 10 and could be provided in word format : "one", "two", "three",  ... "eight".

func main() {

	topics := []string{"maths", "science", "physics", "sports", "games"}
	rankings := []string{"4", "7", "5", "three", "4"}
	baseRanking := 5

	fmt.Println(filerAndOrderBooksTopics(topics, rankings, baseRanking))
}

// complexity // time: nlog(n). space : O(n)
// where n is number of topics or rankings
func filerAndOrderBooksTopics(topics []string, rankings []string, baseRanking int) []string {
	type item struct {
		topic string
		rank  int
	}
	ranks := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "ten": 10}

	// empty list for wrong input.
	if len(topics) != len(rankings) {
		return []string{}
	}

	var items []item
	for index, topic := range topics {
		rank, err := strconv.Atoi(rankings[index])
		if err != nil {
			rank = ranks[rankings[index]]
			if rank == 0 || rank > baseRanking {
				// invalid or higher rank value.
				continue
			}
		}
		if rank > baseRanking {
			// valid but higher rank value.
			continue
		}
		item := item{topic: topic, rank: rank}
		items = append(items, item)
	}

	sort.Slice(items, func(i int, j int) bool {
		return items[i].rank < items[j].rank
	})

	var results []string
	for _, item := range items {
		results = append(results, item.topic)
	}

	return results
}

// program output
// go run .\filter-order-topics.go
// [sports maths games physics]
