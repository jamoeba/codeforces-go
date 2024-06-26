// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/339/E
// https://codeforces.com/problemset/status/339/problem/E
func Test_cf339E(t *testing.T) {
	testCases := [][2]string{
		{
			`5
1 4 3 2 5`,
			`1
2 4`,
		},
		{
			`6
2 1 4 3 6 5`,
			`3
1 2
3 4
5 6`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf339E)
}
