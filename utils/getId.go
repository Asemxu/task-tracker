package utils

import "github.com/google/uuid"

func GetId() string {
	id := uuid.New()
	return id.String()
}
