package checkout

import "strings"

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
