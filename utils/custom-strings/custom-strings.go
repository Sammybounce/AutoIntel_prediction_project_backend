package customStrings

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomNumber(amountOfNumberToGenerate int) string {
	var numbers string

	for i := 0; i < amountOfNumberToGenerate; i++ {
		numbers = fmt.Sprintf("%v%v", numbers, rand.Intn(9))
	}

	return numbers
}

func GenerateRandomString(length int) string {
	rand.NewSource(time.Now().UnixNano())

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

func ReplaceQuotesInSql(s string) string {
	return strings.Replace(s, "'", "''", 1)
}
