package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readCfgFile(fileName string) map[string]string {
	fmt.Println("loading config file:", fileName, "...")

	cfgMap := make(map[string]string)

	// Open cfg file
	cfgFile, err := os.Open(fileName)
	if err != nil {
		log.Println("cfg file error:")
		log.Fatal(err)
	}
	defer cfgFile.Close()

	// prepare scanner, scan each line
	scanner := bufio.NewScanner(cfgFile)
	index := 1       // index
	cfgErrors := 0   // errors
	cfgWarnings := 0 // warnings

	for scanner.Scan() {

		// skip comment lines
		if strings.HasPrefix(scanner.Text(), "#") {
			index++
			continue
		}

		// split at firtst =
		sc := scanner.Text()
		eqi := strings.Index(sc, "=")
		eqc := strings.Count(sc, "=")
		// fmt.Println("= index:", sc, eqi)

		// skip empty lines
		if len(strings.TrimSpace(sc)) == 0 {
			fmt.Println("SKIPPING EMPTY LINE: line", index)
			index++
			continue
		}

		if eqc > 1 {
			fmt.Println("WARNING:", eqc, "\"=\" at line", index)
			cfgWarnings++
		}

		if eqi == -1 {
			fmt.Println("ERROR: no \"=\" at line", index)
			cfgErrors++
			index++
			continue
		}

		cfgMap[strings.TrimSpace(strings.ToLower(sc[:eqi]))] =
			strings.TrimSpace(sc[eqi+1:])
		index++

	} //end for

	fmt.Println("loading config file", fileName, "done",
		"\nwarnings:", cfgWarnings, "\nerrors:", cfgErrors)

	return cfgMap
}

func main() {

	cfg := readCfgFile("cfg.ini")
	// fmt.Println(cfg)
	fmt.Println()

	for k, v := range cfg {
		fmt.Printf("key:[%v] value:[%v]\n", k, v)
	}

}
