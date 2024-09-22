package edc

import (
	"crypto/rand"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// RandomString generates a random string of a given length
func RandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789" +
		"!@#$%&*"

	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	var result strings.Builder
	for _, v := range b {
		result.WriteByte(charset[v%byte(len(charset))])
	}

	return result.String(), nil
}

// PaginateMetadata generates pagination metadata
func PaginateMetadata(currentPage, pageSize, totalItems int) fiber.Map {
	if pageSize <= 0 {
		pageSize = 1
	}

	lastPage := (totalItems + pageSize - 1) / pageSize
	beforePage := currentPage - 1
	nextPage := currentPage + 1

	metadata := fiber.Map{
		"current":  currentPage,
		"last":     lastPage,
		"total":    totalItems,
		"per_page": pageSize,
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
