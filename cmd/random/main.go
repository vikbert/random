package main

import (
	"fmt"
	"github.com/vikbert/random/pkg/command"
)

var (
	goVersion = "unknown"
	version   = "master"
	commit    = "?"
	date      = ""
)

func main() {
	err := command.ExecuteAll()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
