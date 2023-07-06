package utils

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os/exec"
)

func ClipCopy(content string, prompt bool) {
	if clipboard.Unsupported == true {
		return
	}

	if err, _ := exec.Command("pbcopy < /dev/null").Output(); err != nil {
		fmt.Println("failed on clean up clip board by using 'pbcopy' ")
	}

	err := clipboard.WriteAll(content)
	if err != nil {
		fmt.Println("❗️Can't write to clipboard: ", err.Error())
	}

	if prompt {
		fmt.Println("✅ Generated content is copied to clipboard! \n", content)
	}
}
