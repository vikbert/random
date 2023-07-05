package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const defaultCount = 120

func GenerateText(wordCount int) string {
	if 0 == wordCount {
		wordCount = defaultCount
	}

	loremWords := []string{
		"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur",
		"adipiscing", "elit", "sed", "do", "eiusmod", "tempor",
		"incididunt", "ut", "labore", "et", "dolore", "magna",
		"aliqua.", "Ut", "enim", "ad", "minim", "veniam",
	}

	rand.Seed(time.Now().UnixNano())
	len := len(loremWords)

	loremText := make([]string, wordCount)
	for i := 0; i < wordCount; i++ {
		loremText[i] = loremWords[rand.Intn(len)]
	}

	return fmt.Sprintf("%s.", strings.Join(loremText, " "))
}
