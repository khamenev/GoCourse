package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"./readerwriter"
)

func main() {

	lines, err := readerwriter.ReadLines("urls.txt")
	if err != nil {
		fmt.Println("Error: %sn", err)
		return
	}
	for _, line := range lines {
		fmt.Println(line)
	}
	err = readerwriter.WriteLines(lines, "urlsErr.txt")
	fmt.Println(err)

	start := time.Now()
	ch := make(chan string)
	for _, url := range lines {
		go fetch(url, ch)
	}
	for range lines {
		fmt.Print(<-ch) // receive from channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}
