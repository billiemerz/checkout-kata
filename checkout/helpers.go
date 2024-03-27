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

/*
rootDir returns the root directory of the project. This will
remain the same regardless of where the code is called from -
i.e. tests/ a main package etc.
*/
func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
