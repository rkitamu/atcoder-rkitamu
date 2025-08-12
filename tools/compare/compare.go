package main

import (
	"fmt"
	"math"
	"sort"
)

// テストケース生成
func generateAllCombinations(n, minVal, maxVal int) [][]int {
	var result [][]int
	current := make([]int, 0, n)

	var backtrack func(position int)
	backtrack = func(position int) {
		if position == n {
			combination := make([]int, len(current))
			copy(combination, current)
			result = append(result, combination)
			return
		}

		for val := minVal; val <= maxVal; val++ {
			current = append(current, val)
			backtrack(position + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}

// 貪欲法（あなたの実装）
func solveGreedy(tasks []int) int {
	n := len(tasks)
	t := make([]int, n)
	copy(t, tasks)

	sort.Slice(t, func(i, j int) bool {
		return t[i] > t[j]
	})

	oven1 := 0
	oven2 := 0
	maxt := 0
	sum := 0

	for i := 0; i < n; i++ {
		if oven1 <= oven2 {
			oven1 += t[i]
		} else {
			oven2 += t[i]
		}
		if t[i] > maxt {
			maxt = t[i]
		}
		sum += t[i]
	}

	ans := oven1
	if oven2 > ans {
		ans = oven2
	}

	// 下界チェック
	ell := int(math.Ceil(float64(sum) / 2.0))
	if maxt < ell {
		ans = ell
	}	
	return ans
}

// 最適解（DP）
func solveDPOptimal(tasks []int) int {
	n := len(tasks)
	totalTime := 0
	for _, t := range tasks {
		totalTime += t
	}

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, totalTime+1)
	}
	dp[0][0] = true

	for i := 0; i < n; i++ {
		task := tasks[i]
		for sum := 0; sum <= totalTime; sum++ {
			dp[i+1][sum] = dp[i][sum]
			if sum >= task && dp[i][sum-task] {
				dp[i+1][sum] = true
			}
		}
	}

	maxSum := 0
	for i := totalTime / 2; i >= 0; i-- {
		if dp[n][i] {
			maxSum = i
			break
		}
	}

	otherSum := totalTime - maxSum
	if maxSum > otherSum {
		return maxSum
	}
	return otherSum
}

func main() {
	testCases := generateAllCombinations(5, 1, 10)

	for i, testCase := range testCases {
		greedyResult := solveGreedy(testCase)
		optimalResult := solveDPOptimal(testCase)

		if greedyResult != optimalResult {
			fmt.Printf("不一致発見 (ケース %d)\n", i+1)
			fmt.Printf("テストケース: %v\n", testCase)
			fmt.Printf("貪欲法の結果: %d\n", greedyResult)
			fmt.Printf("最適解: %d\n", optimalResult)
			return
		}
	}

	fmt.Printf("全 %d ケースで一致しました。\n", len(testCases))
}
