package main

/*
** Homework-5/Task-3: Самостоятельно изучите пакет encoding/csv. Напишите пример, иллюстрирующий его применение.
** Khamenev Sergei
 */

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {

	// File's URL
	fileUrl := "https://my-files.ru/Save/pey181/Example.csv"

	// Download a file
	if err := DownloadFile("Example.csv", fileUrl); err != nil {
		panic(err)
	}

	// Open a file
	csvfile, err := os.Open("Example.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer csvfile.Close()

	// File's reading
	reader := csv.NewReader(csvfile)
	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rawCSVdata)
}
