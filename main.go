package main

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	t "./tiktok"
	tui "./tui"
)

func main() {
	links, descs, users, dates, stats := t.FetchTikTokTrends()

	tui.SetupTUI(links, descs, users, dates, stats)
}
