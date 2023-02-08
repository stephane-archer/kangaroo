package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"os/exec"
	"strings"
)

func loadConfig(configPath string) *ini.File {
	cfg, err := ini.Load(configPath)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return cfg
}

func main() {
	cfg := loadConfig("config.ini")
	windowName := cfg.Section("").Key("WindowName").String()
	a := app.New()
	w := a.NewWindow(windowName)

	processFileSelected := func(file fyne.URIReadCloser, err error) {
		if err != nil {
			w.SetContent(widget.NewLabel(err.Error()))
			return
		}
		if file == nil {
			log.Fatal("Cancelled")
			return
		}
		infinite := widget.NewProgressBarInfinite()
		w.SetContent(infinite)
		cliToolPath := cfg.Section("").Key("CliToolPath").String()
		cmd := exec.Command(cliToolPath, file.URI().Path())
		var out strings.Builder
		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
			return
		}
		outStr := out.String()
		fmt.Printf("%v\n", outStr)
		entry := widget.NewEntry()
		entry.SetText(outStr)
		w.SetContent(entry)
	}

	windowWidth := cfg.Section("").Key("WindowWidth").MustInt(900)
	windowHeight := cfg.Section("").Key("WindowHeight").MustInt(500)
	windowSize := fyne.NewSize(float32(windowWidth), float32(windowHeight))

	fileDialog := dialog.NewFileOpen(processFileSelected, w)
	fileDialog.Resize(windowSize)
	fileDialog.Show()

	w.Resize(windowSize)
	w.Show()

	a.Run()
}
