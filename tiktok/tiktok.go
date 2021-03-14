package tiktok

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func FetchTikTokTrends() ([]string, []string, []string, []float64, []string) {
	requestUrl := "https://m.tiktok.com/api/recommend/item_list/?aid=1988&count=35"
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}

	var linksSlice []string
	var descSlice []string
	var usersSlice []string
	var datesSlice []float64
	var statsSlice []string

	// collecting around ~90-100 videos at 1 time
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

			video, ok := m["video"].(map[string]interface{})
			if !ok {
				fmt.Println("fatal-3")
				continue
			}
			link, ok := video["playAddr"].(string)
			if !ok {
				fmt.Println("fatal-4")
				continue
			}

			desc, ok := m["desc"].(string)
			if !ok {
				fmt.Println("fatal-5")
				continue
			}

			user, ok := m["author"].(map[string]interface{})
			if !ok {
				fmt.Println("fatal-6")
				continue
			}

			id, ok := user["uniqueId"].(string)
			if !ok {
				fmt.Println("fatal-7")
				continue
			}

			userStats, ok := m["authorStats"].(map[string]interface{})
			if !ok {
				fmt.Println("fatal-8")
				continue
			}

			followers, ok := userStats["followerCount"].(float64)
			if !ok {
				fmt.Println("fatal-9")
				continue
			}

			date, ok := m["createTime"].(float64)
			if !ok {
				fmt.Println("fatal-10")
				continue
			}

			stats, ok := m["stats"].(map[string]interface{})
			if !ok {
				fmt.Println("fatal-11")
				continue
			}

			plays, ok := stats["playCount"].(float64)
			if !ok {
				fmt.Println("fatal-12")
				continue
			}

			shares, ok := stats["shareCount"].(float64)
			if !ok {
				fmt.Println("fatal-13")
				continue
			}

			comments, ok := stats["commentCount"].(float64)
			if !ok {
				fmt.Println("fatal-14")
				continue
			}

			statsStr := "Plays: " + strconv.FormatFloat(plays, 'f', 0, 64) + " Shares: " + strconv.FormatFloat(shares, 'f', 0, 64) + " Comments: " + strconv.FormatFloat(comments, 'f', 0, 64)
			userStr := "User: @" + id + " Followers: " + strconv.FormatFloat(followers, 'f', 0, 64)

			usersSlice = append(usersSlice, userStr)
			statsSlice = append(statsSlice, statsStr)
			datesSlice = append(datesSlice, date)
			descSlice = append(descSlice, desc)
			linksSlice = append(linksSlice, link)
		}
	}

	return linksSlice, descSlice, usersSlice, datesSlice, statsSlice
}
