func pacificAtlantic(heights [][]int) [][]int {
    rows, cols := len(heights),len(heights[0])
	type point struct {
		x int
		y int
	}
	pacific, atlantic := make(map[point]bool), make(map[point]bool)

	var dfs func(r, c int, visited map[point]bool, prevHeight int)
	dfs = func(r, c int, visited map[point]bool, prevHeight int) {
		if visited[point{r,c}] || r < 0 || c < 0 || 
		r == rows || c == cols ||
		heights[r][c] < prevHeight {
			return
		}
		visited[point{r,c}] = true
		dfs(r-1,c,visited,heights[r][c])
		dfs(r+1,c,visited,heights[r][c])
		dfs(r,c-1,visited,heights[r][c])
		dfs(r,c+1,visited,heights[r][c])
	}

	for ix := range cols {
		dfs(0, ix, pacific, heights[0][ix])
		dfs(rows-1, ix, atlantic, heights[rows-1][ix])
	}
	for ix := range rows {
		dfs(ix, 0, pacific, heights[ix][0])
		dfs(ix, cols-1, atlantic, heights[ix][cols-1])
	}

	cellsCanFlow := make([][]int, 0)
	for r := range rows {
		for c := range cols {
			if pacific[point{r,c}] && atlantic[point{r,c}] {
				cellsCanFlow = append(cellsCanFlow, []int{r,c})
			}
		}
	}
	return cellsCanFlow
}
