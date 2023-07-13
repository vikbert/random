package main

import (
	"fmt"
	"github.com/vikbert/random/pkg/command"
)

func main() {
	err := command.ExecuteAll()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
