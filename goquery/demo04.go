/* package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getDoc1() {
	url := "https://gorm.io/zh_CN/docs/"
	d, _ := goquery.NewDocument(url)
	d.Find("")
}

func getDoc2() {
	client := &http.Client{}

	url := "https://gorm.io/zh_CN/docs/"

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	dom, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	dom.Find("")
}

func getDoc3() {

}
func main() {

}
*/