package tiktok

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	t "../types"
)

func FetchTikTokTrends() []string {
	baseTikTokURL := "https://tiktok.com/@"

	//url3 := "https://m.tiktok.com/api/recommend/item_list/?aid=1988&app_name=tiktok_web&device_platform=web&referer=&root_referer=&user_agent=Mozilla%2F5.0+(X11%3B+Linux+x86_64)+AppleWebKit%2F537.36+(KHTML,+like+Gecko)+Chrome%2F88.0.4324.150+Safari%2F537.36&cookie_enabled=true&screen_width=1920&screen_height=1080&browser_language=en-En&browser_platform=Linux+x86_64&browser_name=Mozilla&browser_online=true&ac=4g&page_referer=https:%2F%2Fwww.tiktok.com%2Ftag%2Ftoiktok%3Flang%3Den&priority_region=&verifyFp=&appId=1233&region=US&appType=m&isAndroid=false&isMobile=false&isIOS=false&OS=linux&count=30&itemID=1&_signature="
	//url2 := "https://m.tiktok.com/api/recommend/item_list/?aid=1988&app_name=tiktok_web&device_platform=web&referer=&root_referer=&user_agent=Mozilla%2F5.0+(X11%3B+Linux+x86_64)+AppleWebKit%2F537.36+(KHTML,+like+Gecko)+Chrome%2F88.0.4324.150+Safari%2F537.36&cookie_enabled=true&screen_width=1920&screen_height=1080&browser_language=en-En&browser_platform=Linux+x86_64&browser_name=Mozilla&browser_online=true&ac=4g&page_referer=https:%2F%2Fwww.tiktok.com%2Ftag%2Ftoiktok%3Flang%3Den&priority_region=&verifyFp=&appId=1233&region=US&appType=m&isAndroid=false&isMobile=false&isIOS=false&OS=linux&count=30&itemID=1&_signature="
	//url := baseURL + "/?aid=1988&app_name=tiktok_web&device_platform=web&referer=&root_referer=&user_agent=Mozilla%2F5.0+(X11%3B+Linux+x86_64)+AppleWebKit%2F537.36+(KHTML,+like+Gecko)+Chrome%2F88.0.4324.150+Safari%2F537.36&cookie_enabled=true&screen_width=1920&screen_height=1080&browser_language=ru-RU&browser_platform=Linux+x86_64&browser_name=Mozilla&browser_version=5.0+(X11%3B+Linux+x86_64)+AppleWebKit%2F537.36+(KHTML,+like+Gecko)+Chrome%2F88.0.4324.150+Safari%2F537.36&browser_online=true&ac=4g&timezone_name=Europe%2FMoscow&page_referer=https:%2F%2Fwww.tiktok.com%2Ftag%2Ftoiktok%3Flang%3Den&priority_region=&verifyFp=&appId=1233&region=RU&appType=m&isAndroid=false&isMobile=false&isIOS=false&OS=linux&did=6930141893205345797&count=30&itemID=1&language=ru-RU&_signature="

	url := "https://m.tiktok.com/api/recommend/item_list/?aid=1988&app_name=tiktok_web&device_platform=web&referer=&root_referer=&user_agent=Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:86.0) Gecko/20100101 Firefox/86.0&cookie_enabled=true&screen_width=1920&screen_height=1080&browser_language=ru-RU&browser_platform=Linux+x86_64&browser_name=Mozilla&browser_version=5.0+(X11%3B+Linux+x86_64)+AppleWebKit%2F537.36+(KHTML,+like+Gecko)+Chrome%2F88.0.4324.150+Safari%2F537.36&browser_online=true&ac=4g&timezone_name=Europe%2FMoscow&page_referer=https:%2F%2Fwww.tiktok.com%2Ftag%2Ftoiktok%3Flang%3Den&priority_region=&verifyFp=&appId=1233&region=RU&appType=m&isAndroid=false&isMobile=false&isIOS=false&OS=linux&did=6930141893205345797&count=30&itemID=1&language=ru-RU"
	var tiktoks []string

	method := "GET"
	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatal(err)
	}

	// req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	// req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	// req.Header.Set("Cache-Control", "max-age=0")
	// req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Host", "m.tiktok.com")
	// req.Header.Set("TE", "Trailers")
	// req.Header.Set("Upgrade-Insecure-Requests", "1")
	// req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:86.0) Gecko/20100101 Firefox/86.0")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.StatusCode)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tiktokResp t.TikTokResponse
	json.Unmarshal(body, &tiktokResp)

	for _, v := range tiktokResp.ItemList {
		tiktoks = append(tiktoks, v.Desc+" "+baseTikTokURL+v.Author.UniqueID+"/video/"+v.ID)
	}

	return tiktoks
}
