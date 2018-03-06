package grabber

import (
	"github.com/major1201/goutils"
	"strings"
	"bytes"
	"strconv"
)

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
	subexp1 := string(bytesExp[bracketIndex1+1:bracketIndex2])
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
		ret = append(ret, prefix + goutils.ZeroFill(strconv.Itoa(i), minLength) + postfix)
	}
	return ret
}
