# golang_happy_number

Write an algorithm to determine if a number `n` is happy.

A **happy number** is a number defined by the following process:

- Starting with any positive integer, replace the number by the sum of the squares of its digits.
- Repeat the process until the number equals 1 (where it will stay), or it **loops endlessly in a cycle** which does not include 1.
- Those numbers for which this process **ends in 1** are happy.

Return `true` *if* `n` *is a happy number, and* `false` *if not*.

## Examples

**Example 1:**

```
Input: n = 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1

```

**Example 2:**

```
Input: n = 2
Output: false

```

**Constraints:**

- `1 <= n <= 231 - 1`

## 解析

給定一個正整數 n

定義一個數字 n 是 HappyNumber

必須經過以下運算：

1. 把原本的 n 取代為原本 n 所有digit 的平方和
2. 持續步驟 1 運算直到 n = 1  或是出現重複的迴圈迴圈中不包含 1
3.  如果有出現 n = 1 代表是 happy number 否則就不是

這題的關鍵在於

經果計算平方和的運算

判斷是否有出現 1 或是迴圈

假設出現 1 回傳 true 

假設出現回圈且出現過的值都沒有 1 代表 false

計算下個方式相同

關鍵在於如何判斷有迴圈

1 使用 HashSet

![](https://i.imgur.com/hMHhCCZ.png)


時間複雜度 = O(logN) 其中 N 代表輸入數值大小

因為 digit 的長度 = logN 

而每做一次平方和計算就需要 loop 一次 digit

第2次 digit 會變成 log(logN)

所有時間複雜度 = O(logN + log(logN) + log(log(logN)) + …) = O(logN)

 而空間複雜度 同樣地 因為代表有多少數值在 回圈之前會放入 HashSet 所以會是 O(logN)

2 透過 [Floyd 迴圈判斷法](https://zh.wikipedia.org/zh-tw/Floyd%E5%88%A4%E5%9C%88%E7%AE%97%E6%B3%95)

由 [Floyd 迴圈判斷法](https://zh.wikipedia.org/zh-tw/Floyd%E5%88%A4%E5%9C%88%E7%AE%97%E6%B3%95) 可以從相差一格一前一後的出發點

前面的出發點每次移動一格

後面的出發點每次移動二格

假設有迴圈前面與後面的一定會有重和的時候

![](https://i.imgur.com/TgM87Yg.png)

時間複雜度

假設沒有 cycle 會計算到 n =1 因為一次計算兩步

然而還是跟 digit 長度有關所以 O(logN+ logN) = O(logN)

假設有 cycle 代表假設 cycle 長度是 k,  兩個數值會 起始位置 假設是差了 k-1 則會再 k-1 佈之後追平

但由於每次疊代都跟 digit長度有關 所以時間複雜還是 O(logN)

由於只需要紀錄 slow 跟 fast 兩個數值

所以空間複雜度是 O(1)

## 程式碼
```go
package sol

func isHappy(n int) bool {
	var getNext = func(n int) int {
		sum := 0
		for n > 0 {
			digit := n % 10
			n /= 10
			sum += digit * digit
		}
		return sum
	}
	slow, fast := n, getNext(n)
	for fast != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return fast == 1
}

```
## 困難點

1. 題目本身並沒有解釋的很清楚什麼是 Happy Number
2. 需要理解 Happy Number 的定義
3. 理解如何找到有迴圈出現的方法

## Solve Point

- [x]  建立一個 getNext 方法用來計算輸入數值 n 的 digit 平方和
- [x]  初始化 slow := n, fast := getNext(n)
- [x]  當 fast ≠ 1 && slow ≠ fast 時
- [x]  更新 slow = getNext(slow), fast = getNext(getNext(fast))
- [x]  當迴圈停止時 回傳 fast == 1