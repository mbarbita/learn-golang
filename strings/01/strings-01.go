package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func main() {

	var logins = make(map[string]string)

	lines, err := readLines("test.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for i, line := range lines {
		fmt.Println("Readed line:", i, line)
		if strings.HasPrefix(line, "#") {
			continue
		}
		fld := strings.Fields(line)
		fmt.Printf("Fields are: %q\n", fld)
		logins[fld[0]] = fld[1]

	}
	fmt.Println("map:", logins)

	// if err := writeLines(lines, "test.out.txt"); err != nil {
	// 	log.Fatalf("writeLines: %s", err)
	// }
}
