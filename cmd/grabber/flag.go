package main

import (
	"flag"
	"fmt"
	"os"
	"errors"
	"github.com/major1201/goutils"
)

type GrabberFlags struct {
	help bool
	version bool
	urls arrFlag
	dest string
	yes bool
}

type arrFlag []string

func (v *arrFlag) String() string {
	return fmt.Sprintf("%s", *v)
}

func (v *arrFlag) Set(s string) error {
	*v = append(*v, s)
	return nil
}

var gf = &GrabberFlags{}

func parseFlags() {
	flag.BoolVar(&gf.help, "help", false, "print command usage")
	flag.BoolVar(&gf.version, "version", false, "print version")
	flag.Var(&gf.urls, "url", "urls to download")
	flag.BoolVar(&gf.yes, "y", false, "ignore confirmation")
	pwd, _ := os.Getwd()
	flag.StringVar(&gf.dest, "dest", pwd, "specify the download directory, default: current working directory")
	flag.Parse()

	if gf.help == true || len(os.Args) == 1 {
		fmt.Println("grabber - a simple concurrent download tool written in go (https://github.com/major1201/grabber)")
		flag.Usage()
		os.Exit(0)
	}
	if gf.version == true {
		fmt.Println(ProjectName, "-", Version)
		os.Exit(0)
	}

	// check flags validation
	err := checkFlags()
	if err != nil {
		logger.Fatal(err)
	}
}

func checkFlags() error {
	if len(gf.urls) == 0 {
		return errors.New("at lease one url should be specified")
	}
	if len(gf.dest) == 0 {
		return errors.New("please specify the download directory")
	}
	if !goutils.IsDir(gf.dest) {
		return errors.New(gf.dest + "is not a directory")
	}
	return nil
}
