package main

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	t "./tiktok"
	tui "./tui"
)

func main() {
	tiktoks := t.FetchTikTokTrends()

	tui.SetupTUI(tiktoks)

	// tiktoks2 := t.FetchTikTokTrends(strconv.Itoa(2))
	//	toks := unique(append(tiktoks, tiktoks2...))

}

func unique(slice []string) []string {
	encountered := map[string]int{}
	diff := []string{}

	for _, v := range slice {
		encountered[v] = encountered[v] + 1
	}

	for _, v := range slice {
		if encountered[v] == 1 {
			diff = append(diff, v)
		}
	}
	return diff
}
