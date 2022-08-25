package main

/// colly爬取表情包
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// 保存图片
func httpGet(imgUrl string, title string) {

	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	reader := bufio.NewReaderSize(res.Body, 32*1024)

	imgPath := "D:\\learn\\golang\\go_spider\\images\\" + title + "\\"

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

}
func main() {
	t := time.Now()
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"
	// colly.DetectCharset()检查编码模式

	/* DetectCharset 的作用为检查编码，设置这个为true后，抓取gbk编码的数据就可以解码了。
	 */
	c1 := colly.NewCollector(colly.UserAgent(userAgent), colly.DetectCharset())

	c2 := c1.Clone()

	//异步
	// c2.Async = true
	/* 	//限速
	   	c2.Limit(&colly.LimitRule{
	   		DomainRegexp: "",
	   		DomainGlob:   "*.m.qqtn.com/*",
	   		Delay:        10 * time.Second,
	   		RandomDelay:  0,
	   		Parallelism:  1,
	   	}) */

	//采集器1，获取列表目录
	c1.OnHTML("ul[class='tab-con-tx tab-con-bq clearfix g-dome-list']", func(h *colly.HTMLElement) {
		// 列表中的每一项
		h.ForEach("li", func(i int, item *colly.HTMLElement) {
			// 文章链接
			href := item.ChildAttr("a[href]", "href")
			title := item.ChildText("span")
			ctx := colly.NewContext()
			ctx.Put("href", href)
			ctx.Put("title", title)
			c2.Request("GET", "https://m.qqtn.com"+href, nil, ctx, nil)
			// fmt.Println(href, title)
		})
	})

	//采集器2，获取列表目录详情
	c2.OnHTML("article[class='g-cms-content']", func(e *colly.HTMLElement) {
		href := e.Request.Ctx.Get("href")
		title := e.Request.Ctx.Get("title")

		fmt.Println("----------" + title + "----------")
		// fmt.Println(href)

		//创建目录
		//perm权限设置，os.ModePerm为0777
		err := os.Mkdir("./images/"+title, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(href, "创建目录成功!")
		}

		// 查找遍历img图片
		e.DOM.Find("img").Each(func(i int, s *goquery.Selection) {

			imgUrl, _ := s.Attr("src")
			fmt.Printf("imgUrl: %v\n", imgUrl)
			httpGet(imgUrl, title)
		})

		fmt.Println()
	})

	c2.OnRequest(func(r *colly.Request) {
		fmt.Println("c2爬取页面:", r.URL)
	})

	c1.OnRequest(func(r *colly.Request) {
		fmt.Println("c1爬取页面:", r.URL)
	})

	c1.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c1.Visit("https://m.qqtn.com/bq/keaibq")
	if err != nil {
		fmt.Println(err.Error())
	}

	c2.Wait()
	fmt.Printf("爬取完成，花费时间:%s", time.Since(t))
}
