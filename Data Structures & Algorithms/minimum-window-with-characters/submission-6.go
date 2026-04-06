func minWindow(s string, t string) string {
	if t == ""{
		return ""
	}

    resp := ""
	minWindow := math.MaxInt
	needMap := make(map[byte]int, 0)
	window := make(map[byte]int, 0)
	
	for ix := 0; ix < len(t); ix++ {
		needMap[t[ix]]++
		window[t[ix]] = 0
	}

	have, need := 0, len(needMap)

	left, right := 0, 0
	for right < len(s) {
		if _, ok := needMap[s[right]]; ok {
			window[s[right]]++
			if window[s[right]] == needMap[s[right]] {
				have++
			}
		}
		
		for have == need {
			if right - left + 1 < minWindow {
				resp = s[left:right+1]
				minWindow = right - left + 1
			}
			if _, ok := needMap[s[left]]; ok {
				window[s[left]]--
				if window[s[left]] < needMap[s[left]] {
					have--
				}
			}
			left++
		}
		right++
	}
	return resp
}
