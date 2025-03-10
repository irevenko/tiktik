package tui

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	t "github.com/irevenko/tiktik/tiktok"
)

func SetupTUI(links []string, descs []string, users []string, dates []float64, stats []string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	var tiktoks []string

	for i, v := range links {
		date := time.Unix(int64(dates[i]), 0)
		layout := "2006-01-02 15:04:05"
		formatted := date.Format(layout)

		tiktoks = append(tiktoks, strconv.Itoa(i+1)+") "+users[i]+" "+stats[i]+" [Date:](fg:magenta) "+formatted)
		if descs[i] == "" {
			tiktoks = append(tiktoks, strconv.Itoa(i+1)+") [Description:](fg:blue) "+"no desc")
		} else {
			tiktoks = append(tiktoks, strconv.Itoa(i+1)+") [Description:](fg:blue) "+descs[i])
		}
		tiktoks = append(tiktoks, strconv.Itoa(i+1)+") "+v)
	}

	l := widgets.NewList()
	l.Title = "TikTok Trends" + " - " + strconv.Itoa(len(links)) + " videos "
	l.TitleStyle = ui.NewStyle(ui.ColorCyan)
	l.TextStyle = ui.NewStyle(ui.ColorWhite)
	l.SelectedRowStyle = ui.NewStyle(ui.ColorCyan)
	l.Rows = tiktoks
	l.WrapText = false
	l.SetRect(0, 4, 150, 35)

	g := widgets.NewGauge()
	g.Title = "Downloading Progress"
	g.SetRect(0, 0, 150, 3)
	g.Percent = 0
	g.LabelStyle = ui.NewStyle(ui.ColorGreen)
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
			g.Percent = 0
			l.ScrollDown()
		case "k", "<Up>":
			g.Percent = 0
			l.ScrollUp()
		case "g", "<Home>":
			g.Percent = 0
			l.ScrollTop()
		case "G", "<End>":
			g.Percent = 0
			l.ScrollBottom()
		case "<C-f>", "<PageDown>":
			g.Percent = 0
			l.ScrollPageDown()
		case "<C-b>", "<PageUp>":
			g.Percent = 0
			l.ScrollPageUp()
		case "e", "<Enter>":
			if strings.Contains(l.Rows[l.SelectedRow], "https://") {
				url := strings.Split(l.Rows[l.SelectedRow], ") ")
				OpenBrowser(url[1])
			}
		case "d", "<C-d>":
			if strings.Contains(l.Rows[l.SelectedRow], "https://") {
				url := strings.Split(l.Rows[l.SelectedRow], ") ")
				t.DownloadTikTok(url[1], &g.Percent)
			}
		case "r", "<C-r>":
			ui.Close()
			links, descs, ids, dates, stats := t.FetchTikTokTrends()
			SetupTUI(links, descs, ids, dates, stats)
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
