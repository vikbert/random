package utils

import "github.com/google/uuid"

func GenerateUuid() string {
	return uuid.New().String()
}
