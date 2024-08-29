package errorLogger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	model "ai-project/models"
	"ai-project/utils/array"

	"github.com/google/uuid"
)

var ENV = os.Getenv("APP_ENV")

func CaptureException(trace string, err error) {

	// Get the file name of the current function
	pc, filePath, _, _ := runtime.Caller(1)
	file, _ := runtime.FuncForPC(pc).FileLine(pc)

	// Open the source file
	src, _ := os.ReadFile(file)

	// Split the file content into lines
	lines := strings.Split(string(src), "\n")

	var startingBlack int

	for line, lineContent := range lines {
		if strings.Contains(lineContent, trace) {
			startingBlack = getStartingBlack(1, line, lines)
			break
		}
	}

	// Find the markers and capture the content in between
	var capturedCode strings.Builder

	inBlock := false
	oldMarkerBlock := false
	count := 0

	var previousLine string

	for line, lineContent := range lines {

		if strings.Contains(lineContent, trace) {
			oldMarkerBlock = true
		}

		if line == startingBlack {
			inBlock = true
		}

		if inBlock {

			spaces := " "

			for i := 0; i < len(fmt.Sprintf("%v", len(lines)+2))-len(fmt.Sprintf("%v", line+1)); i++ {
				spaces = fmt.Sprintf("%v ", spaces)
			}

			capturedCode.Write([]byte(fmt.Sprintf("%v%v|%v\n", line+1, spaces, lineContent)))

			previousLine = lineContent
		}

		if inBlock && oldMarkerBlock && (count == 10 || strings.Contains(previousLine, "}")) {
			break
		}
	}

	id, e := uuid.NewV7()
	if e != nil {
		return
	}

	d := model.CaptureError{
		Id:        id.String(),
		Error:     err.Error(),
		CodeBlock: capturedCode.String(),
		FilePath:  filePath,
		Env:       os.Getenv("APP_ENV"),
		CreatedAt: time.Now().Format("2006-01-02T15:04:05-07:00"),
	}

	if ENV != "staging" && ENV != "open" && ENV != "qa" && ENV != "production" {
		fmt.Println(d.FilePath)
		fmt.Println(d.Error)
		fmt.Println(d.CodeBlock)
	} else {

		var errorCache []model.CaptureError

		errorCache = append(errorCache, d)

		jsonData, _ := json.MarshalIndent(errorCache, "", "  ") // Use MarshalIndent for pretty-printed JSON

		var mu sync.RWMutex

		mu.Lock()
		defer mu.Unlock()

		os.WriteFile("cache/files/error.json", jsonData, 0644)
	}
}

func DeleteException(id string) {
	errorCache := array.Filter(ErrorCache(), func(e model.CaptureError, _ int) bool {
		return e.Id != id
	})

	jsonData, _ := json.MarshalIndent(errorCache, "", "  ") // Use MarshalIndent for pretty-printed JSON

	var mu sync.RWMutex

	mu.Lock()
	defer mu.Unlock()

	os.WriteFile("cache/files/error.json", jsonData, 0644)
}

func getStartingBlack(count, line int, lines []string) int {
	if count == 10 || strings.Contains(lines[line], "{") {
		return line
	}

	count++
	return getStartingBlack(count, line-1, lines)
}
