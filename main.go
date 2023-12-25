package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)



func newDocumentFromioReader(r io.Reader) (*Document, error) {
	root, err := r.Read()
	if err != nil {
		return nil, err
	}
	return root, nil
}
func AdventScraper() []string {
	headers := http.Header{}

	headers.Set("Cookie", "session=53616c7465645f5f19e8329ad55cd9cdaba8d6ebe1adaeb5f026767d559608583fb2109b79cf155dcb8801d14379ebac910419b4620f02356501273b2562edb5")

	req, err := http.NewRequest("GET", "https://adventofcode.com/2023/day/1/input", nil)
	if err != nil {
		fmt.Println("newRequest")
		log.Fatal(err)

	}

	req.Header = headers

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do")
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// recreating newDocumentFromReader from goQuery so i dont need the entire library

	doc, err := newDocumentFromioReader(resp.Body)
	if err != nil {
		fmt.Println("goquery.NewDocumentFromReader")
		log.Fatal(err)
	}
	out, err := os.Create("output")
	if err != nil {
		fmt.Println("os.Create")
		log.Fatal(err)
	}
	defer out.Close()

	copy, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("io.Copy")
		log.Fatal(err)
	}

	fmt.Println(copy)
	docString := doc.Text()

	defer resp.Body.Close()

	docStringArray := strings.Split(docString, "\n")

	for i := 0; i < len(docStringArray); i++ {
		docStringArray[i] = strings.TrimSpace(docStringArray[i])
	}

	return docStringArray
}

func AdventDayOne(docStringArray []string) {
	//remove non numeric characters from docStringArray convert the value to integer & place in new integer array
	var docIntArray []int
	for i := 0; i < len(docStringArray); i++ {
		docStringArray[i] = regexp.MustCompile(`[^0-9]+`).ReplaceAllString(docStringArray[i], "")
		stringConv, _ := strconv.Atoi(docStringArray[i])

		strParse, err := strconv.ParseInt(docStringArray[i], 10, 0)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(strParse)

		docIntArray = append(docIntArray, stringConv)

	}
	fmt.Println(docIntArray[2])

}

func main() {

	fmt.Println("Advent of Code Personal Scraper..... ENGAGE!")
	AdventScraper()
	//AdventDayOne(dayOneArray)
}
