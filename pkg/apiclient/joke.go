package apiclient

import (
	"fmt"
	"github.com/icelain/jokeapi"
)

func FetchJoke() string {
	api := jokeapi.New()
	response, err := api.Fetch()
	if err != nil {
		fmt.Println("Can't fetch data from remote joke API")
	}

	return response.Joke[0]
}
