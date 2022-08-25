/* package main

/// goquery Api 选择器
import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func f1() {
	html := `<body>
	<div>div1</div>
	<div>div2</div>
	<span>Span</span>
	</body>
	`

	d, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	d.Find("span").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("s.Text(): %v\n", s.Text())
	})
}

func f2() {
	html := `<body>
	<div id="div1">div1</div>
	<div>div2</div>
	<span>Span</span>
	</body>
	`

	d, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		log.Fatal(err)
	}

	d.Find("#div1").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}

func f3() {
	html := `<body>
	<div id="div1">div1</div>
	<div class="name">manto</div>
	<div>div2</div>
	<span>Span</span>
	</body>
	`

	d, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		log.Fatal(err)
	}

	d.Find(".name").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})

}

func f4() {
	content := `
	<p>最近网上很火的梗就是，都是核酸阴性才配跟我聊天，今天小编就给大家带来一组网上很火的聊天表情给大家啦。</p>
	<p align="center"><img src="https://p.qqan.com/up/2022-8/16608706515165434.jpg"/></p>
	<p align="center"><img src="https://p.qqan.com/up/2022-8/16608706517702452.jpg"/></p>
	<p align="center"><img src="https://p.qqan.com/up/2022-8/16608706512361173.jpg"/></p>
	<p align="center"><img src="https://p.qqan.com/up/2022-8/16608706514898191.jpg"/></p>
	<p align="center"><img src="https://p.qqan.com/up/2022-8/16608706516874357.jpg"/></p>
	<p align="center"><img src="https://p.qqan.com/up/2022-8/16608706514630949.jpg"/></p>
	<p align="center"><img src="https://p.qqan.com/up/2022-8/16608706514777753.jpg"/></p>
	p align="center"><img src="https://p.qqan.com/up/2022-8/16608706515982130.jpg"/></p><p align="center"><img src="https://p.qqan.com/up/2022-8/16608706518519148.jpg"/></p>
	`

	d, err := goquery.NewDocumentFromReader(strings.NewReader(content))

	if err != nil {
		log.Fatal(err)
	}
	d.Find("img").Each(func(i int, s *goquery.Selection) {
		img, _ := s.Attr("src")
		fmt.Printf("img: %v\n", img)
	})
}

func main() {
	// f2()
	// f3()
	f4()
}
*/