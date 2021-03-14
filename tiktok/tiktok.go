package tiktok

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func FetchTikTokTrends() ([]string, []string, []string) {
	requestUrl := "https://m.tiktok.com/api/recommend/item_list/?aid=1988&count=35"
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}

	var linksSlice []string
	var descSlice []string
	var usersSlice []string

	for i := 0; i < 3; i++ {
		req, err := http.NewRequest(method, requestUrl, payload)
		if err != nil {
			fmt.Println(err)
			continue
		}
		req.Header.Add("cookie", " tt_webid_v2=1")
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		data := map[string]interface{}{}

		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = res.Body.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
		it, ok := data["itemList"].([]interface{})
		if !ok {
			fmt.Println("fatal-1")
			continue
		}

		for _, v := range it {
			m, ok := v.(map[string]interface{})
			if !ok {
				fmt.Println("fatal-2")
				continue
			}

			v, ok := m["video"].(map[string]interface{})
			if !ok {
				fmt.Println("fatal-3")
				continue
			}
			link, ok := v["playAddr"].(string)
			if !ok {
				fmt.Println("fatal-4")
				continue
			}

			desc, _ := m["desc"].(string)

			users, ok := m["author"].(map[string]interface{})
			if !ok {
				fmt.Println("fatal-7")
				continue
			}
			id, ok := users["uniqueId"].(string)
			if !ok {
				fmt.Println("fatal-8")
				continue
			}

			usersSlice = append(usersSlice, id)
			descSlice = append(descSlice, desc)
			linksSlice = append(linksSlice, link)
		}
	}

	return linksSlice, descSlice, usersSlice
}
