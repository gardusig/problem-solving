package main

type RowData struct {
	hasDown  bool
	hasRight bool
	hasCoin  bool
	hasEmpty bool
}

func max(a int32, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func getDataFromRow(row []byte) RowData {
	rowData := RowData{
		hasDown:  false,
		hasRight: false,
		hasCoin:  false,
		hasEmpty: false,
	}
	for _, char := range row {
		if char == '*' {
			rowData.hasCoin = true
		}
		if char == '>' {
			rowData.hasRight = true
		}
		if char == 'v' {
			rowData.hasDown = true
		}
		if char == '.' {
			rowData.hasEmpty = true
		}
	}
	return rowData
}

func rowContainsRightLoop(rowData RowData) bool {
	return rowData.hasRight && !rowData.hasDown
}

func rowCanGoDown(rowData RowData) bool {
	return rowData.hasDown || rowData.hasEmpty || rowData.hasCoin
}

func getRowCoinSum(row []byte) int32 {
	sum := int32(0)
	for _, char := range row {
		if char == '*' {
			sum += 1
		}
	}
	return sum
}

func getRowCoinBestDown(rowData RowData, row []byte) int32 {
	best := int32(0)
	if rowData.hasCoin {
		best += 1
	}
	if !rowData.hasDown {
		return best
	}
	dp := make([]int32, len(row))
	for column := range row {
		dp[column] = -1
	}
	for column := range row {
		if row[column] == '>' {
			best = max(best, getBestSequence(row, column, dp))
		}
	}
	return best
}

func getBestSequence(row []byte, column int, dp []int32) int32 {
	if row[column] == 'v' {
		return 0
	}
	if dp[column] != -1 {
		return dp[column]
	}
	best := int32(0)
	if row[column] == '*' {
		best += 1
	}
	nxtColumn := (column + 1) % len(row)
	best += getBestSequence(row, nxtColumn, dp)
	dp[column] = best
	return best
}

func getMaxCollectableCoins(R int32, C int32, G [][]byte) int32 {
	best := int32(0)
	current := int32(0)
	for row := int32(0); row < R; row += 1 {
		rowData := getDataFromRow(G[row])
		if rowContainsRightLoop(rowData) {
			best = max(best, current+getRowCoinSum(G[row]))
		}
		if !rowCanGoDown(rowData) {
			break
		}
		current += getRowCoinBestDown(rowData, G[row])
		best = max(best, current)
	}
	return best
}
