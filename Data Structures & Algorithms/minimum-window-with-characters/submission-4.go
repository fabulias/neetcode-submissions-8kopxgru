func minWindow(s string, t string) string {
	if t == ""{
		return ""
	}

    resp := ""
	minWindow := math.MaxInt
	tMap := make(map[byte]int, 0)
	
	
	for ix := 0; ix < len(t); ix++ {
		tMap[t[ix]]++
	}

	for ix := 0; ix < len(s); ix++ {
		window := make(map[byte]int, 0)
			for iy := ix; iy < len(s); iy++ {
				if _, ok := tMap[s[iy]]; ok {
					window[s[iy]]++
				}
				flag := true
				if len(window) == len(tMap) {
					for k, v := range tMap {
						if window[k] < v {
							flag = false
							break
						}
					}
					if flag && iy - ix + 1 < minWindow {
						resp = s[ix:iy+1]
						minWindow = iy - ix + 1
					}

				}		
			}
	}
	return resp
}
