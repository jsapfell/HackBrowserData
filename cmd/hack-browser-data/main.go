package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/josa42/go-applescript"
	"github.com/urfave/cli/v2"

	"github.com/moond4rk/HackBrowserData/browser"
	"github.com/moond4rk/HackBrowserData/log"
	"github.com/moond4rk/HackBrowserData/utils/fileutil"
)

var (
	browserName  string
	outputDir    string
	outputFormat string
	verbose      bool
	memo         bool
	compress     bool
	profilePath  string
	isFullExport bool
)

func main() {
	Execute()
}

func Execute() {
	app := &cli.App{
		Name:      "hack-browser-data",
		Usage:     "Export password|bookmark|cookie|history|credit card|download|localStorage|extension from browser",
		UsageText: "[hack-browser-data -b chrome -f json -dir results -cc]\nExport all browingdata(password/cookie/history/bookmark) from browser\nGithub Link: https://github.com/moonD4rk/HackBrowserData",
		Version:   "0.5.0",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "verbose", Aliases: []string{"vv"}, Destination: &verbose, Value: false, Usage: "verbose"},
			&cli.BoolFlag{Name: "compress", Aliases: []string{"zip"}, Destination: &compress, Value: false, Usage: "compress result to zip"},
			&cli.BoolFlag{Name: "getmemo", Aliases: []string{"memo"}, Destination: &memo, Value: false, Usage: "get memo info"},
			&cli.StringFlag{Name: "browser", Aliases: []string{"b"}, Destination: &browserName, Value: "all", Usage: "available browsers: all|" + browser.Names()},
			&cli.StringFlag{Name: "results-dir", Aliases: []string{"dir"}, Destination: &outputDir, Value: "results", Usage: "export dir"},
			&cli.StringFlag{Name: "format", Aliases: []string{"f"}, Destination: &outputFormat, Value: "csv", Usage: "file name csv|json"},
			&cli.StringFlag{Name: "profile-path", Aliases: []string{"p"}, Destination: &profilePath, Value: "", Usage: "custom profile dir path, get with chrome://version"},
			&cli.BoolFlag{Name: "full-export", Aliases: []string{"full"}, Destination: &isFullExport, Value: true, Usage: "is export full browsing data"},
		},
		HideHelpCommand: true,
		Action: func(c *cli.Context) error {
			if verbose {
				log.SetVerbose()
			}
			browsers, err := browser.PickBrowsers(browserName, profilePath)
			if err != nil {
				log.Error(err)
			}

			for _, b := range browsers {
				data, err := b.BrowsingData(isFullExport)
				if err != nil {
					log.Error(err)
					continue
				}
				data.Output(outputDir, b.Name(), outputFormat)
			}

			if compress {
				if err = fileutil.CompressDir(outputDir); err != nil {
					log.Error(err)
				}
				log.Noticef("compress success")
			}
			if memo {
				getMemoInfo()
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func getMemoInfo() {
	script := `
    tell application "Notes"
        set notesList to every note
        set notesData to {}
        repeat with aNote in notesList
            set noteTitle to the name of aNote
            set noteContent to the body of aNote
            set noteData to {title:noteTitle, content:noteContent}
            set end of notesData to noteData
        end repeat
        return notesData
    end tell
    `

	result, err := applescript.Run(script)
	if err != nil {
		fmt.Println(err)
		return
	}
	exePath, _ := os.Executable()
	exePath, _ = filepath.Abs(exePath)
	currPath := filepath.Dir(exePath)
	os.Mkdir(path.Join(currPath, outputDir), os.ModePerm)
	memoFile := path.Join(currPath, outputDir, "memo.html")
	ioutil.WriteFile(memoFile, []byte(result), os.ModePerm)
}
