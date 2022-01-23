package runner

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func Run() {
	filenamePtr := flag.String("filename", "", "Name of the file being seeked")
	seekLinesPtr := flag.Int("lines", 10, "Number of lines to seek, defaults to 10.")

	flag.Parse()

	filenameIsValid := validateFilename(*filenamePtr)
	if !filenameIsValid {
		log.Fatal("Invalid filename (must be available CSV file):", *filenamePtr)
	}

	lines := seekCsvFile(*filenamePtr, *seekLinesPtr)

	fmt.Println(lines)
}

func validateFilename(filename string) bool {
	if !strings.HasSuffix(filename, ".csv") || len(filename) < 4 {
		return false
	}

	return true
}

func seekCsvFile(filename string, seekLines int) [][]string {
	csvLines := make([][]string, seekLines)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to read input file:", filename)
	}

	cr := csv.NewReader(f)

	for i := 0; i < seekLines; i++ {
		line, err := cr.Read()
		if err != nil {
			log.Fatalf("Failed to read line #%v.\nError Message: %v", i, err)
		}

		csvLines[i] = line
	}

	return csvLines
}
