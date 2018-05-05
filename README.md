# grabber
A simple concurrent downloader written in go

[![GoDoc](https://godoc.org/github.com/major1201/grabber?status.svg)](https://godoc.org/github.com/major1201/grabber)
[![Go Report Card](https://goreportcard.com/badge/github.com/major1201/grabber)](https://goreportcard.com/report/github.com/major1201/grabber)

## How to install

### Build and Install the Binaries from Source

```
go get -v github.com/major1201/grabber
```

## How to use

**Example1**

```
grabber -url https://www.example.com/img_[1:3].jpg
```

This will download img_1.jpg, img_2.jpg, img_3.jpg to the current wording directory

**Example2**

```
grabber -url https://www.example.com/img_[08:11].jpg -dest /tmp
```

This will download img_08.jpg, img_09.jpg, img_10.jpg, img_11.jpg to `/tmp`
