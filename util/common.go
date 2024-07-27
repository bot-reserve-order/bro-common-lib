package util

import (
	"fmt"
	"runtime"
)

func Caller() string {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s#%d\n", file, no)
	}
	return ""
}
