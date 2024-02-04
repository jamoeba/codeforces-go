// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1822/E
// https://codeforces.com/problemset/status/1822/problem/E
func Test_cf1822E(t *testing.T) {
	testCases := [][2]string{
		{
			`10
10
codeforces
3
abc
10
taarrrataa
10
dcbdbdcccc
4
wwww
12
cabbaccabaac
10
aadaaaaddc
14
aacdaaaacadcdc
6
abccba
12
dcbcaebacccd`,
			`0
-1
1
1
-1
3
-1
2
2
2`,
		},
		{
			`1
10
aadaaaaddc`,
			`-1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1822E)
}