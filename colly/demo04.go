/* package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

//cgolang 判断文件或者文件夹是否存在可以通过os.stat() 方法和os.IsExist() 方法来判断：
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func main() {

	c := colly.NewCollector(colly.DetectCharset())

	c.OnHTML("#content", func(e *colly.HTMLElement) {
		// ret, _ := e.DOM.Html()
		// fmt.Printf("ret: %v\n", ret)
		title := e.DOM.Find("p").Text()
		fmt.Printf("title: %v\n", title)

		e.DOM.Find("img").Each(func(i int, s *goquery.Selection) {

			// if isExist(title) {
			// 	fmt.Printf("%s is exist!", title)

			// } else {

			// 	fmt.Printf("%s is not exist!", title)

			// }
			//创建目录
			//perm权限设置，os.ModePerm为0777
			// err3 := os.Mkdir("./"+title, os.ModePerm)
			// if err3 != nil {
			// 	log.Fatal(err3)
			// }

			imgUrl, _ := s.Attr("src")
			fmt.Printf("imgUrl: %v\n", imgUrl)
			res, err := http.Get(imgUrl)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			reader := bufio.NewReaderSize(res.Body, 32*1024)

			imgPath := "D:\\learn\\golang\\go_spider\\" + title + "\\"
			// 根据图片url获取其文件名
			fileName := path.Base(imgUrl)

			file, err2 := os.Create(imgPath + fileName)
			if err2 != nil {
				panic(err)
			}

			// 获得文件的writer对象
			writer := bufio.NewWriter(file)

			written, _ := io.Copy(writer, reader)
			// 输出文件字节大小
			fmt.Printf("Total lenth: %d\n", written)

		})
	})
	c.Visit("https://m.qqtn.com/c/324778")
}
*/