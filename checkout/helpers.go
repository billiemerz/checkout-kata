package checkout

import (
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

/*
ErrorContains is a helper function to check errors in tests
*/
func ErrorContains(out error, want error) bool {
	if out == nil {
		return want == nil
	}
	if want == nil {
		return false
	}
	return strings.Contains(out.Error(), want.Error())
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
