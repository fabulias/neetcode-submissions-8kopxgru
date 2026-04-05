func checkValidAnswer(s string, tMap map[byte]int) bool {
	sMap := make(map[byte]int, 0)
	for ix := 0; ix < len(s); ix++ {
		sMap[s[ix]]++
	}
	for k, v := range tMap {
		if v > sMap[k]  {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
    resp := ""
	minWindow := 0
	tMap := make(map[byte]int, 0)
	
	for ix := 0; ix < len(t); ix++ {
		tMap[t[ix]]++
	}

	for ix := 0; ix < len(s); ix++ {
		for iy := ix; iy < len(s); iy++ {
			if iy - ix + 1 < len(t) {
				continue
			}
			if checkValidAnswer(s[ix:iy+1], tMap) {
				currentLen := iy - ix + 1
    			if minWindow == 0 || currentLen < minWindow {
     			   minWindow = currentLen
     			   resp = s[ix:iy+1]
				}
			}
		}
	}
	return resp
}
