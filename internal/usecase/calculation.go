package usecase

import (
	"fmt"
	"math"
)

const MAX_OVERHEAD = 249
const INF = math.MaxInt32

func CalculatePacks(n int, packs []int) (*Result, error) {
	maxSum := n + MAX_OVERHEAD
	dp := make([]int, maxSum+1)
	for i := 1; i <= maxSum; i++ {
		dp[i] = INF
	}
	dp[0] = 0

	parent := make([]int, maxSum+1)
	choice := make([]int, maxSum+1)
	for i := range parent {
		parent[i] = -1
		choice[i] = -1
	}

	for x := 0; x <= maxSum; x++ {
		if dp[x] == INF {
			continue
		}
		for _, p := range packs {
			nextSum := x + p
			if nextSum <= maxSum && dp[x]+1 < dp[nextSum] {
				dp[nextSum] = dp[x] + 1
				parent[nextSum] = x
				choice[nextSum] = p
			}
		}
	}

	chosenSum := -1
	for overhead := 0; overhead <= MAX_OVERHEAD; overhead++ {
		if dp[n+overhead] != INF {
			chosenSum = n + overhead
			break
		}
	}

	if chosenSum == -1 {
		return nil, fmt.Errorf("no solution found")
	}

	packCount := make(map[int]int)
	curr := chosenSum
	for curr > 0 {
		p := choice[curr]
		packCount[p]++
		curr = parent[curr]
	}

	result := &Result{
		Requested: n,
		UsedSum:   chosenSum,
		Overhead:  chosenSum - n,
		Packs:     packCount,
	}
	return result, nil
}
