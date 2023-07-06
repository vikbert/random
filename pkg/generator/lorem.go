package generator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func RandomText(wordCount int) string {
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
