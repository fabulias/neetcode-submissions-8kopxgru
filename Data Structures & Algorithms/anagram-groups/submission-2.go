import "maps"
import "slices"

func groupAnagrams(strs []string) [][]string {
	solutionsMap := make(map[[26]int][]string,0)
	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c - 'a']++
		}

		solutionsMap[count] = append(solutionsMap[count], s)
	}

	return slices.Collect(maps.Values(solutionsMap))
}
