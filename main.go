package main

import (
	"fmt"
	"github.com/vikbert/random/cmd"
)

func main() {
	err := cmd.ExecuteAll()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}

}
