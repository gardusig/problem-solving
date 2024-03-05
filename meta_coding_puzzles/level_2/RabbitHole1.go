package main

func getMaxVisitableWebpages(N int32, L []int32) int32 {
	g, t := createGraph(int(N), L)
	finishingTimes := getGraphFinishingTimes(int(N), g)
	components := getGraphComponents(int(N), t, finishingTimes)
	condensedGraph := getCondensedGraph(g, components)
	return getGreatestPath(components, condensedGraph)
}

func createGraph(N int, L []int32) ([][]int, [][]int) {
	g := make([][]int, N)
	t := make([][]int, N)
	for i := 0; i < N; i += 1 {
		j := int(L[i]) - 1
		g[i] = append(g[i], j)
		t[j] = append(t[j], i)
	}
	return g, t
}

func getGraphFinishingTimes(N int, g [][]int) []int {
	finishingTimes := make([]int, 0)
	visited := make([]bool, N)
	for i := 0; i < N; i += 1 {
		visited[i] = false
	}
	for i := 0; i < N; i += 1 {
		if visited[i] {
			continue
		}
		dfs(i, g, &visited, &finishingTimes)
	}
	return finishingTimes
}

func dfs(i int, g [][]int, visited *[]bool, finishingTimes *[]int) {
	(*visited)[i] = true
	for _, j := range g[i] {
		if (*visited)[j] {
			continue
		}
		dfs(j, g, visited, finishingTimes)
	}
	*finishingTimes = append(*finishingTimes, i)
}

func getGraphComponents(N int, t [][]int, finishingTimes []int) [][]int {
	visited := make([]bool, N)
	for i := 0; i < N; i += 1 {
		visited[i] = false
	}
	components := make([][]int, 0)
	for i := len(finishingTimes) - 1; i >= 0; i -= 1 {
		vertex := finishingTimes[i]
		if visited[vertex] {
			continue
		}
		component := make([]int, 0)
		dfsTranspose(vertex, t, &visited, &component)
		components = append(components, component)
	}
	return components
}

func dfsTranspose(i int, t [][]int, visited *[]bool, component *[]int) int32 {
	size := int32(1)
	(*visited)[i] = true
	(*component) = append(*component, i)
	for _, j := range t[i] {
		if (*visited)[j] {
			continue
		}
		size += dfsTranspose(j, t, visited, component)
	}
	return size
}

func getCondensedGraph(graph [][]int, components [][]int) [][]int {
	vertexComponent := make([]int, len(graph))
	for componentIndex, component := range components {
		for _, vertex := range component {
			vertexComponent[vertex] = componentIndex
		}
	}
	componentGraph := make([][]int, len(components))
	for i, edges := range graph {
		for _, j := range edges {
			componentI := vertexComponent[i]
			componentJ := vertexComponent[j]
			if componentI != componentJ {
				componentGraph[componentI] = append(componentGraph[componentI], componentJ)
			}
		}
	}
	return componentGraph
}

func getGreatestPath(components [][]int, condensedGraph [][]int) int32 {
	greatestPath := int32(0)
	dp := make([]int32, len(components))
	for i := 0; i < len(components); i += 1 {
		dp[i] = -1
	}
	for startingVertex := 0; startingVertex < len(components); startingVertex += 1 {
		current := solve(startingVertex, components, condensedGraph, &dp)
		if current > greatestPath {
			greatestPath = current
		}
	}
	return greatestPath
}

func solve(vertex int, components [][]int, condensedGraph [][]int, dp *[]int32) int32 {
	if (*dp)[vertex] != -1 {
		return (*dp)[vertex]
	}
	greatestPath := int32(0)
	for _, nxtVertex := range condensedGraph[vertex] {
		nxtPath := solve(nxtVertex, components, condensedGraph, dp)
		if nxtPath > greatestPath {
			greatestPath = nxtPath
		}
	}
	greatestPath += int32(len(components[vertex]))
	(*dp)[vertex] = greatestPath
	return greatestPath
}

// func main() {
// 	N := int32(5)
// 	L := []int32{5, 3, 4, 5, 2}
// 	fmt.Println(getMaxVisitableWebpages(N, L))
// }
