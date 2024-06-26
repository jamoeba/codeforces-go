// Generated by copypasta/template/generator_test.go
package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// https://codeforces.com/problemset/problem/1789/D
// https://codeforces.com/problemset/status/1789/problem/D
func Test_cf1789D(t *testing.T) {
	testCases := [][2]string{
		{
			`3
5
00111
11000
1
1
1
3
001
000`,
			`2
3 -2
0
-1`,
		},
		{
			`1
3
010
001`,
			`3
-1 1 2`,
		},
		{
			`1
5
00100
00011`,
			`4
-1 -2 2 3 `,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1789D)
}

func TestCheck_cf1789D(_t *testing.T) {
	//return
	lsh := func(s string, k int) string { return s[k:] + strings.Repeat("0", k) }
	rsh := func(s string, k int) string { return strings.Repeat("0", k) + s[:len(s)-k] }
	xor := func(s, t string) string {
		x := []byte(s)
		for i := range x {
			x[i] ^= t[i] ^ '0'
		}
		return string(x)
	}
	assert := assert.New(_t)
	_ = assert
	testutil.DebugTLE = 0
	rg := testutil.NewRandGenerator()
	inputGenerator := func() (string, testutil.OutputChecker) {
		rg.Clear()
		rg.One()
		n := rg.Int(1, 9)
		rg.NewLine()
		s := rg.Str(n,n,'0','1')
		rg.NewLine()
		t := rg.Str(n,n,'0','1')
		return rg.String(), func(myOutput string) (_b bool) {
			// 检查 myOutput 是否符合题目要求
			// * 最好重新看一遍题目描述以免漏判 *
			// 对于 special judge 的题目，可能还需要额外跑个暴力来检查 myOutput 是否满足最优解等
			in := strings.NewReader(myOutput)
			var m ,k int
			Fscan(in, &m)
			if m < 0 {
				return true
			}
			if m > n {
				return 
			}
			for i := 0; i < m; i++ {
				Fscan(in, &k)
				if k == 0 {
					return 
				}
				if k > 0 {
					s = xor(s, lsh(s, k))
				} else {
					s = xor(s, rsh(s, -k))
				}
			}
			return assert.EqualValues(t, s)
		}
	}

	testutil.CheckRunResultsInfWithTarget(_t, inputGenerator, 0, cf1789D)
}
