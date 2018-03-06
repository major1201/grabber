package goutils

import (
	"strings"
	"regexp"
	"github.com/google/uuid"
	"reflect"
)

const EmptyStr = ""

func IsEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

func IsBlank(s string) bool {
	if len(s) == 0 {
		return true
	}
	reg := regexp.MustCompile(`^\s+$`)
	actual := reg.MatchString(s)
	if actual {
		return true
	}
	return false
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

func Trim(str string) string {
	return strings.Trim(str, " ")
}

func TrimLeft(str string) string {
	return strings.TrimLeft(str, " ")
}

func TrimRight(str string) string {
	return strings.TrimRight(str, " ")
}

func LeftPad(s string, padStr string, length int) string {
	prefix := EmptyStr
	if len(s) < length {
		prefix = strings.Repeat(padStr, length-len(s))
	}
	return prefix + s
}

func RightPad(s string, padStr string, length int) string {
	postfix := EmptyStr
	if len(s) < length {
		postfix = strings.Repeat(padStr, length-len(s))
	}
	return s + postfix
}

func ZeroFill(s string, length int) string {
	const zeroStr = "0"
	return LeftPad(s, zeroStr, length)
}

func Len(s string) int {
	return len([]rune(s))
}

func Index(s, substr string) int {
	sRune := []rune(s)
	subRune := []rune(substr)
	if len(subRune) > len(sRune) {
		return -1
	}
	for i := 0; i < len(sRune)-len(subRune)+1; i++ {
		if reflect.DeepEqual(sRune[i:i+len(subRune)], subRune) {
			return i
		}
	}
	return -1
}

func Between(s, from, to string) string {
	indexFrom := Index(s, from)
	if indexFrom == -1 {
		return ""
	}
	runeS := []rune(s)
	indexTo := Index(string(runeS[indexFrom:]), to)
	if indexTo == -1 {
		return ""
	}
	return string(runeS[indexFrom+Len(from):indexFrom+indexTo])
}

func UUID() string {
	return strings.Replace(uuid.New().String(), "-", "", 4)
}
