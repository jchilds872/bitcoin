package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	dt := time.Now()
	fmt.Println("Script started at: ", dt.String())

	var wg sync.WaitGroup
	readFile, err := os.Open("./200k")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(-1)
	}

	var insertions uint64
	var deletions uint64

	fileScanner := bufio.NewReader(readFile)

	for {
		line, err := fileScanner.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Failed to exit routine: ", err)
			os.Exit(-1)
		}

		wg.Add(1)
		go func(s string) {
			defer wg.Done()

			v := strings.Split(line, " ")
			i, err := strconv.ParseInt(v[2], 10, 64)
			if err != nil {
				fmt.Println("Could not parse int: ", err)
				os.Exit(-1)
			}
			o, err := strconv.ParseInt(v[3], 10, 64)
			if err != nil {
				fmt.Println("Could not parse int: ", err)
				os.Exit(-1)
			}

			for x := (4 + i); x < (4 + i + o); x++ {
				f, err := strconv.ParseFloat(v[x], 64)
				if err != nil {
					fmt.Println("Could not parse float: ", err)
					os.Exit(-1)
				}
				vv := int64(f * 100000000)

				if vv > 0 {
					// replacing adding to hashmap here, no need to keep track of the uuids, more performant to just keep track of no
					atomic.AddUint64(&insertions, 1)
				}
			}

			for x := int64(4); x < (4 + i); x++ {
				// replacing map deletion here
				atomic.AddUint64(&deletions, 1)
			}
		}(line)
	}

	// waiting for all threads to complete
	wg.Wait()

	fmt.Println("insertions: ", insertions)
	fmt.Println("deletions: ", deletions)
	fmt.Println("answer: ", insertions-deletions)

	dt2 := time.Now()

	diff := dt2.Sub(dt)
	fmt.Println("Script ended at: ", dt.String())
	fmt.Printf("Seconds: %f\n", diff.Seconds())
}
