初始化 $i=j=0$，按题意要求上下左右移动即可，注意题目保证不会出界。

最后返回 $i\cdot n+j$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Cf421v7Ky/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def finalPositionOfSnake(self, n: int, commands: List[str]) -> int:
        i = j = 0
        for s in commands:
            if s[0] == 'U': i -= 1
            elif s[0] == 'D': i += 1
            elif s[0] == 'L': j -= 1
            else: j += 1
        return i * n + j
```

```py [sol-Python3 match-case]
class Solution:
    def finalPositionOfSnake(self, n: int, commands: List[str]) -> int:
        i = j = 0
        for s in commands:
            match s[0]:
                case 'U': i -= 1
                case 'D': i += 1
                case 'L': j -= 1
                case _:   j += 1  # 匹配其他任意值，相当于 switch 中的 default
        return i * n + j
```

```java [sol-Java]
class Solution {
    public int finalPositionOfSnake(int n, List<String> commands) {
        int i = 0;
        int j = 0;
        for (String s : commands) {
            switch (s.charAt(0)) {
                case 'U' -> i--;
                case 'D' -> i++;
                case 'L' -> j--;
                default  -> j++;
            }
        }
        return i * n + j;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int finalPositionOfSnake(int n, vector<string>& commands) {
        int i = 0, j = 0;
        for (auto& s : commands) {
            switch (s[0]) {
                case 'U': i--; break;
                case 'D': i++; break;
                case 'L': j--; break;
                default:  j++;
            }
        }
        return i * n + j;
    }
};
```

```go [sol-Go]
func finalPositionOfSnake(n int, commands []string) int {
	i, j := 0, 0
	for _, s := range commands {
		switch s[0] {
		case 'U': i--
		case 'D': i++
		case 'L': j--
		default:  j++
		}
	}
	return i*n + j
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m)$，其中 $n$ 是 $\textit{commands}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)