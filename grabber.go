package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/major1201/goutils"
	"github.com/major1201/goutils/logging"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var logger = logging.New("SYS")

// ProjectName means main project name
const ProjectName = "grabber"

// Version means main project version
const Version = "0.1.1"

var wg = sync.WaitGroup{}

func confirmation(urls []string) {
	if !gf.Yes {
		for _, url := range urls {
			fmt.Println(url, "==>", filepath.Join(gf.Dest, filepath.Base(url)))
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

func downloadOne(ch chan int, url string) {
	defer func() {
		fmt.Println(url, " ==> ok")
		<-ch
		wg.Done()
	}()

	ch <- 0

	e := goutils.Download(url, filepath.Join(gf.Dest, filepath.Base(url)))
	if e != nil {
		logger.Error(e)
	}
}

func download(urls []string) {
	ch := make(chan int, gf.Concurrent)
	for _, url := range urls {
		wg.Add(1)
		go downloadOne(ch, url)
	}
	wg.Wait()
}

// ResolveAddr parses a string to multiple url strings
func ResolveAddr(exp string) []string {
	defaultRet := []string{exp}
	bytesExp := []byte(exp)
	bracketIndex1 := bytes.IndexByte(bytesExp, '[')
	if bracketIndex1 == -1 {
		return defaultRet
	}
	prefix := string(bytesExp[:bracketIndex1])
	bracketIndex2 := bracketIndex1 + bytes.IndexByte(bytesExp[bracketIndex1:], ']')
	if bracketIndex2 == -1 {
		return defaultRet
	}
	postfix := string(bytesExp[bracketIndex2+1:])
	subexp1 := string(bytesExp[bracketIndex1+1 : bracketIndex2])
	subexp2 := strings.Split(subexp1, ":")
	if len(subexp2) != 2 {
		return defaultRet
	}
	fromStr := subexp2[0]
	toStr := subexp2[1]
	from, errFrom := strconv.Atoi(fromStr)
	if errFrom != nil {
		return defaultRet
	}
	to, errTo := strconv.Atoi(toStr)
	if errTo != nil {
		return defaultRet
	}
	if from > to || from < 0 || to < 0 {
		return defaultRet
	}
	minLength := len(fromStr)
	if minLength > len(toStr) {
		return defaultRet
	}
	var ret []string
	for i := from; i <= to; i++ {
		ret = append(ret, prefix+goutils.ZeroFill(strconv.Itoa(i), minLength)+postfix)
	}
	return ret
}

func main() {
	// init logging
	logging.AddStdout(0)

	parseFlags()

	var urls []string
	for _, url := range gf.Urls {
		urlClips := ResolveAddr(url)
		urls = append(urls, urlClips...)
	}

	confirmation(urls)

	// download
	logger.Info("==> Started.")
	download(urls)
	logger.Info("==> Done.")
}
