package main

import (
	"github.com/major1201/grabber"
	"github.com/major1201/goutils/logging"
	"fmt"
	"bufio"
	"os"
	"github.com/major1201/goutils"
	"strings"
	"path/filepath"
)

var logger = logging.New("SYS")

const ProjectName = "grabber"
const Version = "0.1.0"

func confirmation(urls []string) {
	if !gf.yes {
		for _, url := range urls {
			fmt.Println(url, "==>", filepath.Join(gf.dest, filepath.Base(url)))
		}
		fmt.Println("-----------------------")
		reader := bufio.NewReader(os.Stdin)
	ReadInput:
		for {
			fmt.Printf("These %d urls would be downloaded, please confirm (y, n) ", len(urls))
			text, _ := reader.ReadString('\n')
			text = strings.ToLower(goutils.Trim(text))
			if len(text) > 1 {
				t := text[:1]
				switch t {
				case "y":
					break ReadInput
				case "n":
					os.Exit(2)
				}
			}
		}
	}
}

func download(urls []string) {
	chs := make([]chan int, len(urls))
	for i, url := range urls {
		addr := url
		chs[i] = make(chan int)
		go func(ch chan int) {
			defer func() {
				fmt.Println(addr, " ==> ok")
				ch <- 1
			}()
			e := goutils.Download(addr, filepath.Join(gf.dest, filepath.Base(addr)))
			if e != nil {
				logger.Error(e)
			}
		}(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}

func main() {
	// init logging
	logging.AddStdout(0)

	parseFlags()

	var urls []string
	for _, url := range gf.urls {
		urlClips := grabber.ResolveAddr(url)
		urls = append(urls, urlClips...)
	}

	confirmation(urls)

	// download
	logger.Info("==> Started.")
	download(urls)
	logger.Info("==> Done.")
}
