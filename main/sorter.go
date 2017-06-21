package main

import (
	"algorithms/bubblesort"
	"algorithms/qsort"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var (
	infile    *string = flag.String("i", "infile", "File contains values for sorting")
	outfile   *string = flag.String("o", "outfile", "File to recive sorted values")
	algorithm *string = flag.String("a", "qsort", "Sort algorithm")
)

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPerfix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPerfix {
			fmt.Println("A too long time, seems unexpected.")
			return
		}

		str := string(line)

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int) {

}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile : ", *infile, "\toutfile : ", *outfile, "\talgorithm : ", *algorithm)
	}

	values, err := readValues(*infile)
	if err != nil {
		fmt.Println(err)
	} else {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm:", *algorithm, "is either unknown or unsupported")
		}
		t2 := time.Now
		fmt.Println("Sorting completed. costing ", t2.Sub(t1))

	}

}
