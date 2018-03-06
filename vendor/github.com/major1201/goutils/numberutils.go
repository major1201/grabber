package goutils

import "strconv"

func ToInt(s string) int {
	return ToIntDv(s, 0)
}

func ToIntDv(s string, dv int) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		return dv
	}
	return i
}
