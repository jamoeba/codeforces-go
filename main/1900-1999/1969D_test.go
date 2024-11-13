// Generated by copypasta/template/generator_test.go
package main

import (
	"container/heap"
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"sort"
	"testing"
)

// https://codeforces.com/problemset/problem/1969/D
// https://codeforces.com/problemset/status/1969/problem/D?friends=on
func Test_cf1969D(t *testing.T) {
	testCases := [][2]string{
		{
			`4
2 0
2 1
1 2
4 1
1 2 1 4
3 3 2 3
4 2
2 1 1 1
4 2 3 2
6 2
1 3 4 9 1 3
7 6 8 10 6 8`,
			`1
1
0
7`,
		},
		{
			`1
4 1
1 4 1 1
3 5 2 3`,
			`2`,
		},
		{
			`1
3 1
1 1 5
3 3 9`,
			`1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1969D)
}

func TestCompare_cf1969D(_t *testing.T) {
	//return
	testutil.DebugTLE = 0
	rg := testutil.NewRandGenerator()
	inputGenerator := func() string {
		//return ``
		rg.Clear()
		rg.One()
		n := rg.Int(1, 9)
		rg.Int(0, n)
		rg.NewLine()
		rg.IntSlice(n, 1, 99)
		rg.IntSlice(n, 1, 99)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var T, n, k int
		for Fscan(in, &T); T > 0; T-- {
			Fscan(in, &n, &k)
			a := make([]int, n)
			for i := range a {
				Fscan(in, &a[i])
			}
			b := make([]int, n)
			for i := range b {
				Fscan(in, &b[i])
			}
			Fprintln(out, solve69(k, a,b))
		}
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, cf1969D)
}

func solve69(k int, a []int, b []int) int {
	n := len(a)
	type Pair struct {
		first  int
		second int
	}
	arr := make([]Pair, n)
	var profit int
	for i := 0; i < n; i++ {
		arr[i] = Pair{a[i], b[i]}
		profit += max(b[i]-a[i], 0)
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].second > arr[j].second
	})

	var best int

	if k == 0 {
		best = profit
	}

	var f int

	pq := make(IntHeap69, 0, n)

	for _, cur := range arr {
		profit -= max(0, cur.second-cur.first)
		heap.Push(&pq, cur.first)
		f += cur.first
		if pq.Len() > k {
			f -= heap.Pop(&pq).(int)
		}
		if pq.Len() == k {
			best = max(profit-f, best)
		}
	}

	return best
}

type IntHeap69 []int

func (h IntHeap69) Len() int           { return len(h) }
func (h IntHeap69) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap69) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap69) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap69) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
