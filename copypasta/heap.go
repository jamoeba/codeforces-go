package copypasta

import (
	"container/heap"
	"sort"
)

/*
思维转换
https://www.luogu.com.cn/problem/P2859
https://www.luogu.com.cn/problem/P4952 枚举中位数
https://codeforces.com/contest/713/problem/C 使序列严格递增的最小操作次数 (+1/-1)
    https://codeforces.com/blog/entry/47094?#comment-315068
    https://codeforces.com/blog/entry/77298 Slope trick
https://codeforces.com/problemset/problem/884/D 从结果倒推（类似霍夫曼编码）

求前缀/后缀的最小的 k 个元素和（k 固定）https://www.luogu.com.cn/problem/P4952 https://www.luogu.com.cn/problem/P3963
滑动窗口中位数 LC480 https://leetcode-cn.com/problems/sliding-window-median/

基于堆的反悔贪心（反悔堆）
https://www.cnblogs.com/nth-element/p/11768155.html
https://codeforces.com/problemset/problem/1526/C2
JSOI07 建筑抢修 https://www.luogu.com.cn/problem/P4053 LC630 https://leetcode-cn.com/problems/course-schedule-iii/
*/

// 下面这些都是最小堆

type hp struct{ sort.IntSlice }

//func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] } // 加上这行变成最大堆
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) push(v int)         { heap.Push(h, v) }
func (h *hp) pop() int           { return heap.Pop(h).(int) }
func (h *hp) popPush(v int) int  { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(h, 0); return top } // h 需要非空
func (h *hp) pushPop(v int) int {
	if len(h.IntSlice) > 0 && v > h.IntSlice[0] { // 最大堆改成 v < h.IntSlice[0]
		v, h.IntSlice[0] = h.IntSlice[0], v
		heap.Fix(h, 0)
	}
	return v
}

//

type hp64 []int64 // 自定义类型

func (h hp64) Len() int               { return len(h) }
func (h hp64) Less(i, j int) bool     { return h[i] < h[j] } // > 为最大堆
func (h hp64) Swap(i, j int)          { h[i], h[j] = h[j], h[i] }
func (h *hp64) Push(v interface{})    { *h = append(*h, v.(int64)) }
func (h *hp64) Pop() interface{}      { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp64) push(v int64)          { heap.Push(h, v) }
func (h *hp64) pop() int64            { return heap.Pop(h).(int64) }
func (h *hp64) popPush(v int64) int64 { top := (*h)[0]; (*h)[0] = v; heap.Fix(h, 0); return top } // h 需要非空
func (h *hp64) pushPop(v int64) int64 {
	if len(*h) > 0 && v > (*h)[0] { // 最大堆改成 v < (*h)[0]
		v, (*h)[0] = (*h)[0], v
		heap.Fix(h, 0)
	}
	return v
}

//

// 支持修改、删除指定元素的堆
// 用法：调用 push 会返回一个 *viPair 指针，记作 p
// 将 p 存于他处（如 slice 或 map），可直接在外部修改 p.v 后调用 fix(p.index)，从而做到修改堆中指定元素
// 调用 remove(p.index) 可以从堆中删除 p.v
// 例题 https://atcoder.jp/contests/abc170/tasks/abc170_e
type viPair struct {
	v     int64
	index int // 该元素在 mh 中的下标，可随着 push pop 等操作自动改变
}
type mh []*viPair // mh 指 modifiable heap

func (h mh) Len() int              { return len(h) }
func (h mh) Less(i, j int) bool    { return h[i].v < h[j].v } // > 为最大堆
func (h mh) Swap(i, j int)         { h[i], h[j] = h[j], h[i]; h[i].index = i; h[j].index = j }
func (h *mh) Push(v interface{})   { *h = append(*h, v.(*viPair)) }
func (h *mh) Pop() interface{}     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *mh) push(v int64) *viPair { p := &viPair{v, len(*h)}; heap.Push(h, p); return p }
func (h *mh) pop() *viPair         { return heap.Pop(h).(*viPair) }
func (h *mh) fix(i int)            { heap.Fix(h, i) }
func (h *mh) remove(i int) *viPair { return heap.Remove(h, i).(*viPair) }

// 对顶堆求动态中位数：medians[i] = a[:i+1] 的中位数
// https://www.luogu.com.cn/problem/P1168
// LC295 https://leetcode-cn.com/problems/find-median-from-data-stream/
// 与树状数组结合 https://leetcode-cn.com/contest/season/2020-fall/problems/5TxKeK/
func dynamicMedians(a []int) []int {
	n := len(a)
	medians := make([]int, 0, n)
	var small, big hp
	for _, v := range a {
		if len(small.IntSlice) == len(big.IntSlice) {
			big.push(-small.pushPop(-v))
		} else {
			small.push(-big.pushPop(v))
		}
		medians = append(medians, big.IntSlice[0])
	}
	return medians
}
