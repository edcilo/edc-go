package edc

import "crypto/rand"

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
