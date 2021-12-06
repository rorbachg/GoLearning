package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


func main() {
	dup4()

}

func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func printCounts(counts map[string]int){
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup1() {
	// First implementation
	counts := make(map[string]int)
	countLines(os.Stdin, counts)
	printCounts(counts)
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files{
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	printCounts(counts)
}

func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	printCounts(counts)
	}
}

func dup4() {
	counts := make(map[string]int)
	sources := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files{
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines2(f, counts, arg, sources)
			f.Close()
		}
	}
	printCounts2(counts, sources)
}

func countLines2(f *os.File, counts map[string]int, name string, sources map[string][]string){
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text]++
		src := sources[text]

		alreadyPresent := false
		for _, entry := range src {
			if entry == name {
				alreadyPresent = true
				break
			}
		}
		if !alreadyPresent {
			sources[text]  = append(sources[text], name)
		}
	}
}



func printCounts2(counts map[string]int, sources map[string][]string){
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, sources[line])
		}
	}
}