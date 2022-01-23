package runner

import (
	"flag"
	"log"
	"strings"
)

func Run() {
	filenamePtr := flag.String("filename", "", "Name of the file being seeked")
	seekLinesPtr := flag.Int("lines", 10, "Number of lines to seek, defaults to 10.")

	flag.Parse()

	filenameIsValid := validateFilename(*filenamePtr)
	if !filenameIsValid {
		log.Fatalln("Invalid filename (must be available CSV file):", *filenamePtr)
	}

	_ = seekLinesPtr
}

func validateFilename(filename string) bool {
	if !strings.HasSuffix(filename, ".csv") || len(filename) < 4 {
		return false
	}

	return true
}
