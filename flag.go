package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/major1201/goutils"
	"os"
)

// GrabberFlags defines the command line flags
type GrabberFlags struct {
	Help       bool
	Version    bool
	Urls       arrFlag
	Dest       string
	Yes        bool
	Concurrent int
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
	flag.BoolVar(&gf.Help, "help", false, "print command usage")
	flag.BoolVar(&gf.Version, "version", false, "print version")
	flag.Var(&gf.Urls, "url", "urls to download")
	flag.BoolVar(&gf.Yes, "y", false, "ignore confirmation")
	flag.IntVar(&gf.Concurrent, "j", 5, "Run n jobs in parallel, default 5")
	pwd, _ := os.Getwd()
	flag.StringVar(&gf.Dest, "dest", pwd, "specify the download directory, default: current working directory")
	flag.Parse()

	if gf.Help == true || len(os.Args) == 1 {
		fmt.Println("grabber - a simple concurrent download tool written in go (https://github.com/major1201/grabber)")
		flag.Usage()
		os.Exit(0)
	}
	if gf.Version == true {
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
	if len(gf.Urls) == 0 {
		return errors.New("at lease one url should be specified")
	}
	if len(gf.Dest) == 0 {
		return errors.New("please specify the download directory")
	}
	if !goutils.IsDir(gf.Dest) {
		return errors.New(gf.Dest + "is not a directory")
	}
	return nil
}
