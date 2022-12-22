// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/abc106/submit?taskScreenName=abc106_d
func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`2 3 1
1 1
1 2
2 2
1 2`,
			`3`,
		},
		{
			`10 3 2
1 5
2 8
7 10
1 7
3 10`,
			`1
1`,
		},
		{
			`10 10 10
1 6
2 9
4 5
4 7
4 7
5 8
6 6
6 7
7 9
10 10
1 8
1 9
1 10
2 8
2 9
2 10
3 8
3 9
3 10
1 10`,
			`7
9
10
6
8
9
6
7
8
10`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/abc106/tasks/abc106_d