/* package main
/// http库实现第一个爬虫
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fech(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	req.Header.Add("Cookie", "_ga=GA1.2.1657007290.1658374901; __gads=ID=70e504ae12532640:T=1659963137:S=ALNI_MatUQSNzXIJ7uupJwp1gEJSU1Ro1g; Hm_lvt_866c9be12d4a814454792b1fd0fed295=1658374900,1659088912,1659963137,1660791595; Hm_lpvt_866c9be12d4a814454792b1fd0fed295=1660791595; _gid=GA1.2.1083222976.1660791595; _gat_gtag_UA_476124_1=1; __gpi=UID=00000865782e9470:T=1659963137:RT=1660791595:S=ALNI_MZtnoTf4cAmlyBgPohndlnhjKtQsg; affinity=1660791599.557.39.860000|a6728cc07008ec0fd0d6b7ff6028a867; .AspNetCore.Session=CfDJ8EOBBtWq0dNFoDS+ZHPSe51ts26t96KCraB2s6VRzHv9AD0mkUHg0dBk2bZoM5c6phI470aNb0Yd/QDD2+coNor1n+cVLndninhVNDhn4qQk7taqBHadMJVM8annvG2Lc6eGdscGPGiJ5wG4hCcN/O6IAUL3HVEfKYpAUksu7ekZ; __utmc=59123430; __utmz=59123430.1660791602.1.1.utmcsr=cnblogs.com|utmccn=(referral)|utmcmd=referral|utmcct=/; __utmt=1; __utma=59123430.1657007290.1658374901.1660791602.1660791599.1; __utmb=59123430.1.10.1660791602")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err: ", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code: ", resp.StatusCode)
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return ""
	}

	return string(body)
}
func main() {
	url := "https://zzk.cnblogs.com/s?w=golang"

	s := fech(url)
	fmt.Printf("s: %v\n", s)
}
*/