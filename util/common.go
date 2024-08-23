package util

import (
	"fmt"
	"runtime"
)

func Caller(skip int) string {
	_, file, no, ok := runtime.Caller(skip)
	if ok {
		return fmt.Sprintf("%s#%d\n", file, no)
	}
	return ""
}
