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

	loremText := make([]string, wordCount)
	len := len(loremWords)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < wordCount; i++ {
		loremText[i] = loremWords[r.Intn(len)]
	}

	return fmt.Sprintf("%s.", strings.Join(loremText, " "))
}
