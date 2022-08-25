/* package main

/// goquery重构爬虫应用
import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://gorm.io/zh_CN/docs/"

	d, _ := goquery.NewDocument(url)
	d.Find(".sidebar-link").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		// fmt.Printf("href: %v\n", href)
		detail_url := url + href

		d, _ := goquery.NewDocument(detail_url)
		// title := d.Find(".article-title").Text()
		// fmt.Printf("title: %v\n", title)

		content, _ := d.Find(".article").Html()
		fmt.Printf("content: %v\n", content)
	})

}
*/