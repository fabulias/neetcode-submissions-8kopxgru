func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int)
	for _, pre := range prerequisites {
		preMap[pre[0]] = append(preMap[pre[0]], pre[1])
	}

	visited := make(map[int]bool, 0)
	
	var dfs func(course int) bool 
	dfs =  func(course int) bool {
		if visited[course] {
			return false
		}

		visited[course] = true	
		for _, pre := range preMap[course] {
			if !dfs(pre) { return false }
		}
		
		visited[course] = false
		return true
	}
	
	for course := range prerequisites {
		if !dfs(course) {
			return false
		}
	}
		
	return true
}
