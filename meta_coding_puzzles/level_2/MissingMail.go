package main

import (
	"math"
)

func getMaxExpectedProfit(N int32, V []int32, C int32, S float64) float64 {
	dp := make([]float64, N)
	for i := int32(0); i < N; i += 1 {
		dp[i] = -1
	}
	expectedValueDp := getExpectedValueDp(N, V, S)
	answer := getBestDays(N, V, S, float64(C), expectedValueDp, dp, 0)
	return answer
}

func getBestDays(N int32, V []int32, S float64, C float64, expectedValueDp [][]float64, dp []float64, day int32) float64 {
	if day >= N {
		return 0
	}
	if dp[day] != -1 {
		return dp[day]
	}
	best := float64(0)
	for finishDay := day; finishDay < N; finishDay += 1 {
		expectedValue := expectedValueDp[day][finishDay]
		collectPackages := math.Max(0, expectedValue-C)
		bestAfterFinishDay := getBestDays(N, V, S, C, expectedValueDp, dp, finishDay+1)
		best = math.Max(best, collectPackages+bestAfterFinishDay)
	}
	dp[day] = best
	return best
}

func getExpectedValueDp(N int32, V []int32, S float64) [][]float64 {
	dp := make([][]float64, N)
	for i := int32(0); i < N; i += 1 {
		dp[i] = make([]float64, N)
		dp[i][i] = float64(V[i])
		for j := i + 1; j < N; j += 1 {
			dp[i][j] = (1.0-S)*dp[i][j-1] + float64(V[j])
		}
	}
	return dp
}

// func main() {
// 	N := int32(5)
// 	V := []int32{10, 2, 8, 6, 4}
// 	C := int32(3)
// 	S := float64(0.15)
// 	answer := getMaxExpectedProfit(N, V, C, S)
// 	fmt.Println("answer:", answer)
// }
