package main

import (
	"sort"
)

func getSecondsElapsed(C int64, N int32, A []int64, B []int64, K int64) int64 {
	tunnelList := createTunnelList(N, A, B)
	fullLapTunnelLength := getFullLapTunnelLength(tunnelList)
	return getTotalLength(C, fullLapTunnelLength, tunnelList, K)
}

type Pair struct {
	First  int64
	Second int64
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].First < p[j].First
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func createTunnelList(N int32, A []int64, B []int64) PairList {
	tunnelList := make(PairList, N)
	for i := int32(0); i < N; i += 1 {
		tunnelList[i] = Pair{
			First:  A[i],
			Second: B[i],
		}
	}
	sort.Sort(tunnelList)
	return tunnelList
}

func getFullLapTunnelLength(tunnelList PairList) int64 {
	totalLength := int64(0)
	for _, tunnel := range tunnelList {
		totalLength += tunnel.Second - tunnel.First
	}
	return totalLength
}

func getTotalLength(C int64, fullLapTunnelLength int64, tunnelList PairList, K int64) int64 {
	fullLaps := K / fullLapTunnelLength
	lastLap := getFinalLapLength(tunnelList, K%fullLapTunnelLength)
	if (K % fullLapTunnelLength) == 0 {
		if fullLaps > 0 {
			fullLaps -= 1
		}
		lastLap = tunnelList[len(tunnelList)-1].Second
	}
	return fullLaps*C + lastLap
}

func getFinalLapLength(tunnelList PairList, remaining int64) int64 {
	for _, tunnel := range tunnelList {
		currentTunnelLength := tunnel.Second - tunnel.First
		if remaining <= currentTunnelLength {
			return tunnel.First + remaining
		}
		remaining -= currentTunnelLength
	}
	return -1
}

// func main() {
// 	C := int64(50)
// 	N := int32(3)
// 	A := []int64{39, 19, 28}
// 	B := []int64{49, 27, 35}
// 	K := int64(15)
// 	answer := getSecondsElapsed(C, N, A, B, K)
// 	fmt.Println(answer)
// }
