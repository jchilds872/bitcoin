package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dolthub/swiss"
)

func main() {
	dt := time.Now()
	fmt.Println("Script started at: ", dt.String())

	a := swiss.NewMap[string, struct{}](42)
	readFile, err := os.Open("./200k")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(-1)

	}
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
				m.Put(v[x+o], struct{}{})
			}
		}

		for x := int64(4); x < (4 + i); x++ {
			m.Delete(v[x])
		}
	}

	fmt.Println(m.Count())

	dt2 := time.Now()

	diff := dt2.Sub(dt)
	fmt.Println("Script ended at: ", dt.String())
	fmt.Printf("Seconds: %f\n", diff.Seconds())
}
