package tui

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

func SetupTUI(tiktoks []string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	l := widgets.NewList()

	l.Title = "tiktik"
	l.Rows = tiktoks
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = true
	l.SetRect(0, 0, 150, 35)

	ui.Render(l)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "z", "<Enter>":
			OpenBrowser(l.Rows[l.SelectedRow])
		}

		ui.Render(l)
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
