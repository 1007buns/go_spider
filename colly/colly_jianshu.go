/* package main

/// 爬取简书首页文章列表
import (
	"fmt"



	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func main() {
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1), colly.Debugger(&debug.LogDebugger{}))
	//文章列表
	// 文章列表为ul标签，中间每一项是li标签，li中包含content，content中包含title，abstract和meta标签
	c.OnHTML("ul[class='note-list']", func(e *colly.HTMLElement) {
		//列表中每一项
		e.ForEach("li", func(i int, item *colly.HTMLElement) {
			//文章链接
			href := item.ChildAttr("div[class='content'] > a[class='title']", "href")
			//文章标题
			title := item.ChildText("div[class='content'] > a[class='title']")
			//文章摘要
			summary := item.ChildText("div[class='content'] > p[class='abstract']")
			fmt.Println(title, href)
			fmt.Println(summary)
			fmt.Println()
		})
	})

	err := c.Visit("https://www.jianshu.com")
	if err != nil {
		fmt.Println(err.Error())
	}
}
*/

/* package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

func main() {
	c1 := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1))
	c2 := c1.Clone()

	//异步
	c2.Async = true
	//限速
	c2.Limit(&colly.LimitRule{
		DomainRegexp: "",
		DomainGlob:   "*.jianshu.com/p/*",
		Delay:        10 * time.Second,
		RandomDelay:  0,
		Parallelism:  1,
	})
	//采集器1，获取文章列表
	c1.OnHTML("ul[class='note-list']", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, item *colly.HTMLElement) {
			href := item.ChildAttr("div[class='content'] > a[class='title']", "href")
			title := item.ChildText("div[class='content'] > a[class='title']")
			summary := item.ChildText("div[class='content'] > p[class='abstract']")
			ctx := colly.NewContext()
			ctx.Put("href", href)
			ctx.Put("title", title)
			ctx.Put("summary", summary)
			//通过Context上下文对象将采集器1采集到的数据传递到采集器2
			c2.Request("GET", "https://www.jianshu.com"+href, nil, ctx, nil)
		})
	})
	//采集器2，获取文章详情
	c2.OnHTML("article", func(e *colly.HTMLElement) {
		href := e.Request.Ctx.Get("href")
		title := e.Request.Ctx.Get("title")
		summary := e.Request.Ctx.Get("summary")
		detail := e.Text

		fmt.Println("----------" + title + "----------")
		fmt.Println(href)
		fmt.Println(summary)
		fmt.Println(detail)
		fmt.Println()
	})

	c2.OnRequest(func(r *colly.Request) {
		fmt.Println("c2爬取页面：", r.URL)
	})

	c1.OnRequest(func(r *colly.Request) {
		fmt.Println("c1爬取页面：", r.URL)
	})

	c1.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c1.Visit("https://www.jianshu.com")
	if err != nil {
		fmt.Println(err.Error())
	}

	c2.Wait()
}
*/

/*
package main

import (
    "fmt"
    "github.com/gocolly/colly"
    "github.com/gocolly/colly/debug"
    "github.com/gocolly/colly/extensions"
    _ "github.com/gocolly/colly/extensions"
    "net/http"
)

func main() {
    url := "https://www.jianshu.com"

    c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1), colly.Debugger(&debug.LogDebugger{}))
    c.OnHTML("ul[class='note-list']", func(e *colly.HTMLElement) {
        e.ForEach("li", func(i int, item *colly.HTMLElement) {
            href := item.ChildAttr("div[class='content'] > a[class='title']", "href")
            title := item.ChildText("div[class='content'] > a[class='title']")
            summary := item.ChildText("div[class='content'] > p[class='abstract']")
            fmt.Println(title, href)
            fmt.Println(summary)
            fmt.Println()
        })
    })

    //设置随机useragent
    extensions.RandomUserAgent(c)
    //设置登录cookie
    c.SetCookies(url, []*http.Cookie{
        &http.Cookie{
            Name:     "remember_user_token",
            Value:    "wNDUxOV0sIiQyYSQxMSRwdkhqWVhHYmxXaDJ6dEU3NzJwbmsuIiwiMTU",
            Path:     "/",
            Domain:   ".jianshu.com",
            Secure:   true,
            HttpOnly: true,
        },
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("爬取页面：", r.URL)
    })

    c.OnError(func(r *colly.Response, err error) {
        fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
    })
    err := c.Visit(url)
    if err != nil {
        fmt.Println(err.Error())
    }
}
*/