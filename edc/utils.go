package edc

import (
	"crypto/rand"

	"github.com/gofiber/fiber/v2"
)

func RandomString(length int) (string, error) {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789" +
		"!@#$%&*"

	b := make([]byte, length)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	var result []byte
	for _, v := range b {
		result = append(result, charset[v%byte(len(charset))])
	}

	return string(result), nil
}

func PaginateMetadata(current int, limit int, total int) fiber.Map {
	lastPage := total/limit + 1
	beforePage := current - 1
	nextPage := current + 1

	metadata := fiber.Map{
		"current": current,
		"last":    lastPage,
		"total":   total,
	}

	if beforePage > 0 {
		metadata["previous"] = beforePage
	} else {
		metadata["previous"] = nil
	}

	if nextPage <= lastPage {
		metadata["next"] = nextPage
	} else {
		metadata["next"] = nil
	}

	return metadata
}
