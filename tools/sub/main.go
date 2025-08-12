package main

import "fmt"

// 全組み合わせ生成関数
func generateAllCombinations(n, minVal, maxVal int) [][]int {
	var result [][]int
	current := make([]int, 0, n)

	var backtrack func(position int)
	backtrack = func(position int) {
		// ベースケース: すべての位置を埋めた
		if position == n {
			// スライスをコピーして追加
			combination := make([]int, len(current))
			copy(combination, current)
			result = append(result, combination)
			return
		}

		// 現在の位置にminValからmaxValまでの値を試す
		for val := minVal; val <= maxVal; val++ {
			current = append(current, val)
			backtrack(position + 1)
			current = current[:len(current)-1] // バックトラック
		}
	}

	backtrack(0)
	return result
}

func main() {
	const (
		n      = 5  // 配列の長さ
		minVal = 1  // 最小値
		maxVal = 10 // 最大値
	)

	combinations := generateAllCombinations(n, minVal, maxVal)

	// 競技プログラミング形式で標準出力
	for _, combination := range combinations {
		fmt.Println(n)
		for i, val := range combination {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}