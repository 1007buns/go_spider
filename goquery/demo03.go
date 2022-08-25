/* package main

/// 爬虫库 goquery
import (
	"fmt"









	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://gorm.io/zh_CN/docs/"
	d, _ := goquery.NewDocument(url)

	d.Find(".sidebar-link").Each(func(i int, s *goquery.Selection) {
		// s2 := s.Text()
		// fmt.Printf("s2: %v\n", s2)
		href, _ := s.Attr("href")
		fmt.Printf("href: %v\n", href)
	})
} */

