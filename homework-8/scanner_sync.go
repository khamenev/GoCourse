package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
	for _, url := range lines {
		fetchSync(url)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchSync(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//bytes, err := ioutil.ReadAll(resp.Body)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}
