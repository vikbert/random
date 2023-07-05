package utils

import (
	"fmt"
	"github.com/atotto/clipboard"
)

func ClipCopy(content string, prompt bool) {
	if clipboard.Unsupported == true {
		return
	}

	err := clipboard.WriteAll(content)
	if err != nil {
		fmt.Println("❗️Can't write to clipboard: ", err.Error())
	}

	if prompt {
		fmt.Println("✅ Generated content is copied to clipboard: ", content)
	}
}
