func validTree(n int, edges [][]int) bool {
	adjList := make(map[int][]int, 0)

	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}
    
	visited := make(map[int]bool)
	
	var dfs func(curr, prev int) bool
	dfs = func(curr, prev int) bool {
		if visited[curr] {
			return false
		}
		visited[curr] = true
		for _, neighbor := range adjList[curr] {
			if neighbor == prev {
				continue
			}
			if !dfs(neighbor, curr) {
				return false
			}
		}
		return true
	}
	return dfs(0, -1) && len(visited) == n
	
}
