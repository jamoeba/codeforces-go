// Generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_c(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, largestPalindrome, "c.txt", 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-411/problems/find-the-largest-palindrome-divisible-by-k/
// https://leetcode.cn/problems/find-the-largest-palindrome-divisible-by-k/