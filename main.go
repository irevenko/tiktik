package main

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	t "./tiktok"
	tui "./tui"
)

func main() {
	links, descs, ids := t.FetchTikTokTrends()

	tui.SetupTUI(links, descs, ids)
}
