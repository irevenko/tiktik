package tui

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"

	t "../tiktok"
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

func SetupTUI(links []string, descs []string, ids []string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	l := widgets.NewList()

	var tiktoks []string

	for i, v := range links {
		tiktoks = append(tiktoks, "User: "+ids[i])
		tiktoks = append(tiktoks, "Desc: "+descs[i])
		tiktoks = append(tiktoks, v)
	}

	l.Title = "TikTok Trends"
	l.TitleStyle.Fg = ui.ColorCyan
	l.Rows = tiktoks
	l.WrapText = false
	l.SetRect(0, 5, 150, 35)

	g := widgets.NewGauge()
	g.Title = "Downloading Progress"
	g.SetRect(0, 0, 40, 4)
	g.Percent = 0
	g.BarColor = ui.ColorBlue
	g.BorderStyle.Fg = ui.ColorWhite
	g.TitleStyle.Fg = ui.ColorCyan

	ui.Render(l, g)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			l.ScrollDown()
			g.Percent = 0
		case "k", "<Up>":
			l.ScrollUp()
			g.Percent = 0
		case "e", "<Enter>":
			if strings.HasPrefix(l.Rows[l.SelectedRow], "http") {
				OpenBrowser(l.Rows[l.SelectedRow])
			}
		case "d":
			if strings.HasPrefix(l.Rows[l.SelectedRow], "http") {
				t.DownloadTikTok(l.Rows[l.SelectedRow], &g.Percent)
			}
		}

		ui.Render(l, g)
	}
}

func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
