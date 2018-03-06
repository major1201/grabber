package goutils

import "reflect"

func Contains(obj interface{}, v ...interface{}) bool {
	for _, o := range v {
		if obj == o {
			return true
		}
	}
	return false
}

func DeepContains(obj interface{}, v ...interface{}) bool {
	for _, o := range v {
		if reflect.DeepEqual(o, obj) {
			return true
		}
	}
	return false
}
