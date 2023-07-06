package utils

import "github.com/google/uuid"

func RandomUuid() string {
	return uuid.New().String()
}
